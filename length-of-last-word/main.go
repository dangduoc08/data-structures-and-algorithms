package lengthoflastword

func lengthOfLastWord(s string) int {
	if len(s) == 1 {
		return 1
	}

	var space rune = 32
	counter := 0

	for i, r := range s {
		if i < len(s)-1 && r == space && s[i+1] != byte(space) {
			counter = 0
		} else if r != space {
			counter++
		}

	}

	return counter
}
