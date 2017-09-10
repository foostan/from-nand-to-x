package logic_gate

import (
	"testing"
)

type ExpectForMuxInput struct {
	A   bool
	B   bool
	SEL bool
	OUT bool
}

func TestBasicMuxCalc(t *testing.T) {
	lg := BasicMux{}

	expects := []ExpectForMuxInput{
		{A: false, B: false, SEL: false, OUT: false},
		{A: true, B: false, SEL: false, OUT: true},
		{A: false, B: true, SEL: false, OUT: false},
		{A: true, B: true, SEL: false, OUT: true},
		{A: false, B: false, SEL: true, OUT: false},
		{A: true, B: false, SEL: true, OUT: false},
		{A: false, B: true, SEL: true, OUT: true},
		{A: true, B: true, SEL: true, OUT: true},
	}

	for i := 0; i < len(expects); i++ {
		actual := lg.Calc(expects[i].A, expects[i].B, expects[i].SEL)
		expected := expects[i].OUT
		if actual != expected {
			t.Errorf("got %v\nwant %v", actual, expected)
		}
	}
}
