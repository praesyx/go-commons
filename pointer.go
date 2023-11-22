package commons

import "time"

func String(val string) *string {
	return &val
}

func Int(n int) *int {
	return &n
}

func Int8(n int8) *int8 {
	return &n
}

func Int16(n int16) *int16 {
	return &n
}

func Int32(n int32) *int32 {
	return &n
}

func Int64(n int64) *int64 {
	return &n
}

func Uint8(n uint8) *uint8 {
	return &n
}

func Uint16(n uint16) *uint16 {
	return &n
}

func Uint32(n uint32) *uint32 {
	return &n
}

func Uint64(n uint64) *uint64 {
	return &n
}

func Time(t time.Time) *time.Time {
	return &t
}

func Ptr[T any](v T) *T {
	return &v
}
