package avg

// sum sums numbers using the iterative Kahan–Babuška algorithm
// https://en.wikipedia.org/wiki/Kahan_summation_algorithm#Further_enhancements
func sum(input []float64) float64 {
	sum := float64(0)
	cs := float64(0)
	ccs := float64(0)
	c := float64(0)
	cc := float64(0)

	for i := range input {
		t := sum + input[i]
		if sum >= input[i] {
			c = (sum - t) + input[i]
		} else {
			c = (input[i] - t) + sum
		}
		sum = t
		t = cs + c
		if cs >= c {
			cc = (cs - t) + c
		} else {
			cc = (c - t) + cs
		}
		cs = t
		ccs = ccs + cc
	}

	return sum + cs + ccs
}

func Avg(input []float64) float64 {
	return sum(input) / float64(len(input))
}
