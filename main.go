package main

import (
	"fmt"
	"leetcode/mysort"
)

func main() {
	arr := []int{3, 2, 5, 6, 9, 1, 4}
	end := mysort.QuickSort(arr)
	fmt.Println(end)
}
