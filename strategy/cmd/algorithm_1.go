package cmd

import "sort"

type algorithm1 struct{}

func (algorithm1) FindBreadth(set []int) int {
	sort.Ints(set)
	return set[len(set) - 1] - set[0]
}