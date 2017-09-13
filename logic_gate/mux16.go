package logic_gate

type Mux16If interface {
	Calc(a [16]bool, b [16]bool, sel bool) [16]bool
}

type BasicMux16 struct{}

func (n BasicMux16) Calc(a [16]bool, b [16]bool, sel bool) (out [16]bool) {
	if sel == false {
		return a
	} else {
		return b
	}
}

type Mux16 struct{}

func (n Mux16) Calc(a [16]bool, b [16]bool, sel bool) (out [16]bool) {
	out[0] = Or{}.Calc(And{}.Calc(a[0], Not{}.Calc(sel)), And{}.Calc(b[0], sel))
	out[1] = Or{}.Calc(And{}.Calc(a[1], Not{}.Calc(sel)), And{}.Calc(b[1], sel))
	out[2] = Or{}.Calc(And{}.Calc(a[2], Not{}.Calc(sel)), And{}.Calc(b[2], sel))
	out[3] = Or{}.Calc(And{}.Calc(a[3], Not{}.Calc(sel)), And{}.Calc(b[3], sel))
	out[4] = Or{}.Calc(And{}.Calc(a[4], Not{}.Calc(sel)), And{}.Calc(b[4], sel))
	out[5] = Or{}.Calc(And{}.Calc(a[5], Not{}.Calc(sel)), And{}.Calc(b[5], sel))
	out[6] = Or{}.Calc(And{}.Calc(a[6], Not{}.Calc(sel)), And{}.Calc(b[6], sel))
	out[7] = Or{}.Calc(And{}.Calc(a[7], Not{}.Calc(sel)), And{}.Calc(b[7], sel))
	out[8] = Or{}.Calc(And{}.Calc(a[8], Not{}.Calc(sel)), And{}.Calc(b[8], sel))
	out[9] = Or{}.Calc(And{}.Calc(a[9], Not{}.Calc(sel)), And{}.Calc(b[9], sel))
	out[10] = Or{}.Calc(And{}.Calc(a[10], Not{}.Calc(sel)), And{}.Calc(b[10], sel))
	out[11] = Or{}.Calc(And{}.Calc(a[11], Not{}.Calc(sel)), And{}.Calc(b[11], sel))
	out[12] = Or{}.Calc(And{}.Calc(a[12], Not{}.Calc(sel)), And{}.Calc(b[12], sel))
	out[13] = Or{}.Calc(And{}.Calc(a[13], Not{}.Calc(sel)), And{}.Calc(b[13], sel))
	out[14] = Or{}.Calc(And{}.Calc(a[14], Not{}.Calc(sel)), And{}.Calc(b[14], sel))
	out[15] = Or{}.Calc(And{}.Calc(a[15], Not{}.Calc(sel)), And{}.Calc(b[15], sel))
	return out
}
