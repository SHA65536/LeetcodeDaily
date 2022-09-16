package problem0383

/*
Given two strings ransomNote and magazine,
Return true if ransomNote can be constructed by using the letters from magazine and false otherwise.
Each letter in magazine can only be used once in ransomNote.
*/

func canConstruct(ransomNote string, magazine string) bool {
	var ransomMap = map[byte]int{}
	var left int = len(ransomNote)
	for i := range ransomNote {
		ransomMap[ransomNote[i]]++
	}
	for i := 0; i < len(magazine) && left > 0; i++ {
		if ransomMap[magazine[i]] > 0 {
			ransomMap[magazine[i]]--
			left--
		}
	}
	return left == 0
}
