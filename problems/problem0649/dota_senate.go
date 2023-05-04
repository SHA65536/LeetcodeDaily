package problem0649

/*
In the world of Dota2, there are two parties: the Radiant and the Dire.

The Dota2 senate consists of senators coming from two parties.
Now the Senate wants to decide on a change in the Dota2 game. The voting for this change is a round-based procedure.
In each round, each senator can exercise one of the two rights:

Ban one senator's right: A senator can make another senator lose all his rights in this and all the following rounds.
Announce the victory: If this senator found the senators who still have rights to vote are all from the same party,
he can announce the victory and decide on the change in the game.

Given a string senate representing each senator's party belonging.
The character 'R' and 'D' represent the Radiant party and the Dire party.
Then if there are n senators, the size of the given string will be n.

The round-based procedure starts from the first senator to the last senator in the given order.
This procedure will last until the end of voting.
All the senators who have lost their rights will be skipped during the procedure.

Suppose every senator is smart enough and will play the best strategy for his own party.
Predict which party will finally announce the victory and change the Dota2 game. The output should be "Radiant" or "Dire".
*/

func predictPartyVictory(senate string) string {
	var totalR, totalD int
	var input = []byte(senate)

	// Counting the senate
	for i := range input {
		if input[i] == 'R' {
			totalR++
		} else {
			totalD++
		}
	}

	// rCur and wCur are the Reading Cursor and Writing cursor
	var rCur, wCur int
	// rVoted and dVoted are the number of radiants and dires to be voted off
	var rVoted, dVoted int

	// Loop until won
	for totalR > 0 && totalD > 0 {
		wCur = 0
		// Loop over voting round
		for rCur = 0; rCur < len(input); rCur++ {
			if input[rCur] == 'R' {
				// If Radiant can vote
				if rVoted == 0 {
					// Write radiant for next loop
					input[wCur] = 'R'
					// Advance writer cur
					wCur++
					// Add a dire to be removed
					dVoted++
					// Decrease total dires
					totalD--
				} else {
					// If radiant can't vote, decrease radiants
					// to be removed
					rVoted--
				}
			} else {
				// Do the same for dire
				if dVoted == 0 {
					input[wCur] = 'D'
					wCur++
					rVoted++
					totalR--
				} else {
					dVoted--
				}
			}
		}
		// Update vote length
		input = input[:wCur]
	}

	if totalD <= 0 {
		return "Radiant"
	}
	return "Dire"
}
