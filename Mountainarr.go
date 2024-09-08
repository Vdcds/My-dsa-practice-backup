package main

// import (
// 	"fmt"
// )

// // A Mountain arr is like this: arr:=int[]{0,2,1,0} where the peak is 2 which is index 1.
// func FindPeakElement(arr []int) int {
// 	start, end := 0, len(arr)-1
// 	for start < end {
// 		mid := start + (end-start)/2
// 		if arr[mid] < arr[mid+1] {
// 			start = mid + 1
// 		} else {
// 			end = mid
// 		}
// 	}
// 	return end
// }
// func main() {
// 	arr := []int{0, 2, 1, 0}
// 	fmt.Println("Index of the peak element is:", FindPeakElement(arr))
// }
