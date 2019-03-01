package easy

import "strings"

func IsPalindrome2(s string) bool {
	l := len(s)
	if l == 0 {
		return true
	}
	s = strings.ToLower(s)
	for i, j := 0, l-1; i != j; {
		a, b := s[i], s[j]
		if a == b {
			i++
			j--
			if i >= j {
				// s="aa",此时存在i>j
				break
			}
			continue
		}
		if !(a >= 'a' && a <= 'z' || a >= '0' && a <= '9') {
			i++
			continue
		}
		if !(b >= 'a' && b <= 'z' || b >= '0' && b <= '9') {
			j--
			continue
		}
		return false
	}
	return true
}
