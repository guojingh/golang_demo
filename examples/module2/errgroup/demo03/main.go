package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"golang.org/x/sync/errgroup"
)

func main() {
	g := new(errgroup.Group)

	files := []string{
		"file1.txt",
		"file2.txt",
		"file3.txt",
	}

	for _, filename := range files {
		filename := filename
		g.Go(func() error {
			file, err := os.Open(filepath.Join("/Users/v_guojinghu/project/git_study/golang_demo/examples/module2/errgroup", filename))
			if err != nil {
				return err
			}
			defer file.Close()
			data, err := ioutil.ReadAll(file)
			if err != nil {
				return err
			}
			fmt.Println(string(data))
			return nil
		})
	}

	if err := g.Wait(); err != nil {
		fmt.Println("one of the file reads returned an error:", err)
	}
}
