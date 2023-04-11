package commons

import "strings"

func ParseStringNumToString(s string) string {
	nLen := 0
	for i := 0; i < len(s); i++ {
		if b := s[i]; '0' <= b && b <= '9' {
			nLen++
		}
	}
	var n = make([]string, 0, nLen)
	for i := 0; i < len(s); i++ {
		if b := s[i]; '0' <= b && b <= '9' {
			n = append(n, string(b))
		}
	}
	return strings.Join(n, "")
}
