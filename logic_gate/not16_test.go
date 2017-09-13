package logic_gate

import (
	"testing"
)

func TestBasicNot16Calc(t *testing.T) {
	test_calc_for_not16(t, BasicNot16{})
}

func TestNot16Calc(t *testing.T) {
	test_calc_for_not16(t, Not16{})
}

func test_calc_for_not16(t *testing.T, lg Single16InputLogicGate) {
	test_input_calc_for_single_16_input(t, lg, []ExpectForSingle16Input{
		{
			A:   [16]bool{true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true},
			OUT: [16]bool{false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false},
		},
		{
			A:   [16]bool{false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false},
			OUT: [16]bool{true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true},
		},
	})
}
