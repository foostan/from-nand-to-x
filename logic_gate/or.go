package logic_gate

type BasicOr struct{}

func (n BasicOr) Calc(a bool, b bool) (out bool) {
	if a == true || b == true {
		return true
	} else {
		return false
	}
}

type Or struct{}

func (n Or) Calc(a bool, b bool) (out bool) {
	not_a := Not{}.Calc(a)
	not_b := Not{}.Calc(b)
	return Nand{}.Calc(not_a, not_b)
}
