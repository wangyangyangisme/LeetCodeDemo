package easy

import "fmt"

func CountAndSay(n int) string {
	str := []byte{'1'}
	for i := 1; i < n; i++ {
		temp := []byte{}
		count, ch := 0, byte(0)
		for _, v := range str {
			if count == 0 {
				ch = v
				count++
				continue
			}
			if ch == v {
				count++
			} else {
				temp = append(temp, []byte(fmt.Sprintf("%v%c", count, ch))...)
				count = 1
				ch = v
			}
		}
		if count > 0 {
			temp = append(temp, []byte(fmt.Sprintf("%v%c", count, ch))...)
		}
		str = temp
	}
	return string(str)
}
