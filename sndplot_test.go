package main

import (
	"math"
	"testing"
)

func TestParseFunctionBad(t *testing.T) {
	bad := []string{"sin(x", "bla(x)", "y", "x +"}

	for _, s := range bad {
		_, err := parseFunction(s)
		if err == nil {
			t.Error("Function should not be parseable: " + s)
		}
	}
}

func TestParseFunction(t *testing.T) {
	tests := map[string]func(float64) float64{
		"sin(x)":            func(x float64) float64 { return math.Sin(x) },
		"cos(x)":            func(x float64) float64 { return math.Cos(x) },
		"cos(x) + sin(x)":   func(x float64) float64 { return math.Cos(x) + math.Sin(x) },
		"sin(x) + cos(x)":   func(x float64) float64 { return math.Cos(x) + math.Sin(x) },
		"5*cos(x) + sin(x)": func(x float64) float64 { return 5*math.Cos(x) + math.Sin(x) },
	}

	for s, expect := range tests {
		f, err := parseFunction(s)
		if err != nil {
			t.Error("Function could not be parsed: " + s)
		} else if !compareFunctions(f, expect) {
			t.Error("Function parsed incorrectly: " + s)
		}
	}
}

func compareFunctions(a, b func(float64) float64) bool {
	for x := -5.0; x < 5.0; x += 0.05 {
		if a(x) != b(x) {
			return false
		}
	}
	return true
}
