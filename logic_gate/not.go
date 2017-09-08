package logic_gate

type Not struct {
}

func (n Not) Calc(a bool) (out bool) {
	return BasicNand{}.Calc(a, a)
}
