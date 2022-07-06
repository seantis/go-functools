package main

import (
	"fmt"
	"math"

	"github.com/seantis/go-functools"
)

func main() {
	vals := []float64{-5.3, 2.8, -1.0, 8.9, 0.1}
	positives := functools.Filter(vals, func(x float64) bool { return x > 0 })
	rounded := functools.Map(positives, func(x float64) int { return int(math.Round(x)) })
	summed := functools.Reduce(rounded, 0, func(x, acc int) int { return x + acc })
	fmt.Printf("%v filtered by positive numbers, rounded and summed up is %d\n", vals, summed)
}
