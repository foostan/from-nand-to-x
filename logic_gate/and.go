package logic_gate

type And struct {
}

func (n And) Calc(a bool, b bool) (out bool) {
	return Not{}.Calc(Nand{}.Calc(a, b))
}
