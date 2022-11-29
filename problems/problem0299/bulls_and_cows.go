package problem0299

import "fmt"

/*
You are playing the Bulls and Cows game with your friend.

You write down a secret number and ask your friend to guess what the number is.
When your friend makes a guess, you provide a hint with the following info:
The number of "bulls", which are digits in the guess that are in the correct position.
The number of "cows", which are digits in the guess that are in your secret number but are located in the wrong position.
Specifically, the non-bull digits in the guess that could be rearranged such that they become bulls.
Given the secret number secret and your friend's guess guess, return the hint for your friend's guess.
The hint should be formatted as "xAyB", where x is the number of bulls and y is the number of cows.
Note that both secret and guess may contain duplicate digits.
*/

func getHint(secret string, guess string) string {
	var freqs = map[byte]int{}
	var bulls, cows int
	for i := range secret {
		// Counting exact matches
		if secret[i] == guess[i] {
			bulls++
		}
		// Creating the letter frequency
		freqs[secret[i]]++
	}
	for i := range guess {
		// If letter in the secret and we didn't count it already
		if val, ok := freqs[guess[i]]; ok && val > 0 {
			cows++
			freqs[guess[i]]--
		}
	}
	// Bulls are also counted as cows so we subtract them
	return fmt.Sprintf("%dA%dB", bulls, cows-bulls)
}

func getHintNoMap(secret string, guess string) string {
	var freqs [10]int
	var bulls, cows int
	for i := range secret {
		// Counting exact matches
		if secret[i] == guess[i] {
			bulls++
		}
		// Creating the letter frequency
		freqs[secret[i]-'0']++
	}
	for i := range guess {
		// If letter in the secret and we didn't count it already
		if freqs[guess[i]-'0'] > 0 {
			cows++
			freqs[guess[i]-'0']--
		}
	}
	// Bulls are also counted as cows so we subtract them
	return fmt.Sprintf("%dA%dB", bulls, cows-bulls)
}
