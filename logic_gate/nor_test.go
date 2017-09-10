package logic_gate

import (
	"testing"
)

func TestNorCalc(t *testing.T) {
	test_calc_for_double_input(t, Nor{}, []ExpectForDoubleInput{
		{A: false, B: false, OUT: true},
		{A: true, B: false, OUT: false},
		{A: false, B: true, OUT: false},
		{A: true, B: true, OUT: false},
	})
}
