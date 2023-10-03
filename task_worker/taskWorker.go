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

// 每个任务执行
func (t *task) do() {
	sum := 0
	for i := t.begin; i <= t.end; i++ {
		sum += i
	}
	t.result <- sum
}

// MainWorker 主处理函数
func MainWorker() {
	taskChan := make(chan task, 10)
	resultChan := make(chan int, 10)
	wg := &sync.WaitGroup{}
	go InitTask(taskChan, resultChan, 100)
	go DistributeTask(taskChan, wg, resultChan)
	sum := ProcessResult(resultChan)
	fmt.Println("sum:", sum)
}

// InitTask 任务初始化
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

// DistributeTask 读取并处理任务
func DistributeTask(taskChan <-chan task, wg *sync.WaitGroup, resultChan chan int) {
	for v := range taskChan {
		wg.Add(1)
		go ProcessTask(v, wg)
	}
	wg.Wait()
	close(resultChan)
}

// ProcessTask 任务处理
func ProcessTask(t task, wg *sync.WaitGroup) {
	t.do()
	wg.Done()
}

// ProcessResult 合并结果
func ProcessResult(resultChan chan int) int {
	sum := 0
	for v := range resultChan {
		sum += v
	}
	return sum
}
