package logic_gate

type BasicAnd struct {
}

func (n BasicAnd) Calc(a bool, b bool) (out bool) {

	if a == b && a == true {
		out = true
	} else {
		out = false
	}

	return
}
