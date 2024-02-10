package main

import "sort"

func TwoNumberSum(array []int, target int) []int {
	sort.Ints(array)
    left := 0
    right := len(array) - 1
    for left <= right {
		currentSum := array[left] + array[right];
        if currentSum == target {
            return []int{array[left], array[right]}
        } else if currentSum < target {
            left += 1
        } else {
            right -= 1
        }
	}
	return []int{}
}

