package jsoncho

import (
	"encoding/json"
)

type choice2 struct{}

var Choice2 = choice2{}

func (choice2) MarshalChoice(c Choice) ([]byte, error) {
	t, err := c.GetType()
	if err != nil {
		return nil, err
	}
	data, err := json.Marshal(c.Value())
	if err != nil {
		return nil, err
	}
	m := map[string]json.RawMessage{
		t: json.RawMessage(data),
	}
	return json.Marshal(m)
}

func (choice2) UnmarshalChoice(data []byte, c Choice) error {
	m := make(map[string]json.RawMessage)
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	tv, err := tvFromMap(m)
	if err != nil {
		return err
	}
	err = c.SetType(tv.Type)
	if err != nil {
		return err
	}
	return json.Unmarshal(tv.Value, c.Value())
}
