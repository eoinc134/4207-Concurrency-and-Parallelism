package main

import (
	"fmt"
	"math/rand"
	"os"
	"runtime"
	"sort"
	"time"
)

var rocNumRuns = runtime.GOMAXPROCS(0)
var numCores = runtime.GOMAXPROCS(0)

func reductionOfCores() {
	//Constants
	const alg = "Quick Sort"
	rand.Seed(420)

	//Create array to average time
	times := make([]time.Duration, rocNumRuns)

	//Algorithm: Quick Sort
	fmt.Println("Sorting", numbers, "numbers with", alg, "...")
	for i := 0; i < rocNumRuns; i++ {
		fmt.Printf("\nNumber of CPUs is %d\n", runtime.GOMAXPROCS(0))

		slice := generateSlice(numbers)
		startTime := time.Now()

		wg.Add(1)
		concurrentQuickSort(slice, 0, numbers-1)
		wg.Wait()

		time := time.Since(startTime)

		if sort.SliceIsSorted(slice, func(i, j int) bool { return slice[i] <= slice[j] }) {
			//fmt.Println("Sorted,", alg, "functional")
			times[i] = time
			fmt.Println("Run", i+1, "time:", times[i])
		} else {
			//fmt.Println("Not sorted,", alg, "not functional")
			os.Exit(1)
		}
		numCores--
		runtime.GOMAXPROCS(numCores)
	}
}
