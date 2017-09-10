package logic_gate

type MuxIf interface {
	Calc(a bool, b bool, sel bool) bool
}

type BasicMux struct{}

func (n BasicMux) Calc(a bool, b bool, sel bool) (out bool) {
	if sel == false {
		return a
	} else {
		return b
	}
}

type Mux struct{}

func (n Mux) Calc(a bool, b bool, sel bool) (out bool) {
	w1 := And{}.Calc(a, Not{}.Calc(sel))
	w2 := And{}.Calc(b, sel)
	return Or{}.Calc(w1, w2)
}
