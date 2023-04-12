package main

func BubbleSort(array []int) {
	lLen := len(array)
	for i := 0; i < lLen-1; i++ {
		for j := lLen - 1; j > i; j-- {
			if array[j] < array[j-1] {
				//tmp := array[j]
				array[j], array[j-1] = array[j-1], array[j]
			}
		}
	}
}
