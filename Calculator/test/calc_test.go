package test

import (
	"Calculator/internal"
	"testing"
)

func TestExpression(t *testing.T) {
	t.Run("Calc : 1+3*(1+2/1-(2-1)", func(t *testing.T) {
		if result := internal.Calculate("1+3*(1+2/1-(2-1))"); result != 7.0 {
			t.Error("Error in calculation ", result)
		}
	})

	t.Run("Calc : 2-2*2+2-5/5", func(t *testing.T) {
		if result := internal.Calculate("2-2*2+2-5/5"); result != -1 {
			t.Error("Error in calculation ", result)
		}
	})

	t.Run("Calc : 2+2*2", func(t *testing.T) {
		if result := internal.Calculate("2+2*2"); result != 6 {
			t.Error("Error in calculation ", result)
		}
	})
}

func TestTokenize(t *testing.T) {
	var opMap = map[string]int{
		"*": 2,
		"/": 2,
		"+": 1,
		"-": 1,
		"(": -1,
		")": -1,
	}
	var expected = []string{"5", "*", "2", "-", "6", "/", "3"}
	t.Run("Tokenize: 5*2-6/3", func(t *testing.T) {
		testEquality := func(a, b []string) bool {
			if len(a) != len(b) {
				return false
			}
			for i := range a {
				if a[i] != b[i] {
					return false
				}
			}
			return true
		}
		if result := internal.Tokenize("5*2-6/3", opMap); !testEquality(expected, result) {
			t.Error("Error in tokenize ", result, expected)
		}
	})
}
