package logic_gate

type BasicNor struct{}

func (n BasicNor) Calc(a bool, b bool) (out bool) {
	if a == true || b == true {
		return false
	} else {
		return true
	}
}

type Nor struct{}

func (n Nor) Calc(a bool, b bool) (out bool) {
	return Not{}.Calc(Or{}.Calc(a, b))
}
