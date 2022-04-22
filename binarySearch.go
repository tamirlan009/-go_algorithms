package main

func BinarySearch(arr []int, number int) int {
	low := 0
	high := len(arr) - 1
	for low <= high {
		mid := (low + high) / 2
		if arr[mid] == number {
			return mid
		}
		if arr[mid] > number {
			high = arr[mid] - 1
		}
		if arr[mid] < number {
			low = arr[mid] + 1
		}
	}
	return 0
}
