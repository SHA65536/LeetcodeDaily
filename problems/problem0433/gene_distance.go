package problem0433

/*
A gene string can be represented by an 8-character long string, with choices from 'A', 'C', 'G', and 'T'.
Suppose we need to investigate a mutation from a gene string start to a gene string end where one mutation is
defined as one single character changed in the gene string.
For example, "AACCGGTT" --> "AACCGGTA" is one mutation.
There is also a gene bank bank that records all the valid gene mutations. A gene must be in bank to make it a valid gene string.
Given the two gene strings start and end and the gene bank bank, return the minimum number of mutations needed to mutate from start to end.
If there is no such a mutation, return -1.
Note that the starting point is assumed to be valid, so it might not be included in the bank.
*/

func minMutation(start string, end string, bank []string) int {
	var res int
	// Available mutations that we haven't tried yet
	// The end word has value of true
	var gene_dict = map[string]bool{}
	var cur = []string{start}
	for i := range bank {
		gene_dict[bank[i]] = false
	}
	// If the desired end is not in the bank, exit
	if _, ok := gene_dict[end]; !ok {
		return -1
	}
	gene_dict[end] = true
	// BFS for mutation
	for len(cur) > 0 {
		res++
		// Next layer in the BFS
		var next = make([]string, 0, len(cur))
		for i := range cur {
			// Finding if any of the current layer
			// matches genes from the bank
			for k, finish := range gene_dict {
				if canMutate(cur[i], k) {
					if finish {
						// If we reached the end, return
						return res
					}
					// Add gene to the next layer
					// removed it from possible genes
					next = append(next, k)
					delete(gene_dict, k)
				}
			}
		}
		cur = next
	}
	return -1
}

func canMutate(a, b string) bool {
	var faults int
	for i := 0; i < len(a) && faults < 2; i++ {
		if a[i] != b[i] {
			faults++
		}
	}
	return faults == 1
}
