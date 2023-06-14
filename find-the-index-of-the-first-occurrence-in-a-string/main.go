package findtheindexofthefirstoccurrenceinastring

func strStr(haystack string, needle string) int {
	if needle == haystack {
		return 0
	}

	if len(needle) > len(haystack) {
		return -1
	}

	for i := 0; i < len(haystack)-len(needle)+1; i++ {
		if haystack[i:len(needle)+i] == needle {
			return i
		}
	}

	return -1
}
