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

type ExpectForSingle16Input struct {
	A   [16]bool
	OUT [16]bool
}

func test_input_calc_for_single_16_input(t *testing.T, lg Single16InputLogicGate, expects []ExpectForSingle16Input) {
	for i := 0; i < len(expects); i++ {
		actual := lg.Calc(expects[i].A)
		expected := expects[i].OUT
		// slice の比較
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

type ExpectForDouble16Input struct {
	A   [16]bool
	B   [16]bool
	OUT [16]bool
}

func test_input_calc_for_double_16_input(t *testing.T, lg Double16InputLogicGate, expects []ExpectForDouble16Input) {
	for i := 0; i < len(expects); i++ {
		actual := lg.Calc(expects[i].A, expects[i].B)
		expected := expects[i].OUT
		// slice の比較
		if actual != expected {
			t.Errorf("got %v\nwant %v", actual, expected)
		}
	}
}
