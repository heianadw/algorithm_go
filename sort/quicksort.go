package main

func sort(array []int, low, high int) {
	if low > high {
		return
	}
	i := low
	j := high
	index := array[i]
	for i < j {
		for i < j && array[j] >= index {
			j--
		}
		if i < j {
			array[i] = array[j]
			i++
		}
		for i < j && array[i] < index {
			i++
		}
		if i < j {
			array[j] = array[i]
			j--
		}
	}
	array[i] = index
	sort(array, low, i-1)
	sort(array, i+1, high)
}

func QuickSort(array []int) {
	sort(array, 0, len(array)-1)
}
