package logic_gate

import (
	"testing"
)

func TestBasicNorCalc(t *testing.T) {
	test_calc_for_nor(t, BasicNor{})
}

func TestNorCalc(t *testing.T) {
	test_calc_for_nor(t, Nor{})
}

func test_calc_for_nor(t *testing.T, lg DoubleInputLogicGate) {
	test_calc_for_double_input(t, lg, []ExpectForDoubleInput{
		{A: false, B: false, OUT: true},
		{A: true, B: false, OUT: false},
		{A: false, B: true, OUT: false},
		{A: true, B: true, OUT: false},
	})
}
