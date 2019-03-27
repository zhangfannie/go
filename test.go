package main

import "sync"
import "sync/atomic"
import "fmt"

const iteration = 1000000

var X, Y, r1, r2 [iteration]uint64

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	go f1()
	go f2()
	wg.Wait()

	for i := 0; i < iteration; i++ {
		if r1[i] == 0 && r2[i] == 0 {
			fmt.Printf("i = %d : r1 = %d : r2 = %d\n", i, r1[i], r2[i])
			panic("broken")
		}
	}
}

// goroutine 1
func f1() {
	for i := 0; i < iteration; i++ {
		atomic.StoreUint64(&X[i], 1)
		r1[i] = Y[i]
	}
	wg.Done()
}

// goroutine 2
func f2() {
	for i := 0; i < iteration; i++ {
		atomic.StoreUint64(&Y[i], 1)
		r2[i] = X[i]
	}
	wg.Done()
}
