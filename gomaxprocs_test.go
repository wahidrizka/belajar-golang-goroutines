package belajar_golang_goroutines

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestGoMaxProcs(t *testing.T) {
	group := sync.WaitGroup{}

	for i := 1; i <= 100; i++ {
		group.Go(func() {
			time.Sleep(3 * time.Second)
		})
	}
	totalCPU := runtime.NumCPU()
	fmt.Println("Total CPU:", totalCPU)

	runtime.GOMAXPROCS(20)
	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("Total Thread:", totalThread)

	totalGoroutine := runtime.NumGoroutine()
	fmt.Println("Total Goroutine:", totalGoroutine)
}
