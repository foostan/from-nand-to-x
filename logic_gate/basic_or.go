package logic_gate

type BasicOr struct {
}

func (n BasicOr) Calc(a bool, b bool) (out bool) {

	if a == true || b == true {
		out = true
	} else {
		out = false
	}

	return
}
