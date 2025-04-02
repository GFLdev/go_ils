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

// Solution represents an Solution solution struct.
type Solution struct {
	seq    []int
	length int
	flow   int
}

// NewEmptySolution creates and returns a pointer to a new, empty instance of the Solution struct.
func NewEmptySolution() *Solution {
	return &Solution{}
}

// NewSolution initializes and returns a pointer to a Solution struct with the given sequence, flow, and length values.
func NewSolution(seq []int, flow, length int) *Solution {
	return &Solution{
		seq:    seq,
		length: length,
		flow:   flow,
	}
}

// Relational compares the flow values of two Solution pointers and returns a ComparisonResult indicating the comparison
// result.
func Relational(i1, i2 *Solution) ComparisonResult {
	if i1.flow == i2.flow {
		return Equal
	} else if i1.flow > i2.flow {
		return Greater
	} else {
		return Less
	}
}

// Get retrieves the integer at the specified index in the sequence.
func (s *Solution) Get(idx int) int {
	return s.seq[idx]
}

// Len returns the length of the sequence in the Solution structure.
func (s *Solution) Len() int {
	return s.length
}

// GetFlow retrieves the value of the flow field from the Solution structure.
func (s *Solution) GetFlow() int {
	return s.flow
}

// SetFlow sets the flow field of the Solution structure to the specified integer value.
func (s *Solution) SetFlow(flow int) {
	s.flow = flow
}

// Set assigns values to the seq, flow, and length fields of the Solution structure.
func (s *Solution) Set(seq []int, flow, length int) {
	s.seq = seq
	s.flow = flow
	s.length = length
}

// Perturb swaps adjacent elements at index i and i+1 in the sequence, ensuring i is within valid bounds.
func (s *Solution) Perturb(i int) {
	if i < 0 {
		i = 0
	} else if i >= s.length-1 {
		i = s.length - 2
	}

	buf := s.seq[i]
	s.seq[i] = s.seq[i+1]
	s.seq[i+1] = buf
}

// IsEmpty checks if the sequence in the Solution structure is nil or has a length of zero (empty).
func (s *Solution) IsEmpty() bool {
	return s.seq == nil || len(s.seq) == 0
}

// String returns a string representation of the Solution structure, including sequence elements and the flow value.
func (s *Solution) String() string {
	out := "Seq: ["
	for idx := range s.length {
		out += strconv.Itoa(s.seq[idx])
		if idx < s.length-1 {
			out += " "
		}
	}
	out += "] F = " + strconv.Itoa(s.flow)
	return out
}
