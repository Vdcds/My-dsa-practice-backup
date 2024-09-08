package main

// import (
// 	"fmt"
// )

// func firstandlastpos() {
// 	arr := []int{1, 3, 5, 8, 9, 9, 12, 14, 16, 21, 27}
// 	fmt.Println("Original Array:", arr)
// 	target := 9
// 	start := 0
// 	end := len(arr) - 1
// 	firstpos := -1
// 	lastpos := -1
// 	for start <= end {
// 		mid := start + (end-start)/2
// 		// First Pos
// 		if arr[mid] == target {
// 			firstpos = mid
// 			end = mid - 1
// 		} else if arr[mid] < target {
// 			start = mid + 1
// 		} else {
// 			end = mid - 1
// 		}
// 	}
// 	// resetting start and end for last pos
// 	start = 0
// 	end = len(arr) - 1

// 	for start <= end {
// 		// Last Pos
// 		mid := start + (end-start)/2
// 		if arr[mid] == target {
// 			lastpos = mid
// 			start = mid + 1
// 		} else if arr[mid] < target {
// 			start = mid + 1
// 		} else {
// 			end = mid - 1
// 		}
// 	}
// 	fmt.Printf("First Position of %d is %d and Last Position is %d\n", target, firstpos, lastpos)
// }

// func main() {
// 	firstandlastpos()
// }
