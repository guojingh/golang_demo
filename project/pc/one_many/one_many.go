package one_many

import (
	"sync"

	"github.com/cncamp/golang/project/pc/out"
)

type Task struct {
	ID int64
}

func (task *Task) run() {
	out.Println(task.ID)
}

var taskNum int64 = 10000
var taskCh = make(chan Task, 10)

func producer(pc chan<- Task) {
	var num int64
	for num = 0; num < taskNum; num++ {
		task := Task{
			ID: num,
		}
		pc <- task
	}
	close(pc)
}

func consumer(cc <-chan Task) {
	for task := range cc {
		if task.ID != 0 {
			task.run()
		}
	}
}

func Exec() {
	wg := &sync.WaitGroup{}
	wg.Add(1)

	//一个生产者
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		producer(taskCh)
	}(wg)

	//多个消费者
	var num int64
	for num = 0; num < taskNum; num++ {
		if num%100 == 10 {
			wg.Add(1)
			go func(wg *sync.WaitGroup) {
				defer wg.Done()
				consumer(taskCh)
			}(wg)
		}
	}

	wg.Wait()
	out.Println("一对多协程打印结束")
}
