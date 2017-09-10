package logic_gate

import (
	"testing"
)

func TestBasicAndCalc(t *testing.T) {
	test_calc_for_and(t, BasicAnd{})
}

func TestAndCalc(t *testing.T) {
	test_calc_for_and(t, And{})
}

func test_calc_for_and(t *testing.T, lg DoubleInputLogicGate) {
	test_calc_for_double_input(t, lg, []ExpectForDoubleInput{
		{A: false, B: false, OUT: false},
		{A: true, B: false, OUT: false},
		{A: false, B: true, OUT: false},
		{A: true, B: true, OUT: true},
	})
}
