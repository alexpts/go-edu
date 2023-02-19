package convert

import "strconv"

func ToPtr[T any](val T) *T {
	return &val
}

func MustInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
