package main

import (
	"fmt"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api/write"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
)

const (
	bucket = "monitor"
	org    = "root"
	token  = "34xdpYgz-xoSCqhQ5-Sy9SWrlZRBG-E8C_rmXP9f8YsayzFCvibKMGDrxVYE24-N2LizCrlttiwN6-HoyhK76w=="
	// Store the URL of your InfluxDB instance
	url = "http://172.16.56.129:8086"
)

var (
	client                 influxdb2.Client
	lastNetIOStatTimeStamp int64    //上一次获取网络IO数据的时间点
	lastNetInfo            *NetInfo // 上一次的网络IO数据
)

func sendMsg(info SysInfo) {
	writeAPI := client.WriteAPI(org, bucket)
	p := new(write.Point)
	// 根据传入数据的类型插入数据
	switch info.Data.(type) {
	case *CpuInfo:
		p = influxdb2.NewPointWithMeasurement("cpu_percent").
			AddTag("cpu", "cpu").
			AddField("cpu_percent", info.Data.(*CpuInfo).CpuPercent)
	case *MemInfo:
		p = influxdb2.NewPointWithMeasurement("memory").
			AddTag("mem", "mem").
			AddField("total", info.Data.(*MemInfo).Total).
			AddField("available", info.Data.(*MemInfo).Available).
			AddField("used", info.Data.(*MemInfo).Used).
			AddField("UsedPercent", info.Data.(*MemInfo).UsedPercent)
	case *DiskInfo:
		for k, v := range info.Data.(*DiskInfo).PartitionUsageStat {
			p = influxdb2.NewPointWithMeasurement("disk").
				AddTag("disk", k).
				AddField("total", v.Total).
				AddField("free", v.Free).
				AddField("used", v.Used).
				AddField("used_percent", v.UsedPercent).
				AddField("inodes_total", v.InodesTotal).
				AddField("inodes_used", v.InodesTotal).
				AddField("inodes_free", v.InodesFree).
				AddField("inodes_used_percent", v.InodesUsedPercent)
		}
	case *NetInfo:
		for k, v := range info.Data.(*NetInfo).NetIOCountersStat {
			p = influxdb2.NewPointWithMeasurement("net").
				AddTag("name", k).
				AddField("bytes_sent_rate", v.BytesSentRate).
				AddField("bytes_recv_rate", v.BytesRecvRate).
				AddField("packets_sent_rate", v.PacketsSentRate).
				AddField("Packets_recv_rate", v.PacketsRecvRate)
		}
	}
	writeAPI.WritePoint(p)
	writeAPI.Flush()
}

func getCpuInfo() {
	var info SysInfo
	var cpuInfo = new(CpuInfo)
	// 获取Cpu使用率
	percent, _ := cpu.Percent(time.Second, false)
	fmt.Printf("cpu precent:%v\n", percent)

	cpuInfo.CpuPercent = percent[0]
	info.InfoType = CpuInfoType
	info.Data = cpuInfo
	sendMsg(info)
}

func initClient() {
	client = influxdb2.NewClient(url, token)
}

// 内存信息
func getMemInfo() {
	var sysInfo SysInfo
	var memInfo = new(MemInfo)
	info, err := mem.VirtualMemory()
	if err != nil {
		fmt.Printf("get mem info failed, err:%v\n", err)
		return
	}
	memInfo.Total = info.Total
	memInfo.Available = info.Available
	memInfo.Used = info.Used
	memInfo.UsedPercent = info.UsedPercent
	memInfo.Buffers = info.Buffers
	memInfo.Cached = info.Cached

	sysInfo.InfoType = MemInfoType
	sysInfo.Data = memInfo
	sendMsg(sysInfo)
}

func getDiskInfo() {
	var sysInfo SysInfo
	var disInfo = &DiskInfo{
		PartitionUsageStat: make(map[string]*disk.UsageStat, 16),
	}
	parts, err := disk.Partitions(true)
	if err != nil {
		fmt.Printf("get disk info failed, err:%v\n", err)
		return
	}

	for _, part := range parts {
		//拿到每一个分区的信息
		usageStat, err := disk.Usage(part.Mountpoint) // 传入挂载点
		if err != nil {
			fmt.Printf("get %s usage stat failed, err:%v\n", part.Mountpoint, err)
			continue
		}
		disInfo.PartitionUsageStat[part.Mountpoint] = usageStat
	}
	sysInfo.InfoType = DiskInfoType
	sysInfo.Data = disInfo
	sendMsg(sysInfo)
}

func getNetInfo() {
	var sysInfo SysInfo
	var netInfo = &NetInfo{
		NetIOCountersStat: make(map[string]*IOStat, 8),
	}

	currentTimeStamp := time.Now().Unix()
	netIOs, err := net.IOCounters(true)
	if err != nil {
		fmt.Printf("get net io info failed, err:%v\n", err)
		return
	}

	for _, netIO := range netIOs {
		var ioStat = new(IOStat)
		ioStat.BytesSent = netIO.BytesSent
		ioStat.BytesRecv = netIO.BytesRecv
		ioStat.PacketsSent = netIO.PacketsSent
		ioStat.BytesRecv = netIO.BytesRecv
		netInfo.NetIOCountersStat[netIO.Name] = ioStat
		// 开始计算网卡相关速率
		if lastNetIOStatTimeStamp == 0 || lastNetInfo == nil {
			continue
		}
		// 计算时间间隔
		interval := currentTimeStamp - lastNetIOStatTimeStamp
		// 计算速率
		ioStat.BytesSentRate = (ioStat.BytesSent - lastNetInfo.NetIOCountersStat[netIO.Name].BytesSent) / uint64(interval)
		ioStat.BytesRecvRate = (ioStat.BytesRecv - lastNetInfo.NetIOCountersStat[netIO.Name].BytesRecv) / uint64(interval)
		ioStat.PacketsSentRate = (ioStat.PacketsSent - lastNetInfo.NetIOCountersStat[netIO.Name].PacketsSent) / uint64(interval)
		ioStat.PacketsRecvRate = (ioStat.PacketsRecv - lastNetInfo.NetIOCountersStat[netIO.Name].PacketsRecv) / uint64(interval)
	}
	lastNetIOStatTimeStamp = currentTimeStamp
	lastNetInfo = netInfo
	sysInfo.InfoType = NetInfoType
	sysInfo.Data = netInfo
	sendMsg(sysInfo)
}

func run(interval time.Duration) {
	ticker := time.Tick(interval)
	for _ = range ticker {
		getCpuInfo()
		getMemInfo()
		getDiskInfo()
		getNetInfo()
	}
}

func main() {
	initClient()
	run(time.Second)
}
