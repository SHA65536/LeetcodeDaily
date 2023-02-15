package problem1502

import "sort"

/*
A sequence of numbers is called an arithmetic progression if the difference between any two consecutive elements is the same.
Given an array of numbers arr, return true if the array can be rearranged to form an arithmetic progression. Otherwise, return false.
*/

func canMakeArithmeticProgression(arr []int) bool {
	sort.Ints(arr)
	var diff = arr[0] - arr[1]
	for i := 1; i < len(arr); i++ {
		if arr[i-1]-arr[i] != diff {
			return false
		}
	}
	return true
}

func canMakeArithmeticProgressionOpt(arr []int) bool {
	var min, max = arr[0], arr[0]
	var dis, pos int
	for i := range arr {
		if arr[i] > max {
			max = arr[i]
		}
		if arr[i] < min {
			min = arr[i]
		}
	}
	if (max-min)%(len(arr)-1) != 0 {
		return false
	}
	dis = (max - min) / (len(arr) - 1)
	for i := 0; i < len(arr); {
		if arr[i] == min+i*dis {
			i++
		} else if (arr[i]-min)%dis != 0 {
			return false
		} else {
			pos = (arr[i] - min) / dis
			if pos < i || arr[pos] == arr[i] {
				return false
			}
			arr[i], arr[pos] = arr[pos], arr[i]
		}
	}
	return true
}
