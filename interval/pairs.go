package interval

func OverlappedPairs(as []Interval) (pairs [][2]int) {
	f := func(i, j int) bool {
		if as[i].Overlaps(as[j]) {
			pair := [2]int{i, j}
			pairs = append(pairs, pair)
		}
		return true
	}
	walkPairsOfIndexes(len(as), f)
	return pairs
}

func walkPairsOfIndexes(n int, f func(i, j int) bool) {
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if !f(i, j) {
				return
			}
		}
	}
}
