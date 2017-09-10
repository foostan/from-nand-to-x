package logic_gate

type BasicXor struct {
}

func (n BasicXor) Calc(a bool, b bool) (out bool) {

	if a != b {
		out = true
	} else {
		out = false
	}

	return
}
