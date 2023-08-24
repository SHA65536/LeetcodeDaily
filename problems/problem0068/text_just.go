package problem0068

/*
Given an array of strings words and a width maxWidth,
format the text such that each line has exactly maxWidth characters and is fully (left and right) justified.
You should pack your words in a greedy approach; that is, pack as many words as you can in each line.
Pad extra spaces ' ' when necessary so that each line has exactly maxWidth characters.
Extra spaces between words should be distributed as evenly as possible.
If the number of spaces on a line does not divide evenly between words,
the empty slots on the left will be assigned more spaces than the slots on the right.
For the last line of text, it should be left-justified, and no extra space is inserted between words.
Note:
    A word is defined as a character sequence consisting of non-space characters only.
    Each word's length is guaranteed to be greater than 0 and not exceed maxWidth.
    The input array words contains at least one word.
*/

func fullJustify(words []string, maxWidth int) []string {
	// Assign each word to a line
	var wordsPerLine [][]string = [][]string{nil}
	var curLine, curIdx int
	for i := 0; i < len(words); i++ {
		if curIdx+len(words[i]) > maxWidth {
			wordsPerLine = append(wordsPerLine, []string{})
			curLine++
			curIdx = 0
			i--
		} else {
			wordsPerLine[curLine] = append(wordsPerLine[curLine], words[i])
			curIdx += len(words[i]) + 1
		}
	}

	// Justify each line
	var res = make([]string, len(wordsPerLine))
	for i := 0; i < len(wordsPerLine)-1; i++ {
		res[i] = justifyLine(wordsPerLine[i], maxWidth)
	}

	// Last line is left just
	last := emptyLine(maxWidth)
	var i int
	for _, word := range wordsPerLine[len(wordsPerLine)-1] {
		for j := range word {
			last[i] = word[j]
			i++
		}
		i++
	}
	res[len(wordsPerLine)-1] = string(last)
	return res
}

func justifyLine(words []string, maxWidth int) string {
	var sumLen, spaces, extra int
	var res []byte = emptyLine(maxWidth)
	// One word is special case, left just
	if len(words) == 1 {
		for i := range words[0] {
			res[i] = words[0][i]
		}
		return string(res)
	}

	for i := range words {
		sumLen += len(words[i])
	}
	var idx int
	// Normal spaces between words
	spaces = (maxWidth - sumLen) / (len(words) - 1)
	// Extra spaces for the left words (if not evenly divided)
	extra = (maxWidth - sumLen) % (len(words) - 1)
	for i := range words {
		// Copy word
		for j := range words[i] {
			res[idx] = words[i][j]
			idx++
		}
		// Add normal spaces
		idx += spaces
		// Add extra spaces if there are any
		if extra > 0 {
			extra--
			idx++
		}
	}
	return string(res)
}

func emptyLine(n int) []byte {
	var res = make([]byte, n)
	for i := range res {
		res[i] = ' '
	}
	return res
}
