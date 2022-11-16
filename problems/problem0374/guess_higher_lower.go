package problem0374

/*
We are playing the Guess Game. The game is as follows:
I pick a number from 1 to n. You have to guess which number I picked.
Every time you guess wrong, I will tell you whether the number I picked is higher or lower than your guess.
You call a pre-defined API int guess(int num), which returns three possible results:
    -1: Your guess is higher than the number I picked (i.e. num > pick).
    1: Your guess is lower than the number I picked (i.e. num < pick).
    0: your guess is equal to the number I picked (i.e. num == pick).
Return the number that I picked.
*/

func guessNumber(n int, guess func(int) int) int {
	var start, mid int
	mid = n / 2
	for {
		switch guess(mid) {
		case -1:
			n = mid - 1
		case 1:
			start = mid + 1
		case 0:
			return mid
		}
		mid = ((n - start) / 2) + start
	}
}
