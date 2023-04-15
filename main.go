package main

import (
	"fmt"
)

// Constants
const numRuns = 3
const numbers = 3000000
const alg1 = "Merge Sort"
const alg2 = "Quick Sort"
const sequential = false
const concurrent = false
const reduction = true

func main() {
	if sequential {
		fmt.Println("-----------------------------------------------------")
		fmt.Println("Comparing Sort Times without Parallelism/Concurrency:")
		fmt.Println("-----------------------------------------------------")

		sequentialSorting()
	}

	if concurrent {
		fmt.Println("\n-----------------------------------------------------")
		fmt.Println("Comparing Sort Times with Parallelism/Concurreny:")
		fmt.Println("-----------------------------------------------------")

		concurrentSorting()
	}

	if reduction {
		fmt.Println("\n-----------------------------------------------------")
		fmt.Println("Effect of Reduction of Cores on Execution Time:")
		fmt.Println("-----------------------------------------------------")

		reductionOfCores()
	}
}
