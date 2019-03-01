package easy

func IsValid(s string) bool {
	l := len(s)
	if l < 1 {
		return true
	}
	data := stack{}
	for _, k := range s {
		switch k {
		case '(', '{', '[':
			data.push(k)
		case ')', '}', ']':
			if r, ok := data.pop(); ok {
				if m[r] != k {
					return false
				}
			} else {
				return false
			}
		default:
			return false
		}
	}
	if _, ok := data.pop(); ok {
		return false
	}
	return true
}

var m map[rune]rune = map[rune]rune{
	'(': ')',
	'{': '}',
	'[': ']',
}

type stack []rune

func (s *stack) push(b rune) {
	*s = append(*s, b)
}

func (s *stack) pop() (rune, bool) {
	if l := len(*s); l > 0 {
		b := (*s)[l-1]
		*s = (*s)[:l-1]
		return b, true
	}
	return 0, false
}
