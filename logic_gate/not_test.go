package logic_gate

import (
	"testing"
)

func TestNotCalc(t *testing.T) {
	test_input_calc_for_single_input(t, Not{}, []ExpectForSingleInput{
		{A: true, OUT: false},
		{A: false, OUT: true},
	})
}
