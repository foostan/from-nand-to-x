package logic_gate

import (
	"testing"
)

func TestBasicNotCalc(t *testing.T) {
	test_input_calc_for_single_input(t, BasicNot{}, []ExpectForSingleInput{
		{A: true, OUT: false},
		{A: false, OUT: true},
	})
}
