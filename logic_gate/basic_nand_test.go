package logic_gate

import (
	"testing"
)

func TestBasicNandCalc(t *testing.T) {
	test_calc_for_double_input(t, BasicNand{}, []ExpectForDoubleInput{
		{A: false, B: false, OUT: true},
		{A: true, B: false, OUT: true},
		{A: false, B: true, OUT: true},
		{A: true, B: true, OUT: false},
	})
}
