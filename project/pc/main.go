package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/cncamp/golang/project/pc/many_many"
	"github.com/cncamp/golang/project/pc/out"
)

func main() {

	o := out.NewOut()

	go o.OutPut()
	//one_one.Exec()
	//one_many.Exec()
	//many_one.Exec()

	many_many.Exec()

	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig
}
