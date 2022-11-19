package problem1629

/*
A newly designed keypad was tested, where a tester pressed a sequence of n keys, one at a time.
You are given a string keysPressed of length n, where keysPressed[i] was the ith key pressed in the testing sequence, and a sorted list releaseTimes,
where releaseTimes[i] was the time the ith key was released.
Both arrays are 0-indexed. The 0th key was pressed at the time 0, and every subsequent key was pressed at the exact time the previous key was released.
The tester wants to know the key of the keypress that had the longest duration.
The ith keypress had a duration of releaseTimes[i] - releaseTimes[i - 1], and the 0th keypress had a duration of releaseTimes[0].
Note that the same key could have been pressed multiple times during the test, and these multiple presses of the same key may not have had the same duration.
Return the key of the keypress that had the longest duration.
If there are multiple such keypresses, return the lexicographically largest key of the keypresses.
*/

func slowestKey(releaseTimes []int, keysPressed string) byte {
	var maxTime int = releaseTimes[0] // The first key took releaseTimes[0] ms to release
	var maxIdx, newTime int
	for i := 1; i < len(releaseTimes); i++ {
		newTime = releaseTimes[i] - releaseTimes[i-1] // Time a key took to release
		if newTime >= maxTime {                       // If this is the slowest
			if newTime == maxTime && keysPressed[maxIdx] > keysPressed[i] {
				// "If there are multiple such keypresses, return the lexicographically largest key of the keypresses.""
				// We skip if both are the same time and current is larger lexicographically
				continue
			}
			// Make current time the new slowest
			maxTime = releaseTimes[i] - releaseTimes[i-1]
			maxIdx = i
		}
	}
	return keysPressed[maxIdx]
}
