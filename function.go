package main

import (
	"errors"
	"math"
)

func parseFunction(s string) (generator, error) {
	if s == "cos(x)" {
		return func(x float64) float64 { return math.Cos(x) }, nil
	}

	return nil, errors.New("Not implemented")
}
