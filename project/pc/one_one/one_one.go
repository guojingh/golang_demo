package one_one

import (
	"fmt"
	"sync"

	"github.com/cncamp/golang/project/pc/out"
)

type Task struct {
	ID int64
}

func (task *Task) run() {
	out.Println(task.ID)
}

var taskCh = make(chan Task, 10)

const taskNum int64 = 10000

func producer(wo chan<- Task) {
	var i int64
	for i = 1; i < taskNum; i++ {
		t := Task{
			ID: i,
		}

		wo <- t
	}
	close(wo)
}

func consumer(ro <-chan Task) {
	for t := range ro {
		if t.ID != 0 {
			t.run()
		}
	}
}

func Exec() {
	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		producer(taskCh)
		//wg.Done()
	}(wg)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		consumer(taskCh)
		//wg.Done()
	}(wg)

	wg.Wait()
	fmt.Println("一对一生产消费结束")
}
