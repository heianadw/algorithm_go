package main

import "fmt"

func main() {
	fmt.Println("------Bubble sort------")
	bubbleArray := []int{36, 25, 48, 12, 25, 65, 43, 57}
	fmt.Println("unordered array:", bubbleArray)
	BubbleSort(bubbleArray)
	fmt.Println("sorted    array:", bubbleArray)
	fmt.Println("------Quick sort------")
	quickArray := []int{5, 4, 9, 8, 7, 6, 0, 1, 3, 2}
	fmt.Println("unordered array:", quickArray)
	QuickSort(quickArray)
	fmt.Println("sorted    array:", quickArray)
}
