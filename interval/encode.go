package interval

import (
	"encoding"
)

func _() {
	var t Interval
	var (
		_ encoding.TextMarshaler   = t
		_ encoding.TextUnmarshaler = &t
	)
}

func (v Interval) MarshalText() (text []byte, err error) {
	s := defaultStringFormatter.Format(v)
	return []byte(s), nil
}

func (p *Interval) UnmarshalText(text []byte) error {
	v, err := Parse(string(text))
	if err != nil {
		return err
	}
	*p = v
	return nil
}
