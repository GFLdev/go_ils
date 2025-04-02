package ils

import "errors"

type LCGConstant int64

const (
	Modulus     LCGConstant = 2147483647 // (2^31) - 1
	Multiplier  LCGConstant = 16807
	DivFactor   LCGConstant = 127773
	AdjustConst LCGConstant = 2836
)

type LocalSearch struct {
	best Solution
	list Solution
	// nJobs represents the number of jobs for each machine.
	nJobs int
	// nMach represents the number of machines used in search.
	nMach int
	// nPert represents the number of perturbations.
	nPert     int
	idxThread int
	nThreads  int
	removed   [5]int
	pos       int
	seed      int64
	p         [501][501]int
	s         [501]int
	operation int
}

// Max return the max value between two integers.
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func NewLocalSearch(sol Solution, seed int64, nThreads, idxThread, nPert, nJobs, nMach int, l *[][]int) *LocalSearch {
	ls := &LocalSearch{
		best:      sol,
		nJobs:     nJobs,
		nMach:     nMach,
		nPert:     nPert,
		idxThread: idxThread,
		nThreads:  nThreads,
		seed:      seed,
	}
	for i := range nMach {
		for j := range nJobs {
			ls.p[i][j] = (*l)[i][j]
		}
	}
	return ls
}

func (ls *LocalSearch) Set(sol Solution, seed int64, nThreads, idxThread, nPert, nJobs, nMach int, l *[][]int) {
	ls.best = sol
	ls.nJobs = nJobs
	ls.nMach = nMach
	ls.nPert = nPert
	ls.idxThread = idxThread
	ls.nThreads = nThreads
	ls.seed = seed
	for i := range nMach {
		for j := range nJobs {
			ls.p[i][j] = (*l)[i][j]
		}
	}
}

func (ls *LocalSearch) GetSolution() Solution {
	return ls.best
}

func (ls *LocalSearch) SetSolution(sol Solution) {
	ls.best = sol
}

func (ls *LocalSearch) SetPosition(pos int) {
	ls.pos = pos
}

// Random generates a random integer within the range [min, max].
func (ls *LocalSearch) Random(min, max int, seed *int64) int {
	quotient := *seed / int64(DivFactor)
	*seed = int64(Multiplier)*(*seed%int64(DivFactor)) - quotient*int64(AdjustConst)
	if *seed < 0 {
		*seed += int64(Modulus)
	}

	normalized := float64(*seed) / float64(Modulus)
	randomValue := int(normalized*float64(max-min+1)) + min
	return randomValue
}

func (ls *LocalSearch) GetFlow() (int, error) {
	if len(ls.s) >= ls.nJobs {
		return 0, errors.New("slice length is greater than the number of jobs")
	}

	flowChart := make([][]int, ls.nMach)
	for i := range flowChart {
		flowChart[i] = make([]int, ls.nJobs)
	}

	flowChart[0][0] = ls.p[0][ls.s[0]]
	for i := 1; i < ls.nJobs; i++ {
		flowChart[0][i] = ls.p[0][ls.s[i]] + flowChart[0][i-1]
	}
	for i := 1; i < ls.nMach; i++ {
		flowChart[i][0] = ls.p[i][ls.s[0]] + flowChart[i-1][0]
		for j := 1; j < ls.nJobs; j++ {
			flowChart[i][j] = Max(flowChart[i-1][j], flowChart[i][j-1]) + ls.p[i][ls.s[j]]
		}
	}

	total := 0
	for i := range ls.nJobs {
		total += flowChart[ls.nMach-1][i]
	}
	return total, nil
}

func (ls *LocalSearch) RandomSolution() (*Solution, error) {
	for i := range ls.nJobs {
		ls.s[i] = i
	}

	// Generate random sequence
	for i := range ls.nJobs - 1 {
		r := ls.Random(1, ls.nJobs-1, &ls.seed) + i
		buf := ls.s[i]
		ls.s[i] = ls.s[r]
		ls.s[r] = buf
	}

	flow, err := ls.GetFlow()
	if err != nil {
		return nil, err
	}
	return NewSolution(ls.s[:], flow, ls.nJobs), nil
}
