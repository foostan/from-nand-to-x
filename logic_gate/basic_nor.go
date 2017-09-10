package logic_gate

type BasicNor struct {
}

func (n BasicNor) Calc(a bool, b bool) (out bool) {

	if a == true || b == true {
		out = false
	} else {
		out = true
	}

	return
}
