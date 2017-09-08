package logic_gate

type BasicNand struct {
}

func (n BasicNand) Calc(a bool, b bool) (out bool) {

	if a == b && a == true {
		out = false
	} else {
		out = true
	}

	return
}
