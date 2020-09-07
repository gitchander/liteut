package jsoncho

import (
	"encoding/json"
)

type choice1 struct{}

var Choice1 = choice1{}

func (choice1) MarshalChoice(c Choice) ([]byte, error) {
	t, err := c.GetType()
	if err != nil {
		return nil, err
	}
	data, err := json.Marshal(c.Value())
	if err != nil {
		return nil, err
	}
	tv := typeValue{
		Type:  t,
		Value: json.RawMessage(data),
	}
	return json.Marshal(tv)
}

func (choice1) UnmarshalChoice(data []byte, c Choice) error {
	var tv typeValue
	err := json.Unmarshal(data, &tv)
	if err != nil {
		return err
	}
	err = c.SetType(tv.Type)
	if err != nil {
		return err
	}
	return json.Unmarshal(tv.Value, c.Value())
}
