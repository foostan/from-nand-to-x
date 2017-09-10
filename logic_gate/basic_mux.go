package logic_gate

type BasicMux struct {
}

func (n BasicMux) Calc(a bool, b bool, sel bool) (out bool) {

	if sel == false {
		out = a
	} else {
		out = b
	}

	return
}
