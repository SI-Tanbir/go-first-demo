package main

import "fmt"

// Return indices of duplicate elements
func duplicateArray(num []int) []int {
	duplicate := []int{}
	for i := 0; i < len(num)-1; i++ {
		if num[i] == num[i+1] {
			duplicate = append(duplicate, i+1) // store index of duplicate element
		}
	}
	return duplicate
}

// Remove elements at given indices from the slice
func removeIndices(num []int, indices []int) []int {
	toRemove := make(map[int]bool)
	for _, idx := range indices {
		toRemove[idx] = true
	}
	result := []int{}
	for i, v := range num {
		if !toRemove[i] {
			result = append(result, v)
		}
	}
	return result
}

func main() {
	num := []int{-5, -5, -3, -3, -1, -1, 0, 1, 1, 2}
	fmt.Println("Original:", num)

	indicesToRemove := duplicateArray(num)
	fmt.Println("Duplicate indices:", indicesToRemove)

	numNoDuplicates := removeIndices(num, indicesToRemove)
	fmt.Println("After removing duplicates:", numNoDuplicates)
}
