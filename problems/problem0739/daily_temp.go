package problem0739

/*
Given an array of integers temperatures represents the daily temperatures,
return an array answer such that answer[i] is the number of days you have to wait after the ith day to get a warmer temperature.
If there is no future day for which this is possible, keep answer[i] == 0 instead.
*/

func dailyTemperatures(temperatures []int) []int {
	var res = make([]int, len(temperatures))
	for i := len(res) - 2; i >= 0; i-- {
		for j := 1; i+j < len(temperatures); j += res[i+j] {
			if temperatures[i] < temperatures[i+j] {
				res[i] = j
				break
			}
			if res[i+j] == 0 {
				break
			}
		}
	}
	return res
}
