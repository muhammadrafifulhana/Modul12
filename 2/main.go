package main

import "fmt"

func SelectionSort(arr []int, ascending bool) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		selectedIndex := i
		for j := i + 1; j < n; j++ {
			if (ascending && arr[j] < arr[selectedIndex]) || (!ascending && arr[j] > arr[selectedIndex]) {
				selectedIndex = j
			}
		}
		arr[i], arr[selectedIndex] = arr[selectedIndex], arr[i]
	}
}

func separateOddEven(arr []int) ([]int, []int) {
	odd := []int{}
	even := []int{}
	for _, v := range arr {
		if v%2 == 0 {
			even = append(even, v)
		} else {
			odd = append(odd, v)
		}
	}
	return odd, even
}

func main() {
	var n int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		var m int
		fmt.Scan(&m)
		arr := make([]int, m)

		for j := 0; j < m; j++ {
			fmt.Scan(&arr[j])
		}

		odd, even := separateOddEven(arr)

		SelectionSort(odd, true)
		SelectionSort(even, false)

		for _, v := range odd {
			fmt.Print(v, " ")
		}

		for _, v := range even {
			fmt.Print(v, " ")
		}

		fmt.Println()
	}
}
