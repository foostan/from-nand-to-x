package logic_gate

type BasicNand struct{}

func (n BasicNand) Calc(a bool, b bool) (out bool) {
	if a == b && a == true {
		return false
	} else {
		return true
	}
}

type Nand struct{}

func (n Nand) Calc(a bool, b bool) (out bool) {
	return BasicNand{}.Calc(a, b)
}
