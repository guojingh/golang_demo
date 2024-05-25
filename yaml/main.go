package main

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Broker        string `yaml:"broker" envconfig:"BROKER"`
	DefaultQueue  string `yaml:"default_queue" envconfig:"DEFAULT_QUEUE"`
	ResultBackend string `yaml:"result_backend" envconfig:"RESULT_BACKEND"`
	Redis         string `yaml:"redis"`
}

func main() {

	//打开文件
	file, err := os.Open("yaml/test.yaml")
	if err != nil {
		fmt.Println("打开文件失败")
	}

	//读取文件
	bytes := make([]byte, 1024)
	count, err := file.Read(bytes)
	if err != nil {
		fmt.Println("文件读取失败")
	}

	cnf := new(Config)
	if err := yaml.Unmarshal(bytes[:count], cnf); err != nil {
		fmt.Println("yaml 序列化失败")
	}

	fmt.Printf("conf=%v", cnf)
}
