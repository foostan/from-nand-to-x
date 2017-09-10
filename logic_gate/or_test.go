package logic_gate

import (
	"testing"
)

func TestBasicOrCalc(t *testing.T) {
	test_calc_for_or(t, BasicOr{})
}

func TestOrCalc(t *testing.T) {
	test_calc_for_or(t, Or{})
}

func test_calc_for_or(t *testing.T, lg DoubleInputLogicGate) {
	test_calc_for_double_input(t, lg, []ExpectForDoubleInput{
		{A: false, B: false, OUT: false},
		{A: true, B: false, OUT: true},
		{A: false, B: true, OUT: true},
		{A: true, B: true, OUT: true},
	})
}
