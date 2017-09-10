package logic_gate

type BasicNot struct{}

func (n BasicNot) Calc(a bool) (out bool) {
	return !a
}

type Not struct{}

func (n Not) Calc(a bool) (out bool) {
	return BasicNand{}.Calc(a, a)
}
