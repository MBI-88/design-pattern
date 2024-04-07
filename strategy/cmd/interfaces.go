package cmd

type StrategyInt interface {
	FindBreadth([]int) int
}

func NewStrategy(variant string) StrategyInt {

	switch variant {
	case "a":
		return &algorithm1{}
	default:
		return &algorithm2{}

	}
}
