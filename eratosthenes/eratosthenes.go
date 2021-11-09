package main

import (
	"fmt"
	"math"
)

const UINT32SIZE = 4_294_967_296 // 2**32

func main() {
	const max int = UINT32SIZE

	// Initialize Data - Using a bitmap
	var xs [max]bool
	xs[0] = false
	xs[1] = false
	for i := 2; i < max; i++ {
		xs[i] = true
	}

	// Eratosthenes
	max_sqrt := int(math.Sqrt(float64(max)))
	//fmt.Printf("Checking from 2 to %d.\n", max_sqrt) // DEBUG

	var multiple int
	for p := 2; p <= max_sqrt; p++ {

		// Skip if it's a multiple
		if !xs[p] {
			// fmt.Printf("Skipping the multiple: %d\n", p) // DEBUG
			continue
		}

		multiple = p + p
		// fmt.Printf("Prime: %d\n", p) // DEBUG
		// fmt.Printf("%d ", p) // DEBUG
		for multiple < max {
			xs[multiple] = false
			// fmt.Printf("Marking %d as false.\n", multiple) // DEBUG
			multiple += p
		}
	}

	// Results
	count := 0
	for j := 0; j < max; j++ {
		if xs[j] {
			fmt.Println(j) // Show results
			count += 1
		}
	}
	// fmt.Printf("\nFound %d primes.\n", count) // DEBUG
}
