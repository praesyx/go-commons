package commons

import "unicode"

func TruncateByLen(str string, max int) string {
	lastSpaceIx := -1
	n := 0
	for i, r := range str {
		if unicode.IsSpace(r) {
			lastSpaceIx = i
		}
		n++
		if n >= max {
			if lastSpaceIx != -1 {
				return str[:lastSpaceIx] + "..."
			}
		}
	}

	return str
}
