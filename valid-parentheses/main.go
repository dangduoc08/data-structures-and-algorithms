package validparentheses

// func calcCloseBracketOffset(i, totalOpenBrackets int) int {
// 	return i + (totalOpenBrackets * 2) + 1
// }

// func isValid(s string) bool {
// 	bracketMap := map[string]string{
// 		"(": ")",
// 		"[": "]",
// 		"{": "}",
// 	}
// 	checkedCase := map[int]bool{}

// 	length := len(s)
// 	firstChar := s[0:1]
// 	lastChar := s[length-1:]

// 	if length%2 != 0 ||
// 		bracketMap[firstChar] == "" ||
// 		(lastChar != bracketMap["("] &&
// 			lastChar != bracketMap["["] &&
// 			lastChar != bracketMap["{"]) {
// 		return false
// 	}

// 	if length == 2 {
// 		return bracketMap[firstChar] == lastChar
// 	}

// 	i := 0
// 	for i < len(s) {
// 		if checkedCase[i] {
// 			i++
// 			continue
// 		}
// 		totalOpenBrackets := 0
// 		totalCloseBrackets := 0
// 		j := i + 1
// 		for j < len(s) {
// 			if bracketMap[string(s[j])] != "" {
// 				totalOpenBrackets++
// 			} else if string(s[j]) == bracketMap["("] ||
// 				string(s[j]) == bracketMap["["] ||
// 				string(s[j]) == bracketMap["{"] {
// 				if totalOpenBrackets == 0 {
// 					break
// 				}

// 				totalCloseBrackets++
// 			}
// 			if totalOpenBrackets-totalCloseBrackets == 0 && bracketMap[string(s[j+1])] == "" {
// 				break
// 			}

// 			j++
// 		}

// 		closeBracketOffset := calcCloseBracketOffset(i, totalOpenBrackets)
// 		if closeBracketOffset >= length || bracketMap[string(s[i])] != string(s[closeBracketOffset]) {
// 			return false
// 		}
// 		checkedCase[closeBracketOffset] = true

// 		i++
// 	}

// 	return true
// }

// // 0 1 2 3 4 5 6 7
// // ( ( [ ] ) { } )
// // ( ( [ ] ) ) { }

func isValid(s string) bool {
	bracketMap := map[int32]int32{
		'(': ')',
		'[': ']',
		'{': '}',
	}
	arr := []int32{}

	length := len(s)
	firstChar := s[0:1]
	lastChar := s[length-1:]

	if length%2 != 0 ||
		bracketMap[[]rune(firstChar)[0]] == 0 ||
		(lastChar != string(bracketMap['(']) &&
			lastChar != string(bracketMap['[']) &&
			lastChar != string(bracketMap['{'])) {
		return false
	}

	for _, r := range s {
		if bracketMap[r] != 0 {
			arr = append(arr, bracketMap[r])
			continue
		}

		if len(arr) == 0 || arr[len(arr)-1] != r {
			return false
		}

		arr = arr[:len(arr)-1]
	}

	return len(arr) == 0
}
