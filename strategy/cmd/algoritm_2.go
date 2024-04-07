package cmd

import "math"

type algorithm2 struct{}

func (algorithm2) FindBreadth(set []int) int {
	var (
		min = math.MaxInt32
		max = math.MinInt64
	)

	for i := 0; i < len(set); i++ {
		if set[i] < min {
			min = set[i]
		} else if set[i] > max {
			max = set[i]
		}
	}

	return max - min
}