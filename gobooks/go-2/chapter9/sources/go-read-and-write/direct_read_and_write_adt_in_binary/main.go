package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"os"
)

// 通过标准库提供的binary包直接读写抽象数据类型
// 各个字段的类型都采用了定长类型，这是binary包对直接操作的抽象数据类型的约束
type Player struct {
	Name   [20]byte
	Age    int16
	Gender [6]byte
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
		// 采用大端字节序对Player示例进行编码
		err = binary.Write(f, binary.BigEndian, &player)
		if err != nil {
			return err
		}
	}

	return nil
}

func main() {
	var players [3]Player

	copy(players[0].Name[:], []byte("Tommy"))
	players[0].Age = 18
	copy(players[0].Gender[:], []byte("male"))

	copy(players[1].Name[:], []byte("Lily"))
	players[1].Age = 19
	copy(players[1].Gender[:], []byte("female"))

	copy(players[2].Name[:], []byte("Lily"))
	players[2].Age = 20
	copy(players[2].Gender[:], []byte("female"))

	err := directWriteADTToFile("players.dat", players[:])
	if err != nil {
		fmt.Println("write file error:", err)
		return
	}

	f, err := os.Open("players.dat")
	if err != nil {
		fmt.Println("open file error:", err)
		return
	}

	var player Player
	for {
		err := binary.Read(f, binary.BigEndian, &player)
		if err == io.EOF {
			fmt.Println("Read meet EOF")
			return
		}
		if err != nil {
			fmt.Println("read file error:", err)
			return
		}
		fmt.Printf("%s %d %s\n", player.Name, player.Age, player.Gender)
	}
}
