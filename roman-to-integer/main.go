package romantointeger

func romanToInt(s string) int {
	romanDic := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	sum := 0
	isSkip := false

	for i, r := range s {
		if isSkip {
			isSkip = false
			continue
		}
		str := string(r)
		nextStr := str

		if i < len(s)-1 {
			nextStr = s[i+1 : i+2]
		}

		if romanDic[nextStr] > romanDic[str] {
			sum += romanDic[nextStr] - romanDic[str]
			isSkip = true
			continue
		}

		sum += romanDic[str]
	}

	return sum
}
