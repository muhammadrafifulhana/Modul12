package main

import "fmt"

func selectionSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}
}

func main() {
	var n int
	fmt.Scan(&n)

	var results [][]int

	for i := 0; i < n; i++ {
		var m int
		fmt.Scan(&m)

		houses := make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Scan(&houses[j])
		}

		selectionSort(houses)
		results = append(results, houses)
	}

	fmt.Print("\n")

	for _, houses := range results {
		for i, house := range houses {
			if i > 0 {
				fmt.Print(" ")
			}
			fmt.Print(house)
		}
		fmt.Println()
	}
}
