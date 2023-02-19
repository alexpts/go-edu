package convert

import (
	"encoding/json"
)

type IJsonMarshaler interface {
	Marshal(v any) ([]byte, error)
	Unmarshal(data []byte, v any) error
}

// StdJsonMarshaler - adapter for std json package to IJsonMarshaler
type StdJsonMarshaler struct{}

func (m *StdJsonMarshaler) Marshal(v any) ([]byte, error) {
	return json.Marshal(v)
}

func (m *StdJsonMarshaler) Unmarshal(data []byte, v any) error {
	return json.Unmarshal(data, v)
}
