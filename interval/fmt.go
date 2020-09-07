package interval

import (
	"fmt"
	"strconv"
	"strings"
)

// https://www.mathsisfun.com/sets/intervals.html

/*

type StringFormatter interface {
	Format(Interval) string
}

type StringConverter interface {
	Convert(Interval) string
}

*/

// format: ('(' or '[') + min + ".." + max + (')' or ']')

const delimiter = ".."

// (min, max) - an open interval
// [min, max] - a closed interval
// (min, max] - open on min, closed on max
// [min, max) - closed on min, open on max

type StringFormatter struct {
	MinClosed bool
	MaxClosed bool
}

var defaultStringFormatter = StringFormatter{
	MinClosed: true,
	MaxClosed: false,
}

func (sf StringFormatter) Format(v Interval) string {

	if v.Empty() {
		v = ZI
	}

	var (
		min = v.Min // closed
		max = v.Max // open
	)

	var openBracket, closeBracket byte

	if sf.MinClosed {
		openBracket = '['
	} else {
		openBracket = '('
		min-- // closed -> open
	}

	if sf.MaxClosed {
		closeBracket = ']'
		max-- // open -> closed
	} else {
		closeBracket = ')'
	}

	var b strings.Builder
	b.WriteByte(openBracket)
	b.WriteString(strconv.Itoa(min))
	b.WriteString(delimiter)
	b.WriteString(strconv.Itoa(max))
	b.WriteByte(closeBracket)
	return b.String()
}

func Parse(s string) (v Interval, err error) {

	s = strings.TrimSpace(s)

	// brackets: "[]", "(]", "[)", "()"
	if len(s) < 2 {
		return ZI, fmt.Errorf("parse interval: there are not brackets")
	}

	// Check open bracket
	openBracket := s[0]
	if (openBracket != '[') && (openBracket != '(') {
		return ZI, fmt.Errorf("parse interval: open bracket is wrong %q", openBracket)
	}

	// Check close bracket
	closeBracket := s[len(s)-1]
	if (closeBracket != ']') && (closeBracket != ')') {
		return ZI, fmt.Errorf("parse interval: close bracket is wrong %q", closeBracket)
	}

	// trim brackets
	s = s[1 : len(s)-1]

	xs := strings.Split(s, delimiter)

	const n = 2 // 2 params (min, max)

	if len(xs) != n {
		return ZI, fmt.Errorf("parse interval: wrong number of values: have %d, want %d", len(xs), n)
	}

	var min, max int

	ds := [n]*int{
		&min,
		&max,
	}

	for i, x := range xs {
		d, err := strconv.Atoi(x)
		if err != nil {
			return ZI, err
		}
		*(ds[i]) = d
	}

	if openBracket == '(' {
		min++ // open -> closed
	}
	if closeBracket == ']' {
		max++ // closed -> open
	}

	v = Interval{
		Min: min, // closed
		Max: max, // open
	}

	if v.Empty() {
		v = ZI
	}

	return v, nil
}
