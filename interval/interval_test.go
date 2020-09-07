package interval

import (
	"testing"
)

func TestIntervalEmpty(t *testing.T) {

	var samples = []struct {
		r     Interval
		empty bool
	}{
		{
			Interval{Min: 0, Max: 0},
			true,
		},
		{
			Interval{Min: -1, Max: -1},
			true,
		},
		{
			Interval{Min: 1, Max: 1},
			true,
		},
		{
			Interval{Min: 0, Max: -1},
			true,
		},
		{
			Interval{Min: 0, Max: 1},
			false,
		},
		{
			Interval{Min: -1, Max: 0},
			false,
		},
	}

	for i, sample := range samples {
		if sample.r.Empty() != sample.empty {
			t.Fatalf("invalid sample %d: %t != %t", i, sample.r.Empty(), sample.empty)
		}
	}
}

func TestParse(t *testing.T) {
	samples := []struct {
		v  Interval
		cs []string
	}{
		{
			v: Interval{
				Min: 0,
				Max: 0,
			},
			cs: []string{
				"[0..0)",
				"[0..-1]",
				"(0..0)",
				"(0..0]",
				"[10..-7)",
				"[100..99)",
				"[-1..-7)",
			},
		},
		{
			v: Interval{
				Min: 5,
				Max: 17,
			},
			cs: []string{
				"[5..17)",
				"[5..16]",
				"(4..17)",
				"(4..16]",
			},
		},
		{
			v: Interval{
				Min: 0,
				Max: 1,
			},
			cs: []string{
				"[0..1)",
				"[0..0]",
				"(-1..1)",
				"(-1..0]",
			},
		},
	}

	for _, sample := range samples {
		a := sample.v
		for _, c := range sample.cs {
			b, err := Parse(c)
			if err != nil {
				t.Fatal(err)
			}
			if !(a.Equal(b)) {
				t.Fatalf("%s != %s", a, b)
			}
		}
	}
}
