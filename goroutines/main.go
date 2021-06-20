package main

import (
	"fmt"
	"sync"
)

func fatorial(n int, group *sync.WaitGroup) {
	prod := 1
	for i := 1; i <= n; i++ {
		prod *= i
	}

	fmt.Println(prod)
	group.Done()
}

func main() {
	var group, group2 sync.WaitGroup

	group.Add(3)
	group2.Add(1)
	go fatorial(5, &group)
	go fatorial(4, &group)
	go fatorial(3, &group)
	go fatorial(2, &group)
	group.Wait()
}
