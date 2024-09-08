package main

// import (
// 	"fmt"
// )

// func BinarySearch(arr []int, target int, start int, end int) int {
// 	for start <= end {
// 		mid := start + (end-start)/2
// 		if arr[mid] == target {
// 			fmt.Println("Bingo, find ya !!! ", mid)
// 			return mid
// 		}
// 		if arr[mid] < target {
// 			start = mid + 1
// 		} else {
// 			end = mid - 1
// 		}
// 	}
// 	return -1
// }
// func SearchInfinite(arr []int, target int) int {
// 	// Lets do Inital Ranges
// 	start := 0
// 	end := 1
// 	for arr[end] < target {
// 		newstart := end + 1
// 		newend := end + (end-start+1)*2
// 		start = newstart
// 		end = newend
// 	}
// 	return BinarySearch(arr, target, start, end)
// }
// func main() {
// 	arr := []int{1, 3, 5, 8, 9, 9, 12, 14, 16, 21, 27}
// 	target := 9
// 	fmt.Println("Search Result: ", SearchInfinite(arr, target))
// }
