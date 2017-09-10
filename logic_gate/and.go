package logic_gate

type BasicAnd struct{}

func (n BasicAnd) Calc(a bool, b bool) (out bool) {
	if a == b && a == true {
		return true
	} else {
		return false
	}
}

type And struct{}

func (n And) Calc(a bool, b bool) (out bool) {
	return Not{}.Calc(Nand{}.Calc(a, b))
}
