package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestMutexForSolveRaceCondition(t *testing.T) {
	var mutex sync.Mutex
	x := 0
	for i := 1; i <= 1000; i++ {
		go func() {
			for j := 0; j <= 100; j++ {
				mutex.Lock()
				x = x + 1
				mutex.Unlock()
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Counter: ", x)
}
