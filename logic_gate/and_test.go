package logic_gate

import (
	"testing"
)

func TestAndCalc(t *testing.T) {
	test_calc_for_double_input(t, And{}, []ExpectForDoubleInput{
		{A: false, B: false, OUT: false},
		{A: true, B: false, OUT: false},
		{A: false, B: true, OUT: false},
		{A: true, B: true, OUT: true},
	})
}
