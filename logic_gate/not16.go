package logic_gate

type BasicNot16 struct{}

func (n BasicNot16) Calc(a [16]bool) (out [16]bool) {
	for i:=0; len(a) > i; i++ {
		out[i] = !a[i]
	}
	return
}

type Not16 struct{}

func (n Not16) Calc(a [16]bool) (out [16]bool) {
	out[0] = Nand{}.Calc(a[0], a[0])
	out[1] = Nand{}.Calc(a[1], a[1])
	out[2] = Nand{}.Calc(a[2], a[2])
	out[3] = Nand{}.Calc(a[3], a[3])
	out[4] = Nand{}.Calc(a[4], a[4])
	out[5] = Nand{}.Calc(a[5], a[5])
	out[6] = Nand{}.Calc(a[6], a[6])
	out[7] = Nand{}.Calc(a[7], a[7])
	out[8] = Nand{}.Calc(a[8], a[8])
	out[9] = Nand{}.Calc(a[9], a[9])
	out[10] = Nand{}.Calc(a[10], a[10])
	out[11] = Nand{}.Calc(a[11], a[11])
	out[12] = Nand{}.Calc(a[12], a[12])
	out[13] = Nand{}.Calc(a[13], a[13])
	out[14] = Nand{}.Calc(a[14], a[14])
	out[15] = Nand{}.Calc(a[15], a[15])
	return
}
