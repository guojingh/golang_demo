package util

import (
	"regexp"
	"strconv"
	"strings"
)

const (
	B = 1 << (iota * 10)
	KB
	MB
	GB
	TB
	PB
)

const (
	l = 2 + iota
	n
	d
)

const (
	A = 1
	R
	C
	D
	E
)

func ParseSize(size string) (int64, string) {
	//fmt.Printf("输入内容为%s\n", size)

	//TODO 做传参解析 size: 1KB 100KB 1MB 2MB 1GB
	re, _ := regexp.Compile("[0-9]+")
	unit := string(re.ReplaceAll([]byte(size), []byte("")))
	num, _ := strconv.ParseInt(strings.Replace(size, unit, "", 1), 10, 64)
	upper := strings.ToUpper(unit)

	var maxMemory int64
	switch upper {
	case "B":
		maxMemory = num * B
	case "KB":
		maxMemory = num * KB
	case "MB":
		maxMemory = num * MB
	case "GB":
		maxMemory = num * GB
	case "TB":
		maxMemory = num * TB
	case "PB":
		maxMemory = num * PB
	default:
		maxMemory = num * KB
	}
	//fmt.Printf("输出内容为数字部分=%d，字符串部分=%s\n", num, upper)
	return maxMemory, strconv.FormatInt(maxMemory, 10)
}

func GetValSize(val interface{}) int64 {
	return int64(1009)
}

/*func main() {
	ParseSize("100mb")
}*/
