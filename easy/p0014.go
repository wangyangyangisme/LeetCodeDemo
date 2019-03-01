package easy

func LongestCommonPrefix(strs []string) string {
	if len(strs) < 1 {
		return ""
	}
	minLen := len(strs[0])
	for i := range strs {
		if len(strs[i]) < minLen {
			minLen = len(strs[i])
		}
	}
	if minLen < 1 {
		return ""
	}
	ret := make([]byte, 0, minLen)
	for i := 0; i < minLen; i++ {
		b := strs[0][i]
		for j := range strs {
			if strs[j][i] != b {
				return string(ret)
			}
		}
		ret = append(ret, b)
	}
	return string(ret)
}
