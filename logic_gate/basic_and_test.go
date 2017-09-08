package logic_gate

import (
	"testing"
)

func TestBasicAndCalc(t *testing.T) {
	test_calc(t, BasicAnd{}, []Expect{
		{IN: []bool{false, false}, OUT: false},
		{IN: []bool{true, false}, OUT: false},
		{IN: []bool{false, true}, OUT: false},
		{IN: []bool{true, true}, OUT: true},
	})
}
