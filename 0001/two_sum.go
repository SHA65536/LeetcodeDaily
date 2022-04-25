package two_sum

// twoSum recieves an array of numbers, and a target number
// returns the indices of two numbers that add up to target
func twoSum(nums []int, target int) []int {
	var seen = map[int]int{}
	for idx, num := range nums {
		compliment := target - num
		if compIdx, ok := seen[compliment]; ok {
			return []int{compIdx, idx}
		} else {
			seen[num] = idx
		}
	}
	return []int{}
}
