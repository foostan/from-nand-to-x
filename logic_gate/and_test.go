package logic_gate

import (
	"testing"
)

func TestAndCalc(t *testing.T) {
	test_calc(t, And{}, []Expect{
		{IN: []bool{false, false}, OUT: false},
		{IN: []bool{true, false}, OUT: false},
		{IN: []bool{false, true}, OUT: false},
		{IN: []bool{true, true}, OUT: true},
	})
}
