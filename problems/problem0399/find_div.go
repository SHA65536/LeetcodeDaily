package problem0399

/*
You are given an array of variable pairs equations and an array of real numbers values,
where equations[i] = [Ai, Bi] and values[i] represent the equation Ai / Bi = values[i].
Each Ai or Bi is a string that represents a single variable.
You are also given some queries, where queries[j] = [Cj, Dj] represents the jth query where
you must find the answer for Cj / Dj = ?.
Return the answers to all queries. If a single answer cannot be determined, return -1.0.
Note: The input is always valid.
You may assume that evaluating the queries will not result in division by zero and that there is no contradiction.
*/

type Relation struct {
	// Target variable
	Target string
	// Ratio between the variables
	Value float64
}

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	var res = make([]float64, len(queries))
	// relations[i] is a list of relations between variable i and others
	var relations = MakeRelations(equations, values)

	for i := range queries {
		// Don't bother searching thigns we don't know about
		if _, ok := relations[queries[i][0]]; !ok {
			res[i] = -1
			continue
		}
		if _, ok := relations[queries[i][1]]; !ok {
			res[i] = -1
			continue
		}
		// Get solution for one equation
		res[i] = SolveEquation(relations, queries[i][0], queries[i][1])
	}

	return res
}

func SolveEquation(graph map[string][]Relation, start, end string) float64 {
	// seen [i] means we already checked this node
	var seen = map[string]bool{}
	// dfs(i, j) returns i / j if calculable, otherwise false
	var dfs func(start, end string) (float64, bool)

	dfs = func(start, end string) (float64, bool) {
		// If we found the target return 1 (x/x = 1)
		if start == end {
			return 1, true
		}
		// Mark current node
		seen[start] = true
		// Search each connected node
		for _, rel := range graph[start] {
			// Ignore marked nodes
			if seen[rel.Target] {
				continue
			}
			// Recurse on node, and if solution found
			// return it times the current relation value
			if val, ok := dfs(rel.Target, end); ok {
				return val * rel.Value, true
			}
		}
		return -1, false
	}

	// Start the dfs
	if val, ok := dfs(start, end); ok {
		return val
	}
	return -1
}

func MakeRelations(equations [][]string, values []float64) map[string][]Relation {
	var Relations = map[string][]Relation{}
	for i, eq := range equations {
		Relations[eq[0]] = append(Relations[eq[0]], Relation{
			Target: eq[1],
			Value:  values[i],
		})
		Relations[eq[1]] = append(Relations[eq[1]], Relation{
			Target: eq[0],
			Value:  1 / values[i],
		})
	}
	return Relations
}
