package logic_gate

import (
	"testing"
)

type Expect struct {
	IN  []bool
	OUT bool
}

func TestCalc(t *testing.T) {
	nand := Nand{}

	expects := []Expect{
		{IN: []bool{false, false}, OUT: true},
		{IN: []bool{true, false}, OUT: true},
		{IN: []bool{false, true}, OUT: true},
		{IN: []bool{true, true}, OUT: false},
	}

	for i := 0; i < len(expects); i++ {
		expect := expects[i]
		a := expect.IN[0]
		b := expect.IN[1]

		actual := nand.Calc(a, b)
		expected := expect.OUT
		if actual != expected {
			t.Errorf("got %v\nwant %v", actual, expected)
		}
	}
}
