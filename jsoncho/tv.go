package jsoncho

import (
	"encoding/json"
	"fmt"
)

type typeValue struct {
	Type  string          `json:"$type"`
	Value json.RawMessage `json:"$value"`
}

func tvFromMap(m map[string]json.RawMessage) (*typeValue, error) {
	if n := len(m); n != 1 {
		return nil, fmt.Errorf("keys number: %d, must be one", n)
	}
	keys := mapKeys(m)
	t := keys[0]
	return &typeValue{
		Type:  t,
		Value: m[t],
	}, nil
}

func mapKeys(m map[string]json.RawMessage) []string {
	keys := make([]string, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}
	return keys
}
