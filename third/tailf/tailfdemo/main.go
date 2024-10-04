package main

import (
	"fmt"
	"time"

	"github.com/hpcloud/tail"
)

func main() {

	filename := `D:/git/xx.log`
	config := tail.Config{
		ReOpen:    true,                                 // 如果文件被重命名或移动，重新打开文件
		Follow:    true,                                 // 持续跟踪文件的新增内容
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2}, // 从文件末尾开始读取
		MustExist: false,                                // 文件不存在时不报错
		Poll:      true,                                 // 使用轮询方式检查文件变化
	}

	// 打开文件开始读取数据
	tails, err := tail.TailFile(filename, config)
	if err != nil {
		fmt.Printf("tailfile %s failed, err:%v\n", filename, err)
		return
	}

	// 开始读取数据
	var (
		msg *tail.Line
		ok  bool
	)

	for {
		msg, ok = <-tails.Lines // chan
		if !ok {
			fmt.Printf("tailfile file close reopen, filename:%s\n", tails.Filename)
			time.Sleep(time.Second)
			continue
		}
		fmt.Println("msg:", msg.Text)
	}
}
