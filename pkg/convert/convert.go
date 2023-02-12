package convert

import (
	"encoding/json"
	"strconv"
)

func MustInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

type IJsonMarshaler interface {
	Marshal(v any) ([]byte, error)
}

type StdMarshaler struct{}

func (m *StdMarshaler) Marshal(v any) ([]byte, error) {
	return json.Marshal(v)
}
