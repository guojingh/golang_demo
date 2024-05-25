package many_one

import (
	"sync"

	"github.com/cncamp/golang/project/pc/out"
)

type Task struct {
	ID int64
}

const taskNum int64 = 10000

var taskCh = make(chan Task, 10)

func (task *Task) run() {
	out.Println(task.ID)
}

// 生产者
func producer(wo chan<- Task, startNum, nums int64) {
	var i int64
	for i = startNum; i < nums; i++ {
		task := Task{
			ID: i,
		}
		wo <- task
	}
}

func consumer(co <-chan Task) {
	for task := range co {
		if task.ID != 0 {
			task.run()
		}
	}
}

func Exec() {
	wg := sync.WaitGroup{}
	proWg := sync.WaitGroup{}
	wg.Add(1)

	var i int64
	for i = 0; i < taskNum; i += 100 {
		wg.Add(1)
		proWg.Add(1)
		go func(i int64) {
			defer wg.Done()
			defer proWg.Done()
			producer(taskCh, i, 100)
		}(i)
	}

	go func() {
		defer wg.Done()
		consumer(taskCh)
	}()

	proWg.Wait()
	go close(taskCh)
	wg.Wait()
	out.Println("多对一执行完成")

}
