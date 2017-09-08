package logic_gate

type BasicNot struct {
}

func (n BasicNot) Calc(a bool) (out bool) {
	return !a
}
