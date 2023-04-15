package main

import (
	"fmt"
	"math/rand"
	"os"
	"sort"
	"sync"
	"time"
)

var wg sync.WaitGroup
var max = 30000

func concurrentSorting() {
	rand.Seed(420)

	//Create array to average time for each algorithm
	times1 := make([]time.Duration, numRuns)
	times2 := make([]time.Duration, numRuns)

	//Algorithm 1: Merge Sort
	fmt.Println("Sorting", numbers, "numbers with", alg1, "...")
	for i := 0; i < numRuns; i++ {
		//Generate slice of n random numbers
		slice := generateSlice(numbers)
		startTime := time.Now()

		//Sort the array using given sort algorithm
		slice = concurrentMergeSort(slice)

		time := time.Since(startTime)

		//Check array has been sorted correctly
		if sort.SliceIsSorted(slice, func(i, j int) bool { return slice[i] <= slice[j] }) {
			fmt.Println("Sorted,", alg1, "functional")
			times1[i] = time
			fmt.Println("Run", i+1, "time:", times1[i])
		} else {
			fmt.Println("Not sorted,", alg1, "not functional")
			os.Exit(1)
		}
	}

	//Algorithm 2: Quick Sort
	fmt.Println("\nSorting", numbers, "numbers with", alg2, "...")
	for i := 0; i < numRuns; i++ {
		slice := generateSlice(numbers)
		startTime := time.Now()

		//Sort array using concurrent quick sort algorithm
		//Add 1 to the WaitGroup counter
		wg.Add(1)
		concurrentQuickSort(slice, 0, numbers-1)
		wg.Wait()

		time := time.Since(startTime)

		//Check array has been sorted
		if sort.SliceIsSorted(slice, func(i, j int) bool { return slice[i] <= slice[j] }) {
			fmt.Println("Sorted,", alg2, "functional")
			times2[i] = time
			fmt.Println("Run", i+1, "time:", times2[i])
		} else {
			fmt.Println("Not sorted,", alg2, "not functional")
			os.Exit(1)
		}
	}

	//Get total runtime for each algorithm
	totalTime1 := time.Duration(0)
	totalTime2 := time.Duration(0)
	for _, v := range times1 {
		totalTime1 += v
	}
	for _, v := range times2 {
		totalTime2 += v
	}

	//Print out the average runtime for each algorithm
	fmt.Println("\nAverage time [", alg1, "]:", totalTime1/numRuns)
	fmt.Println("Average time [", alg2, "]:", totalTime2/numRuns)
}

// Algorithm 1: Merge Sort
func concurrentMergeSort(arr []int64) []int64 {
	if len(arr) < 2 {
		return arr
	}

	if len(arr) <= max {
		return mergeSort(arr)
	} else {
		first := make(chan []int64)

		go func() { first <- concurrentMergeSort(arr[:len(arr)/2]) }()
		second := concurrentMergeSort(arr[len(arr)/2:])
		return merge(<-first, second)
	}
}

// Algorithm 2: Quick Sort
func concurrentQuickSort(arr []int64, low, high int) {
	defer wg.Done()

	if low < high {
		var p int
		wg.Add(2)

		arr, p = partition(arr, low, high)

		go concurrentQuickSort(arr, low, p-1)
		concurrentQuickSort(arr, p+1, high)
	}
}
