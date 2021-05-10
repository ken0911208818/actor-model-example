package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func main() {
	testdata := generateRandomSlice(100)
	qsortGood(testdata[50:])
	qsortGood(testdata[:50])
}

func generateRandomSlice(n int) []int {
	slice := make([]int, n, n)
	for i := 0; i < n; i++ {
		slice[i] = int(rand.Int31())
	}
	return slice
}

func qsortGoodWorker(inputCh chan int, wg *sync.WaitGroup, remainingTaskNum *sync.WaitGroup) {
	defer wg.Done()

	for input := range inputCh {
		// do something
		s := fmt.Sprintf("Job is : %v", input)
		fmt.Println(s)
		time.Sleep(1 * time.Second)
		remainingTaskNum.Done()
	}
}

// WARNING: this qsortGood is for demo only, not for production usage.
// The actual performance of qsortGood is MUCH worse than the standard library
func qsortGood(input []int) {
	wg := sync.WaitGroup{}
	remainingTaskNum := sync.WaitGroup{}

	threadNum := runtime.NumCPU() * 2
	inputCh := make(chan int)
	wg.Add(threadNum)
	for i := 0; i < threadNum; i++ {
		fmt.Printf("start Workder: %v\n", i)
		go qsortGoodWorker(inputCh, &wg, &remainingTaskNum)
	}
	for _, v := range input {
		// add the input to channel, and wait for all subtask completed
		remainingTaskNum.Add(1)
		inputCh <- v
	}
	remainingTaskNum.Wait()

	// let worker thread die peacefully, we SHOULD NOT leave the worker thread behind
	close(inputCh)
	wg.Wait()
}
