package logic_gate

import (
	"testing"
)

func TestNandCalc(t *testing.T) {
	test_calc(t, Nand{}, []Expect{
		{IN: []bool{false, false}, OUT: true},
		{IN: []bool{true, false}, OUT: true},
		{IN: []bool{false, true}, OUT: true},
		{IN: []bool{true, true}, OUT: false},
	})
}
