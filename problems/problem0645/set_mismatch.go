package problem0645

/*
You have a set of integers s, which originally contains all the numbers from 1 to n.
Unfortunately,due to some error, one of the numbers in s got duplicated to another number in the set,
which results in repetition of one number and loss of another number.
You are given an integer array nums representing the data status of this set after the error.
Find the number that occurs twice and the number that is missing and return them in the form of an array.
*/

func findErrorNums(nums []int) []int {
	var double, sum, expected int
	var freqs = map[int]int{}
	for _, num := range nums {
		sum += num
		if freqs[num] != 0 {
			double = num
		}
		freqs[num]++
	}
	sum -= double
	expected = (len(nums) * (len(nums) + 1)) / 2
	return []int{double, expected - sum}
}
