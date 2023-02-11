package convert

import "strconv"

func MustInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
