package gophig

import (
	"github.com/go-yaml/yaml"
)

// YAMLMarshaler is a Marshaler that uses the go-yaml/yaml package.
type YAMLMarshaler struct{}

// Marshal ...
func (YAMLMarshaler) Marshal(v interface{}) ([]byte, error) {
	return yaml.Marshal(v)
}

// Unmarshal ...
func (YAMLMarshaler) Unmarshal(data []byte, v interface{}) error {
	return yaml.Unmarshal(data, v)
}
