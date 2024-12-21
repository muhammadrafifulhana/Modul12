package main

import "fmt"

func insertionSort(data []int) {
	for i := 1; i < len(data); i++ {
		key := data[i]
		j := i - 1
		for j >= 0 && data[j] > key {
			data[j+1] = data[j]
			j--
		}
		data[j+1] = key
	}
}

func main() {
	var data []int
	var num int

	for {
		fmt.Scan(&num)
		if num == -5313 {
			break
		}
		if num != 0 {
			data = append(data, num)
		} else {
			insertionSort(data)
			n := len(data)
			if n%2 == 1 {
				fmt.Println(data[n/2])
			} else {
				median := (data[n/2-1] + data[n/2]) / 2
				fmt.Println(median)
			}
		}
	}
}
