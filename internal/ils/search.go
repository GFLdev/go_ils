package ils

type LocalSearch struct {
	best ILS
	list ILS
	n    int
}

// Max return the max value between two integers.
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
