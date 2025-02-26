package main

import (
	"fmt"
	"sync"
)




func main() {

	var score = []int{0}

	wg := &sync.WaitGroup{} //usual practice to use pointer to waitgroup
	mut := &sync.Mutex{} // can also use RWMutex to allow multiple readers but only one writer using RLock and RUnlock, it's just a performance optimization thing otherwise normal mutex also does the job but when there are multiple readers  it's better to use RWMutex.
	wg.Add(3) //approach 2 - no difference
 	// wg.Add(1) -- approach 1
	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		fmt.Println("First goroutine")
		m.Lock()
		score = append(score, 1)
		m.Unlock()
		wg.Done() 
	}(wg,mut) //iife - immediately invoked function expression
	// wg.Add(1)
	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		fmt.Println("Second goroutine")
		m.Lock()
		score = append(score, 2)
		m.Unlock()
		wg.Done()
	}(wg,mut)
	// wg.Add(1)
	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		fmt.Println("Third goroutine")
		m.Lock()
		score = append(score, 3)
		m.Unlock()
		wg.Done()
	}(wg,mut)

	wg.Wait() //wait for all goroutines to finish - but without mutex we tend to go in race condition

	fmt.Println("Final score is", score) //order of execution is not guaranteed 
}
