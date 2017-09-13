package logic_gate

import (
	"testing"
)

type ExpectForMux16Input struct {
	A   [16]bool
	B   [16]bool
	SEL bool
	OUT [16]bool
}

func TestBasicMux16Calc(t *testing.T) {
	test_calc_for_mux16(t, BasicMux16{})
}

func TestMux16Calc(t *testing.T) {
	test_calc_for_mux16(t, Mux16{})
}

func test_calc_for_mux16(t *testing.T, lg Mux16If) {
	expects := []ExpectForMux16Input{
		{
			A:   [16]bool{false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false},
			B:   [16]bool{false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false},
			SEL: false,
			OUT: [16]bool{false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false},
		},
		{
			A:   [16]bool{true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true},
			B:   [16]bool{false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false},
			SEL: false,
			OUT: [16]bool{true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true},
		},
		{
			A:   [16]bool{false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false},
			B:   [16]bool{true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true},
			SEL: false,
			OUT: [16]bool{false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false},
		},
		{
			A:   [16]bool{true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true},
			B:   [16]bool{true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true},
			SEL: false,
			OUT: [16]bool{true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true},
		},
		{
			A:   [16]bool{false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false},
			B:   [16]bool{false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false},
			SEL: true,
			OUT: [16]bool{false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false},
		},
		{
			A:   [16]bool{true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true},
			B:   [16]bool{false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false},
			SEL: true,
			OUT: [16]bool{false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false},
		},
		{
			A:   [16]bool{false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false},
			B:   [16]bool{true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true},
			SEL: true,
			OUT: [16]bool{true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true},
		},
		{
			A:   [16]bool{true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true},
			B:   [16]bool{true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true},
			SEL: true,
			OUT: [16]bool{true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true},
		},
	}

	for i := 0; i < len(expects); i++ {
		actual := lg.Calc(expects[i].A, expects[i].B, expects[i].SEL)
		expected := expects[i].OUT
		if actual != expected {
			t.Errorf("got %v\nwant %v", actual, expected)
		}
	}
}
