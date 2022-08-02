package binsearch

import "fmt"

//Take in a sorted integer array we will call "list" and an int to search for, then return the int position (int)
func binSearch(list []int, findMe int) int {
	// constants of array positions to search between (these two values make a search range)
	low := 0              // arrays always bottom at 0
	high := len(list) - 1 // leng of the array -1 will be the max position

	// loop though array as long as low is less than high
	for low <= high {
		// find the middle position of the sorted array by adding high + low / 2
		mid := (low + high) / 2
		// if the int value of the middle position is = findMe then return mids position
		if list[mid] == findMe {
			return mid
		}
		// if the int value of the middle position is less than findMe, shorten the search range by changing low -> mid
		if list[mid] < findMe {
			low = mid + 1
		}
		// if the int value of the middle position is less than findMe, shorten the search range by changing high -> mid
		if list[mid] > findMe {
			high = mid - 1
		}
	}
	// this will mean the int is not in the sorted list
	return -1
}

func main() {
	fmt.Println(binSearch([]int{1, 2, 3, 3, 5, 9, 19, 33, 88, 99}, 1)) // 0 (position 0 has our searched int)
	fmt.Println(binSearch([]int{1, 2, 3, 4, 22, 44, 54, 88, 99}, -1))  // -1 (not in the data set)
}
