package logic_gate

import (
	"testing"
)

type ExpectForDMuxInput struct {
	IN  bool
	SEL bool
	A   bool
	B   bool
}

func TestBasicDMuxCalc(t *testing.T) {
	test_calc_for_dmux(t, BasicDMux{})
}

func TestDMuxCalc(t *testing.T) {
	test_calc_for_dmux(t, DMux{})
}

func test_calc_for_dmux(t *testing.T, lg DMuxIf) {
	expects := []ExpectForDMuxInput{
		{IN: false, SEL: false, A: false, B: false},
		{IN: true, SEL: false, A: true, B: false},
		{IN: false, SEL: true, A: false, B: false},
		{IN: true, SEL: true, A: false, B: true},
	}

	for i := 0; i < len(expects); i++ {
		actual_a, actual_b := lg.Calc(expects[i].IN, expects[i].SEL)
		expected_a := expects[i].A
		expected_b := expects[i].B
		if actual_a != expected_a {
			t.Errorf("got %v\nwant %v", actual_a, expected_b)
		}
		if actual_b != expected_b {
			t.Errorf("got %v\nwant %v", actual_b, expected_b)
		}
	}
}
