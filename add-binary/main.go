package addbinary

func swap(a, b string) (string, string) {
	if len(a) < len(b) {
		return a, b
	}
	return b, a
}

func addBinary(a string, b string) string {
	a, b = swap(a, b)
	res := ""
	mem := ""

	for i := len(a) - 1; i >= 0; i-- {
		j := len(b) - len(a) + i

		if (a[i] != b[j] && mem == "") ||
			(a[i] == b[j] && string(a[i]) == "1" && mem == "1") {
			res = "1" + res
		} else if a[i] == b[j] && string(a[i]) != "1" && mem == "" ||
			(a[i] != b[j] && mem == "1") {
			res = "0" + res
		} else if a[i] == b[j] && string(a[i]) == "1" && mem == "" {
			res = "0" + res
			mem = "1"
		} else if a[i] == b[j] && string(a[i]) != "1" && mem == "1" {
			res = "1" + res
			mem = ""
		}

		if j > 0 && i == 0 {
			if mem != "" {
				res = addBinary(b[:j], mem) + res
			} else {
				res = b[:j] + res
			}
		} else if j == 0 {
			res = mem + res
		}
	}

	return res
}
