package main

func QuickSort(arr []int) []int {

	if len(arr) < 2 {
		return arr
	}
	pivot := arr[0]
	less := make([]int, 0)
	greater := make([]int, 0)

	for _, value := range arr[1:] {
		if value > pivot {
			greater = append(greater, value)
		}
		if value < pivot {
			less = append(less, value)
		}
	}
	arr = append(QuickSort(less), pivot)
	arr = append(arr, QuickSort(greater)...)

	return arr
}
