package problem0997

/*
n a town, there are n people labeled from 1 to n. There is a rumor that one of these people is secretly the town judge.
If the town judge exists, then:
    The town judge trusts nobody.
    Everybody (except for the town judge) trusts the town judge.
    There is exactly one person that satisfies properties 1 and 2.
You are given an array trust where trust[i] = [ai, bi] representing that the person labeled ai trusts the person labeled bi.
Return the label of the town judge if the town judge exists and can be identified, or return -1 otherwise.
*/

func findJudge(n int, trust [][]int) int {
	var trusts = make([]int, n)
	var trusted = make([]int, n)
	for _, e := range trust {
		trusts[e[0]-1]++
		trusted[e[1]-1]++
	}
	for i := range trusts {
		if trusts[i] == 0 && trusted[i] == n-1 {
			return i + 1
		}
	}
	return -1
}
