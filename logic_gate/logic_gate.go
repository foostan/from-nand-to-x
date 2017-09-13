package logic_gate

type SingleInputLogicGate interface {
	Calc(a bool) (out bool)
}

type Single16InputLogicGate interface {
	Calc(a [16]bool) (out [16]bool)
}

type DoubleInputLogicGate interface {
	Calc(a bool, b bool) (out bool)
}

type Double16InputLogicGate interface {
	Calc(a [16]bool, b [16]bool) (out [16]bool)
}
