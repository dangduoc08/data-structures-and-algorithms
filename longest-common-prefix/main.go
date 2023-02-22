package longestcommonprefix

func findlongestCommonPrefix2Strs(str1, str2 string) string {
	res := ""
	s := str1
	if len(str2) < len(str1) {
		s = str2
	}

	for i := range s {
		if str1[i:i+1] == str2[i:i+1] {
			res += str1[i : i+1]
			continue
		}
		break
	}

	return res
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 1 {
		return strs[0]
	}

	res := findlongestCommonPrefix2Strs(strs[0], strs[1])
	if res == "" {
		return res
	}

	i := 2
	for i < len(strs) {
		str := strs[i]
		res = findlongestCommonPrefix2Strs(res, str)
		if res == "" {
			return res
		}
		i++
	}

	return res
}
