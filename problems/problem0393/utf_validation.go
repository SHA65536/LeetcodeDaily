package problem0393

/*
Given an integer array data representing the data, return whether it is a valid UTF-8 encoding
(i.e. it translates to a sequence of valid UTF-8 encoded characters).
A character in UTF8 can be from 1 to 4 bytes long, subjected to the following rules:
For a 1-byte character, the first bit is a 0, followed by its Unicode code.
For an n-bytes character, the first n bits are all one's, the n + 1 bit is 0, followed by n - 1 bytes with the most significant 2 bits being 10.
This is how the UTF-8 encoding would work:
*/

// Masks to look at the first couple of bits of a number
const (
	oneBit    = 0b10000000
	twobits   = 0b11000000
	threebits = 0b11100000
	fourbits  = 0b11110000
	fivebits  = 0b11111000
)

// Values to look for in the headers
const (
	oneChar    = 0b00000000
	twoChars   = 0b11000000
	threeChars = 0b11100000
	fourChars  = 0b11110000
	contChar   = 0b10000000
)

func validUtf8(data []int) bool {
	var continuesLeft int
	for i := range data {
		// For each byte we check which type of byte it is
		// And set the number of expected future bytes accordingly
		switch {
		case data[i]&oneBit == oneChar:
			if continuesLeft != 0 {
				return false
			}
		case data[i]&twobits == contChar:
			if continuesLeft == 0 {
				return false
			}
			continuesLeft--
		case data[i]&threebits == twoChars:
			if continuesLeft != 0 {
				return false
			}
			continuesLeft = 1
		case data[i]&fourbits == threeChars:
			if continuesLeft != 0 {
				return false
			}
			continuesLeft = 2
		case data[i]&fivebits == fourChars:
			if continuesLeft != 0 {
				return false
			}
			continuesLeft = 3
		default:
			return false
		}
	}
	// If we're expecting more bytes, it's invalid
	return continuesLeft == 0
}
