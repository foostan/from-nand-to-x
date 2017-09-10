package logic_gate

import (
	"testing"
)

func TestXorCalc(t *testing.T) {
	test_calc_for_double_input(t, Xor{}, []ExpectForDoubleInput{
		{A: false, B: false, OUT: false},
		{A: true, B: false, OUT: true},
		{A: false, B: true, OUT: true},
		{A: true, B: true, OUT: false},
	})
}
