package many_many

import (
	"fmt"
	"time"

	"github.com/cncamp/golang/project/pc/out"
)

type Task struct {
	ID int64
}

func (task *Task) run() {
	out.Println(task.ID)
}

var taskCh = make(chan Task, 10)
var done = make(chan struct{})
var taskNums int64 = 10000

func producer(pr chan<- Task, do chan struct{}) {
	var i int64
	for {
		if i == taskNums {
			i = 0
		}
		i++

		task := Task{
			ID: i,
		}

		select {
		case pr <- task:
		//如果关闭 done channel 那么这个 case 必定会被执行
		case <-do:
			out.Println("生产者退出")
			return
		}
	}
}

func consumer(co <-chan Task, do chan struct{}) {
	for {
		select {
		case t := <-co:
			if t.ID != 0 {
				t.run()
			}
		//如果关闭 done channel 这个 case 必定会被执行
		case <-do:
			for n := range taskCh {
				out.Println(n.ID)
			}
			out.Println("消费者退出")
			return
		}
	}
}

func Exec() {
	go producer(taskCh, done)
	go producer(taskCh, done)
	go producer(taskCh, done)
	go producer(taskCh, done)

	go consumer(taskCh, done)
	go consumer(taskCh, done)
	go consumer(taskCh, done)
	go consumer(taskCh, done)

	time.Sleep(time.Second * 5)
	close(done)
	close(taskCh)
	time.Sleep(time.Second * 5)
	fmt.Println(len(taskCh))
}
