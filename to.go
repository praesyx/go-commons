package commons

import "strconv"

func ToInt(s string) int {
	if "" == s {
		return 0
	}
	n, _ := strconv.Atoi(s)
	return n
}

func ToUInt8(s string) uint8 {
	if "" == s {
		return 0
	}
	n, _ := strconv.Atoi(s)
	return uint8(n)
}

func ToUInt16(s string) uint16 {
	if "" == s {
		return 0
	}
	n, _ := strconv.Atoi(s)
	return uint16(n)
}

func ToUInt32(s string) uint32 {
	if "" == s {
		return 0
	}
	n, _ := strconv.Atoi(s)
	return uint32(n)
}

func ToInt8(s string) int8 {
	if "" == s {
		return 0
	}
	n, _ := strconv.Atoi(s)
	return int8(n)
}

func ToInt16(s string) int16 {
	if "" == s {
		return 0
	}
	n, _ := strconv.Atoi(s)
	return int16(n)
}

func ToInt32(s string) int32 {
	if "" == s {
		return 0
	}
	n, _ := strconv.Atoi(s)
	return int32(n)
}
