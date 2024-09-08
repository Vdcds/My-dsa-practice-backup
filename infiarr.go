package main

// import (
// 	"fmt"
// )

// func BinarySearch(nums []int, target, start int, end int) int {
// 	for start <= end {
// 		mid := start + (end-start)/2
// 		if nums[mid] == target {
// 			return mid
// 		} else if nums[mid] < target {
// 			start = mid + 1
// 		} else {
// 			end = mid - 1
// 		}
// 	}
// 	return -1
// }

// func searchInInfiniteArray(nums []int, target int) int {
// 	start := 0
// 	end := 1

// 	// Find a range where target might exist
// 	for nums[end] < target {
// 		newStart := end + 1
// 		end = end + (end-start+1)*2
// 		start = newStart
// 	}

// 	// Binary search in the found range
// 	return BinarySearch(nums, target, start, end)
// }

// func main() {
// 	// a 20 element array of even numbers, starting from 0
// 	nums := make([]int, 20)

// 	for i := range nums {
// 		nums[i] = i * 2 // Even numbers
// 	}

// 	target := 24
// 	result := searchInInfiniteArray(nums, target)
// 	fmt.Printf("Target %d found at index: %d\n", target, result)
// }
