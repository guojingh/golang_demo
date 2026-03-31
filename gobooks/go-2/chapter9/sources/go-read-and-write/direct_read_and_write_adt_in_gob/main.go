package main

import (
	"encoding/gob"
	"fmt"
	"io"
	"os"
)

// 利用god包的Decode和Encode的方法进行读写
// god包支持对任意抽象数据类型示例的直接读写，唯一的约束是自定义结构体类型中的字段至少有一个是导出的
type Player struct {
	Name   string
	Age    int
	Gender string
}

func directWriteADTToFile(path string, players []Player) error {
	f, err := os.Create(path)
	if err != nil {
		fmt.Println("open file error:", err)
		return err
	}
	defer func() {
		f.Sync()
		f.Close()
	}()
	enc := gob.NewEncoder(f)

	for _, player := range players {
		err = enc.Encode(player)
		if err != nil {
			return err
		}
	}

	return nil
}

func main() {
	var players = []Player{
		{Name: "Tommy", Age: 18, Gender: "male"},
		{Name: "Lily", Age: 19, Gender: "female"},
		{Name: "Lily", Age: 20, Gender: "female"},
	}

	err := directWriteADTToFile("players.dat", players)
	if err != nil {
		fmt.Println("write file error:", err)
		return
	}
	f, err := os.Open("players.dat")
	if err != nil {
		fmt.Println("open the file error:", err)
		return
	}
	var player Player
	dec := gob.NewDecoder(f)
	for {
		err := dec.Decode(&player)
		if err == io.EOF {
			fmt.Println("read meet EOF")
			return
		}
		if err != nil {
			fmt.Println("read file error:", err)
			return
		}
		fmt.Printf("%v\n", player)
	}
}
