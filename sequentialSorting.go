package main

import (
	"fmt"
	"math/rand"
	"os"
	"sort"
	"time"
)

func sequentialSorting() {
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

		//Sort the array using merge sort algorithm
		slice = mergeSort(slice)

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
		//Generate slice of n random numbers
		slice := generateSlice(numbers)
		startTime := time.Now()

		//Sort the array using quick sort algorithm
		slice = quickSort(slice, 0, numbers-1)

		time := time.Since(startTime)

		//Check array has been sorted correctly
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

// Generates a slice of size n, size filled with random positive 64bit numbers
func generateSlice(size int) []int64 {
	slice := make([]int64, size)
	for i := 0; i < size; i++ {
		slice[i] = rand.Int63()
	}
	return slice
}

// Algorithm 1: Merge Sort
func mergeSort(arr []int64) []int64 {
	if len(arr) < 2 {
		return arr
	}
	first := mergeSort(arr[:len(arr)/2])
	second := mergeSort(arr[len(arr)/2:])
	return merge(first, second)
}

// Algorithm 2: Quick Sort
func quickSort(arr []int64, low, high int) []int64 {
	if low < high {
		var p int
		arr, p = partition(arr, low, high)
		arr = quickSort(arr, low, p-1)
		arr = quickSort(arr, p+1, high)
	}
	return arr
}

// Merge function to merge 2 sorted arrays into a single sorted array
func merge(a []int64, b []int64) []int64 {
	final := []int64{}
	i := 0
	j := 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			final = append(final, a[i])
			i++
		} else {
			final = append(final, b[j])
			j++
		}
	}
	for ; i < len(a); i++ {
		final = append(final, a[i])
	}
	for ; j < len(b); j++ {
		final = append(final, b[j])
	}
	return final
}

// Partition function to find pivot and move everything to the correct side of the pivot
func partition(arr []int64, low, high int) ([]int64, int) {
	pivot := arr[high]
	i := low
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i]
	return arr, i
}
