package logic_gate

import (
	"testing"
)

func TestBasicOr16Calc(t *testing.T) {
	test_calc_for_or16(t, BasicOr16{})
}

func TestOr16Calc(t *testing.T) {
	test_calc_for_or16(t, Or16{})
}

func test_calc_for_or16(t *testing.T, lg Double16InputLogicGate) {
	test_input_calc_for_double_16_input(t, lg, []ExpectForDouble16Input{
		{
			A:   [16]bool{false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false},
			B:   [16]bool{false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false},
			OUT: [16]bool{false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false},
		},
		{
			A:   [16]bool{true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true},
			B:   [16]bool{false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false},
			OUT: [16]bool{true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true},
		},
		{
			A:   [16]bool{false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false},
			B:   [16]bool{true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true},
			OUT: [16]bool{true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true},
		},
		{
			A:   [16]bool{true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true},
			B:   [16]bool{true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true},
			OUT: [16]bool{true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true},
		},
	})
}
