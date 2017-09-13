package logic_gate

type BasicOr16 struct{}

func (n BasicOr16) Calc(a [16]bool, b[16]bool) (out [16]bool) {
	for i:=0; len(a) > i; i++ {
		out[i] = a[i] || b[i]
	}
	return
}

type Or16 struct{}

func (n Or16) Calc(a [16]bool, b[16]bool) (out [16]bool) {
	out[0] = Or{}.Calc(a[0], b[0])
	out[1] = Or{}.Calc(a[1], b[1])
	out[2] = Or{}.Calc(a[2], b[2])
	out[3] = Or{}.Calc(a[3], b[3])
	out[4] = Or{}.Calc(a[4], b[4])
	out[5] = Or{}.Calc(a[5], b[5])
	out[6] = Or{}.Calc(a[6], b[6])
	out[7] = Or{}.Calc(a[7], b[7])
	out[8] = Or{}.Calc(a[8], b[8])
	out[9] = Or{}.Calc(a[9], b[9])
	out[10] = Or{}.Calc(a[10], b[10])
	out[11] = Or{}.Calc(a[11], b[11])
	out[12] = Or{}.Calc(a[12], b[12])
	out[13] = Or{}.Calc(a[13], b[13])
	out[14] = Or{}.Calc(a[14], b[14])
	out[15] = Or{}.Calc(a[15], b[15])
	return
}
