package interval

// Some ideas for Interval get from image.Rectangle

// Interval
// [min..max) - Max value isn't contained

// Axioms:
// if (min >= max) it is empty interval
// if interval contains x than (min <= x < max)
type Interval struct {
	Min int `json:"min"`
	Max int `json:"max"`
}

var ZI Interval // Zero Interval

// Ivl is shorthand for Interval{Min, Max}
func Ivl(min, max int) Interval {
	return Interval{
		Min: min,
		Max: max,
	}
}

func (a Interval) String() string {
	return defaultStringFormatter.Format(a)
}

func (a Interval) Empty() bool {
	return (a.Min >= a.Max)
}

func (a Interval) notEmpty() bool {
	return (a.Min < a.Max)
}

func (a Interval) Width() int {
	if a.notEmpty() {
		return a.Max - a.Min
	}
	return 0
}

func (a Interval) Equal(b Interval) bool {
	return (a == b) || (a.Empty() && b.Empty())
}

func (a Interval) Contains(v int) bool {
	return a.notEmpty() && (a.Min <= v) && (v < a.Max)
}

func (a Interval) Overlaps(b Interval) bool {
	return a.notEmpty() && b.notEmpty() &&
		(a.Min < b.Max) && (b.Min < a.Max)
}

func (a Interval) Add(d int) Interval {
	return Interval{
		Min: a.Min + d,
		Max: a.Max + d,
	}
}

func (a Interval) Sub(d int) Interval {
	return Interval{
		Min: a.Min - d,
		Max: a.Max - d,
	}
}

func (a Interval) Split(count int) []Interval {

	if a.Empty() {
		return nil
	}

	width := a.Width()
	if count > width {
		count = width
	}

	if count <= 0 {
		return nil
	}

	quo, rem := quoRem(width, count)

	bs := make([]Interval, count)

	var b Interval
	b.Min = a.Min
	for i := 0; i < count; i++ {
		b.Max = b.Min + quo
		if rem > 0 {
			b.Max++
			rem--
		}
		bs[i] = b
		b.Min = b.Max
	}

	return bs
}

func quoRem(a, b int) (quo, rem int) {
	quo = a / b
	rem = a % b
	return
}
