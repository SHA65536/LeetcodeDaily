package problem2140

/*
You are given a 0-indexed 2D integer array questions where questions[i] = [pointsi, brainpoweri].
The array describes the questions of an exam, where you have to process the questions in order
(i.e., starting from question 0) and make a decision whether to solve or skip each question.
Solving question i will earn you pointsi points but you will be unable to solve each of the next brainpoweri questions.
If you skip question i, you get to make the decision on the next question.
    For example, given questions = [[3, 2], [4, 3], [4, 4], [2, 5]]:
    If question 0 is solved, you will earn 3 points but you will be unable to solve questions 1 and 2.
    If instead, question 0 is skipped and question 1 is solved, you will earn 4 points but you will be unable to solve questions 2 and 3.
Return the maximum points you can earn for the exam.
*/

func mostPoints(questions [][]int) int64 {
	// dp[i] represents the number of points you can make
	// from the ith index
	var dp = make([]int64, len(questions))

	// helper(i) solves for dp(i)
	var helper func(int) int64

	helper = func(i int) int64 {
		var solve, skip int64
		// Out of bounds
		if i >= len(dp) {
			return 0
		}
		// Already calculated
		if dp[i] != 0 {
			return dp[i]
		}
		// Value if we solve this question (current value + next solvable)
		solve = int64(questions[i][0]) + helper(i+questions[i][1]+1)
		// Value if we skip this question (next question value)
		skip = helper(i + 1)
		dp[i] = max(solve, skip)
		return dp[i]
	}

	return helper(0)
}

func mostPointsIterative(questions [][]int) int64 {
	var dp = make([]int64, len(questions)+1)
	for i := len(questions) - 1; i >= 0; i-- {
		var points, jump = questions[i][0], questions[i][1]
		dp[i] = max(int64(points)+dp[min(jump+i+1, len(questions))], dp[i+1])
	}
	return dp[0]
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
