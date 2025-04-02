package ils

import "strconv"

// ComparisonResult represents an integer value used to indicate the result of a relational comparison.
type ComparisonResult int

const (
	// Equal represents the equality comparison result.
	Equal ComparisonResult = iota
	// Less represents the lesser comparison result.
	Less
	// Greater represents the greater comparison result.
	Greater
)

// ILS represents an ILS solution struct.
type ILS struct {
	seq    []int
	length int
	flow   int
}

// NewEmptyILS creates and returns a pointer to a new, empty instance of the ILS struct.
func NewEmptyILS() *ILS {
	return &ILS{}
}

// NewILS initializes and returns a pointer to an ILS struct with the given sequence, flow, and length values.
func NewILS(seq []int, flow, length int) *ILS {
	return &ILS{
		seq:    seq,
		length: length,
		flow:   flow,
	}
}

// Relational compares the flow values of two ILS pointers and returns a ComparisonResult indicating the comparison
// result.
func Relational(i1, i2 *ILS) ComparisonResult {
	if i1.flow == i2.flow {
		return Equal
	} else if i1.flow > i2.flow {
		return Greater
	} else {
		return Less
	}
}

// Get retrieves the integer at the specified index in the sequence.
func (i *ILS) Get(idx int) int {
	return i.seq[idx]
}

// Len returns the length of the sequence in the ILS structure.
func (i *ILS) Len() int {
	return i.length
}

// Flow retrieves the value of the flow field from the ILS structure.
func (i *ILS) Flow() int {
	return i.flow
}

// SetFlow sets the flow field of the ILS structure to the specified integer value.
func (i *ILS) SetFlow(flow int) {
	i.flow = flow
}

// Set assigns values to the seq, flow, and length fields of the ILS structure.
func (i *ILS) Set(seq []int, flow, length int) {
	i.seq = seq
	i.flow = flow
	i.length = length
}

// IsEmpty checks if the sequence in the ILS structure is nil or has a length of zero (empty).
func (i *ILS) IsEmpty() bool {
	return i.seq == nil || len(i.seq) == 0
}

// String returns a string representation of the ILS structure, including sequence elements and the flow value.
func (i *ILS) String() string {
	out := "Seq: ["
	for idx := range i.length {
		out += strconv.Itoa(i.seq[idx])
		if idx < i.length-1 {
			out += " "
		}
	}
	out += "] F = " + strconv.Itoa(i.flow)
	return out
}
