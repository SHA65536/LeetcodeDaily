package problem1125

/*
In a project, you have a list of required skills req_skills, and a list of people.
The ith person people[i] contains a list of skills that the person has.
Consider a sufficient team: a set of people such that for every required skill in req_skills, there is at least one person in the team who has that skill.
We can represent these teams by the index of each person.
For example, team = [0, 1, 3] represents the people with skills people[0], people[1], and people[3].
Return any sufficient team of the smallest possible size, represented by the index of each person.
You may return the answer in any order.
It is guaranteed an answer exists.
*/

func smallestSufficientTeam(req_skills []string, people [][]string) []int {
	// dp[skills-set] is a team that satisfies skills-set
	var dp = make(map[int][]int, 1<<len(req_skills))
	dp[0] = []int{}
	// skills[x] is the serial number of the x skill
	var skills = make(map[string]int, len(req_skills))
	for i := range req_skills {
		skills[req_skills[i]] = i
	}
	for i := range people {
		// skill is the bitmap of the current employee skill
		var skill int
		for j := range people[i] {
			skill |= 1 << skills[people[i][j]]
		}
		// for each skill-set in our dp
		for skill_set, team := range dp {
			// try adding the person to the team
			var comb = skill_set | skill
			// if this is a new skill-set or if we can make this skill-set with fewer people
			if val, ok := dp[comb]; !ok || len(val) > 1+len(dp[skill_set]) {
				// update the new skill-set
				dp[comb] = make([]int, len(team))
				copy(dp[comb], team)
				dp[comb] = append(dp[comb], i)
			}
		}
	}
	// 1<<len(req_skills))-1 is a skill-set with all the required skills
	return dp[(1<<len(req_skills))-1]
}
