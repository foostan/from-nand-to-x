package logic_gate

import (
	"testing"
)

func TestNandCalc(t *testing.T) {
	test_calc_for_double_input(t, Nand{}, []ExpectForDoubleInput{
		{A: false, B: false, OUT: true},
		{A: true, B: false, OUT: true},
		{A: false, B: true, OUT: true},
		{A: true, B: true, OUT: false},
	})
}
