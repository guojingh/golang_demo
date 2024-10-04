package main

import (
	"github.com/shirou/gopsutil/disk"
)

const (
	CpuInfoType  = "cpu"
	MemInfoType  = "mem"
	DiskInfoType = "disk"
	NetInfoType  = "net"
)

type SysInfo struct {
	InfoType string
	Ip       string
	Data     interface{}
}

type CpuInfo struct {
	CpuPercent float64 `json:"cpu_percent"`
}

type MemInfo struct {
	Total       uint64  `json:"total"`
	Available   uint64  `json:"available"`
	Used        uint64  `json:"used"`
	UsedPercent float64 `json:"used_percent"`
	Buffers     uint64  `json:"buffers"`
	Cached      uint64  `json:"cached"`
}

type UsageStat struct {
	Path              string  `json:"path"`
	Fstype            string  `json:"fstype"`
	Total             uint64  `json:"total"`
	Free              uint64  `json:"free"`
	Used              uint64  `json:"used"`
	UsedPercent       float64 `json:"used_percent"`
	InodesTotal       uint64  `json:"inodes_total"`
	InodesUsed        uint64  `json:"inodes_used"`
	InodesFree        uint64  `json:"inodes_free"`
	InodesUsedPercent float64 `json:"inodes_used_percent"`
}

type DiskInfo struct {
	PartitionUsageStat map[string]*disk.UsageStat
}

type IOStat struct {
	BytesSent       uint64 `json:"bytes_sent"`
	BytesRecv       uint64 `json:"byte_recv"`
	PacketsSent     uint64 `json:"packets_sent"`
	PacketsRecv     uint64 `json:"packets_recv"`
	BytesSentRate   uint64 `json:"bytes_sent_rate"`
	BytesRecvRate   uint64 `json:"bytes_recv_rate"`
	PacketsSentRate uint64 `json:"packets_sent_rate"`
	PacketsRecvRate uint64 `json:"packets_recv_rate"`
}

type NetInfo struct {
	NetIOCountersStat map[string]*IOStat
}
