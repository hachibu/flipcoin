package web

import "time"

type PriceChart struct {
	X []time.Time `json:"x"`
	Y []int32     `json:"y"`
}

func NewPriceChart() PriceChart {
	return PriceChart{
		[]time.Time{},
		[]int32{},
	}
}
