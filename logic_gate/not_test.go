package logic_gate

import (
	"testing"
)

func TestBasicNotCalc(t *testing.T) {
	test_calc_for_not(t, BasicNot{})
}

func TestNotCalc(t *testing.T) {
	test_calc_for_not(t, Not{})
}

func test_calc_for_not(t *testing.T, lg SingleInputLogicGate) {
	test_input_calc_for_single_input(t, lg, []ExpectForSingleInput{
		{A: true, OUT: false},
		{A: false, OUT: true},
	})
}
