package logic_gate

type Nand struct {
}

func (n Nand) Calc(a bool, b bool) bool {

	if a == b && a == true {
		return false
	} else {
		return true
	}
}
