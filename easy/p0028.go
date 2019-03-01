package easy

func StrStr(haystack string, needle string) int {
	lh, ln := len(haystack), len(needle)
	if ln < 1 || haystack == needle {
		return 0
	}
	if lh < 1 || lh < ln {
		return -1
	}
	for i := range haystack {
		if ln+i > lh {
			return -1
		}
		if haystack[i:ln+i] == needle {
			return i
		}
	}
	return -1
}
