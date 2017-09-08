package logic_gate

type Nand struct {
}

func (n Nand) Calc(a bool, b bool) (out bool) {

	if a == b && a == true {
		out = false
	} else {
		out = true
	}

	return
}
