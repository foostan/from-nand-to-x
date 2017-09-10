package logic_gate

import (
	"testing"
)

func TestBasicXorCalc(t *testing.T) {
	test_calc_for_xor(t, BasicXor{})
}

func TestXorCalc(t *testing.T) {
	test_calc_for_xor(t, Xor{})
}

func test_calc_for_xor(t *testing.T, lg DoubleInputLogicGate) {
	test_calc_for_double_input(t, lg, []ExpectForDoubleInput{
		{A: false, B: false, OUT: false},
		{A: true, B: false, OUT: true},
		{A: false, B: true, OUT: true},
		{A: true, B: true, OUT: false},
	})
}
