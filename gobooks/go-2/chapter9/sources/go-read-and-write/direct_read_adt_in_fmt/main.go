package main

import (
	"fmt"
	"io"
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

func main() {

	f, err := os.Open("players.dat")
	if err != nil {
		fmt.Println("open file error:", err)
		return
	}
	var player Player
	for {
		_, err := fmt.Fscanf(f, "%s %d %s", &player.name, &player.age, &player.gender)
		if err == io.EOF {
			fmt.Println("read meet EOF")
			return
		}
		if err != nil {
			fmt.Println("read file error:", err)
			return
		}
		fmt.Printf("%s %d %s\n", player.name, player.age, player.gender)
	}

}
