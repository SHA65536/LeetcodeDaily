package problem2483

/*
You are given the customer visit log of a shop represented by a 0-indexed string customers consisting only of characters 'N' and 'Y':
    if the ith character is 'Y', it means that customers come at the ith hour
    whereas 'N' indicates that no customers come at the ith hour.
If the shop closes at the jth hour (0 <= j <= n), the penalty is calculated as follows:
    For every hour when the shop is open and no customers come, the penalty increases by 1.
    For every hour when the shop is closed and customers come, the penalty increases by 1.
Return the earliest hour at which the shop must be closed to incur a minimum penalty.
Note that if a shop closes at the jth hour, it means the shop is closed at the hour j.
*/

func bestClosingTime(customers string) int {
	var curPen int
	// We start from all closed
	for i := range customers {
		if customers[i] == 'Y' {
			// Penalty of being close when a customer arrives
			curPen++
		}
	}

	var minIdx, minPen = -1, curPen

	// Start trying to open
	for i := range customers {
		if customers[i] == 'Y' {
			curPen--
		} else {
			curPen++
		}

		// Keep track of minimal penalty
		if curPen < minPen {
			minPen = curPen
			minIdx = i
		}
	}
	return minIdx + 1
}
