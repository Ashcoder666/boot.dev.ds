package ch1

import (
	"math"
	"slices"
)

func Find_minimum(nums []int) (int, int) {
	leastval := nums[0]
	if len(nums) == 0 {
		return math.MaxInt, math.MaxInt
	}
	for _, num := range nums {

		if leastval > num {
			leastval = num
		}
	}

	builtIn := slices.Min(nums)

	return leastval, builtIn
}

func FindMedian(arr []int) int {
	var median int
	if len(arr)%2 == 0 {
		median = (arr[len(arr)/2] + arr[len(arr)/2-1]) / 2
	} else {
		median = arr[len(arr)/2]
	}

	return median

}
