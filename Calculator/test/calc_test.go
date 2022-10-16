package test

import (
	"Calculator/internal"
	"testing"
)

func TestExpression(t *testing.T) {
	expression := "1+3*(1+2/1-(2-1))"
	expected := 7.0
	t.Run("1+3*(1+2/1-(2-1)", func(t *testing.T) {
		if result := internal.Calculate(expression); result != expected {
			t.Error("Error in calculation ", result)
		}
	})

}
