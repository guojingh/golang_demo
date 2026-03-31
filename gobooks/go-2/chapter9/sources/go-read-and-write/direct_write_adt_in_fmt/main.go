package main

import (
	"fmt"
	"os"
)

type Player struct {
	name   string
	age    int
	gender string
}

func (p Player) String() string {
	return fmt.Sprintf("%s %d %s", p.name, p.age, p.gender)
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
	for _, player := range players {
		//当使用fmt.Fprintf向文件（io.Writer实例）写入数据时（通过%s），Player类型的String方法便会被调用
		_, err := fmt.Fprintf(f, "%s\n", player)
		if err != nil {
			return err
		}
	}

	return nil
}

func main() {
	var players = []Player{
		{"Tommy", 18, "male"},
		{"Lucy", 17, "female"},
		{"George", 19, "male"},
	}
	err := directWriteADTToFile("players.dat", players)
	if err != nil {
		fmt.Println("write file error:", err)
		return
	}
}
