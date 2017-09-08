package logic_gate

type Nand struct {
}

func (n Nand) Calc(a bool, b bool) (out bool) {
	return BasicNand{}.Calc(a, b)
}
