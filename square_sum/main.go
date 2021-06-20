package main

import (
	"fmt"
	"sync"
	"time"
)

func sequential(n int) {
	// 0*0 + 1*1 + 2*2 + ... * (n-1) * (n-1)
	for i := 0; i < n; i++ {
		func() {
			sum := int64(0)
			for j := 0; j <= i; j++ {
				sum += int64(j * j)
			}
		}()
	}
}

func parallel(n int) {
	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(m int) {
			sum := int64(0)
			for j := 0; j <= m; j++ {
				sum += int64(j * j)
			}
			wg.Done()
		}(i)
	}

	wg.Wait()
}

func main() {
	initial := time.Now()
	sequential(10000)
	finish := time.Now()
	elapsed := finish.Sub(initial)
	fmt.Println(elapsed.Seconds())
}
