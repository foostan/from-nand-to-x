package logic_gate

type Nor struct {
}

func (n Nor) Calc(a bool, b bool) (out bool) {
	return Not{}.Calc(Or{}.Calc(a, b))
}
