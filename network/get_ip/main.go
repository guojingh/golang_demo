package main

import (
	"fmt"
	"log"
	"net"
)

func GetLocalIP() (ip string, err error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return
	}

	for _, addr := range addrs {
		ipAddr, ok := addr.(*net.IPNet)
		if !ok {
			continue
		}
		fmt.Println(ipAddr)
		if ipAddr.IP.IsLoopback() {
			continue
		}
		if !ipAddr.IP.IsGlobalUnicast() {
			continue
		}
		return ipAddr.IP.String(), nil
	}
	return
}

func GetOutBoundIP() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)
	//fmt.Println(localAddr.String())
	return localAddr.IP.String()
}

func main() {
	/*	ip, err := GetLocalIP()
		if err != nil {
			return
		}
		fmt.Println(ip)*/
	fmt.Println(GetOutBoundIP())
}
