package logic_gate

type Mux struct {
}

func (n Mux) Calc(a bool, b bool, sel bool) (out bool) {
	w1 := And{}.Calc(a, Not{}.Calc(sel))
	w2 := And{}.Calc(b, sel)
	return Or{}.Calc(w1, w2)
}
