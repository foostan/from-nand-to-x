package logic_gate

import (
	"testing"
)

type ExpectForSingleInput struct {
	A   bool
	OUT bool
}

func test_input_calc_for_single_input(t *testing.T, lg SingleInputLogicGate, expects []ExpectForSingleInput) {

	for i := 0; i < len(expects); i++ {
		actual := lg.Calc(expects[i].A)
		expected := expects[i].OUT
		if actual != expected {
			t.Errorf("got %v\nwant %v", actual, expected)
		}
	}
}

type ExpectForDoubleInput struct {
	A   bool
	B   bool
	OUT bool
}

func test_calc_for_double_input(t *testing.T, lg DoubleInputLogicGate, expects []ExpectForDoubleInput) {

	for i := 0; i < len(expects); i++ {
		actual := lg.Calc(expects[i].A, expects[i].B)
		expected := expects[i].OUT
		if actual != expected {
			t.Errorf("got %v\nwant %v", actual, expected)
		}
	}
}
