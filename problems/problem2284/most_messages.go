package problem2284

import "strings"

/*
You have a chat log of n messages. You are given two string arrays messages and senders where messages[i] is a message sent by senders[i].
A message is list of words that are separated by a single space with no leading or trailing spaces.
The word count of a sender is the total number of words sent by the sender. Note that a sender may send more than one message.
Return the sender with the largest word count.
If there is more than one sender with the largest word count, return the one with the lexicographically largest name.
Note:
Uppercase letters come before lowercase letters in lexicographical order.
"Alice" and "alice" are distinct.
*/

func largestWordCount(messages []string, senders []string) string {
	var maxWords, count int
	var maxName string
	var senderDict = map[string]int{}
	for i := range messages {
		count = strings.Count(messages[i], " ") + 1
		senderDict[senders[i]] += count
		if maxWords < senderDict[senders[i]] {
			maxWords = senderDict[senders[i]]
		}
	}
	for k, v := range senderDict {
		if maxWords == v {
			if k > maxName {
				maxName = k
			}
		}
	}
	return maxName
}
