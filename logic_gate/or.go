package logic_gate

type Or struct {
}

func (n Or) Calc(a bool, b bool) (out bool) {
	not_a := Not{}.Calc(a)
	not_b := Not{}.Calc(b)
	return Nand{}.Calc(not_a, not_b)
}
