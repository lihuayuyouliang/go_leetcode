package task_worker

import (
	"fmt"
	"sync"
)

type task struct {
	begin  int
	end    int
	result chan<- int
}

func (t *task) do() {
	sum := 0
	for i := t.begin; i <= t.end; i++ {
		sum += i
	}
	t.result <- sum
}
func MainWorker() {
	taskChan := make(chan task, 10)
	resultChan := make(chan int, 10)
	wg := &sync.WaitGroup{}
	go InitTask(taskChan, resultChan, 100)
	go DistributeTask(taskChan, wg, resultChan)
	sum := ProcessResult(resultChan)
	fmt.Println("sum:", sum)
}
func InitTask(taskChan chan<- task, r chan int, p int) {
	qu := p / 10
	mod := p % 10
	high := qu * 10
	for i := 0; i < qu; i++ {
		b := 10*i + 1
		e := 10 * (i + 1)
		tsk := task{
			begin:  b,
			end:    e,
			result: r,
		}
		taskChan <- tsk
	}
	if mod != 0 {
		tsk := task{
			begin:  high + 1,
			end:    p,
			result: r,
		}
		taskChan <- tsk
	}
	close(taskChan)
}
func DistributeTask(taskChan <-chan task, wg *sync.WaitGroup, resultChan chan int) {
	for v := range taskChan {
		wg.Add(1)
		go ProcessTask(v, wg)
	}
	wg.Wait()
	close(resultChan)
}
func ProcessTask(t task, wg *sync.WaitGroup) {
	t.do()
	wg.Done()
}
func ProcessResult(resultChan chan int) int {
	sum := 0
	for v := range resultChan {
		sum += v
	}
	return sum
}
