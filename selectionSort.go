package main

func findSmallest(arr []int) int {
	smallestIndex := 0
	smallestValue := arr[smallestIndex]

	for i, v := range arr {
		if v < smallestValue {
			smallestValue = v
			smallestIndex = i
		}
	}

	return smallestIndex
}

func SelectionSort(arr []int) []int {
	sizeArr := len(arr)
	if sizeArr == 0 {
		return arr
	}
	sortedArr := make([]int, sizeArr)
	for i := 0; i < sizeArr; i++ {
		min := findSmallest(arr)
		sortedArr[i] = arr[min]
		arr = append(arr[:min], arr[min+1:]...)

	}
	return sortedArr

}
