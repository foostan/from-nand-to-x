package logic_gate

type BasicXor struct{}

func (n BasicXor) Calc(a bool, b bool) (out bool) {
	if a != b {
		return true
	} else {
		return false
	}
}

type Xor struct{}

func (n Xor) Calc(a bool, b bool) (out bool) {
	not_a := Not{}.Calc(a)
	not_b := Not{}.Calc(b)
	w1 := And{}.Calc(a, not_b)
	w2 := And{}.Calc(not_a, b)
	return Or{}.Calc(w1, w2)
}
