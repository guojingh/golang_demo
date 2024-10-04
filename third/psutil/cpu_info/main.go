package main

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
)

// 获取cpu信息
func getCpuInfo() {
	cpuInfos, err := cpu.Info()
	if err != nil {
		fmt.Printf("get cpu info failed, err:%v\n", err)
	}

	//获取cpu信息
	for _, ci := range cpuInfos {
		fmt.Println(ci)
	}

	for {
		// 获取Cpu使用率
		percent, _ := cpu.Percent(time.Second, false)
		fmt.Printf("cpu precent:%v\n", percent)
	}
}

// cpu 负载 在windows下可能会出现问题
func getLoad() {
	info, err := load.Avg()
	if err != nil {
		fmt.Printf("get load failed, err:%v\n", err)
		return
	}
	fmt.Println(info)
}

// 内存信息
func getMemInfo() {
	info, err := mem.VirtualMemory()
	if err != nil {
		fmt.Printf("get mem info failed, err:%v\n", err)
		return
	}

	fmt.Println(info)
}

// 主机信息
func getHostInfo() {
	hInfo, _ := host.Info()
	fmt.Printf("host info:%v uptime:%v boottime:%v\n", hInfo, hInfo.Uptime, hInfo.BootTime)
}

// Disk 磁盘信息
func getDiskInfo() {
	// 获取所有分区信息
	parts, err := disk.Partitions(true)
	if err != nil {
		fmt.Printf("get disk info failed, err:%v\n", err)
		return
	}
	fmt.Println(parts)
	for _, part := range parts {
		partInfo, err := disk.Usage(part.Mountpoint)
		if err != nil {
			fmt.Printf("get disk info failed, err:%v\n", err)
			return
		}
		fmt.Println(partInfo)
	}
	// 磁盘IO
	ioStat, _ := disk.IOCounters()
	for k, v := range ioStat {
		fmt.Printf("%v:%v\n", k, v)
	}
}

// net 相关
func getNetInfo() {
	netIOs, err := net.IOCounters(true)
	if err != nil {
		fmt.Printf("get net info failed, err:%v\n", err)
		return
	}

	for _, v := range netIOs {
		fmt.Printf("%v:%v:%v\n", v.Name, v.BytesRecv, v.BytesSent)
	}
}

func main() {
	getCpuInfo()
	//getLoad()
	//getMemInfo()
	//getHostInfo()
	//getDiskInfo()
	//getNetInfo()
}
