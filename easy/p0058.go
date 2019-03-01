package easy

func LengthOfLastWord(s string) int {
	l := len(s)
	if l == 0 {
		return 0
	}
	c := 0
	for i := l - 1; i >= 0; i-- {
		if s[i] == ' ' {
			if c != 0 {
				return c
			}
			continue
		}
		c++
	}
	return c
}
