package common

import "net"

// 要搜集的日志的配置结构体
type CollectEntry struct {
	Path  string `json:"path"`  // 去哪个路径读取日志文件
	Topic string `json:"topic"` // 日志文件发往哪个 topic
}

// 获取本机IP
func GetOutBoundIP() (ip string, err error) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return
	}
	defer conn.Close()

	ip = conn.LocalAddr().(*net.UDPAddr).IP.String()
	//fmt.Println(localAddr.String())
	return
}
