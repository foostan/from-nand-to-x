package logic_gate

type DMuxIf interface {
	Calc(in bool, sel bool) (a bool, b bool)
}

type BasicDMux struct{}

func (n BasicDMux) Calc(in bool, sel bool) (a bool, b bool) {
	if sel == false {
		a = in
	} else {
		b = in
	}
	return
}

type DMux struct{}

func (n DMux) Calc(in bool, sel bool) (a bool, b bool) {
	a = And{}.Calc(in, Not{}.Calc(sel))
	b = And{}.Calc(in, sel)
	return
}

