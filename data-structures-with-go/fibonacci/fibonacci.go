package fibonacci

func Fibonacci(position int, seq []int) int {
	if seq == nil {
		seq = []int{0, 1}
	}
	for len(seq) <= position {
		seq = append(seq, seq[len(seq)-1]+seq[len(seq)-2])
	}
	return seq[len(seq)-1]
}
