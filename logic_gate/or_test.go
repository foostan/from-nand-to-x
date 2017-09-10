package logic_gate

import (
	"testing"
)

func TestOrCalc(t *testing.T) {
	test_calc_for_double_input(t, Or{}, []ExpectForDoubleInput{
		{A: false, B: false, OUT: false},
		{A: true, B: false, OUT: true},
		{A: false, B: true, OUT: true},
		{A: true, B: true, OUT: true},
	})
}