package main

import "fmt"

func main() {
	fmt.Println("------Bubble sort------")
	bubbleArray := []int{36, 25, 48, 12, 25, 65, 43, 57}
	fmt.Println("unordered array:", bubbleArray)
	BubbleSort(bubbleArray)
	fmt.Println("sorted    array:", bubbleArray)
}
