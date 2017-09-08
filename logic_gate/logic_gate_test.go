package logic_gate

import (
	"testing"
)

type Expect struct {
	IN  []bool
	OUT bool
}

func test_calc(t *testing.T, lg LogicGate, expects []Expect) {

	for i := 0; i < len(expects); i++ {
		expect := expects[i]
		a := expect.IN[0]
		b := expect.IN[1]

		actual := lg.Calc(a, b)
		expected := expect.OUT
		if actual != expected {
			t.Errorf("got %v\nwant %v", actual, expected)
		}
	}
}
