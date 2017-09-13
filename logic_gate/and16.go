package logic_gate

type BasicAnd16 struct{}

func (n BasicAnd16) Calc(a [16]bool, b[16]bool) (out [16]bool) {
	for i:=0; len(a) > i; i++ {
		out[i] = a[i] && b[i]
	}
	return
}

type And16 struct{}

func (n And16) Calc(a [16]bool, b[16]bool) (out [16]bool) {
	out[0] = And{}.Calc(a[0], b[0])
	out[1] = And{}.Calc(a[1], b[1])
	out[2] = And{}.Calc(a[2], b[2])
	out[3] = And{}.Calc(a[3], b[3])
	out[4] = And{}.Calc(a[4], b[4])
	out[5] = And{}.Calc(a[5], b[5])
	out[6] = And{}.Calc(a[6], b[6])
	out[7] = And{}.Calc(a[7], b[7])
	out[8] = And{}.Calc(a[8], b[8])
	out[9] = And{}.Calc(a[9], b[9])
	out[10] = And{}.Calc(a[10], b[10])
	out[11] = And{}.Calc(a[11], b[11])
	out[12] = And{}.Calc(a[12], b[12])
	out[13] = And{}.Calc(a[13], b[13])
	out[14] = And{}.Calc(a[14], b[14])
	out[15] = And{}.Calc(a[15], b[15])
	return
}
