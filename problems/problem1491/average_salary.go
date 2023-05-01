package problem1491

/*
You are given an array of unique integers salary where salary[i] is the salary of the ith employee.
Return the average salary of employees excluding the minimum and maximum salary.
Answers within 10-5 of the actual answer will be accepted.
*/

func average(salary []int) float64 {
	var res float64
	var min, max int = salary[0], salary[0]

	for i := range salary {
		if salary[i] > max {
			max = salary[i]
		}
		if salary[i] < min {
			min = salary[i]
		}
		res += float64(salary[i])
	}

	res -= float64(min + max)

	return res / float64(len(salary)-2)
}
