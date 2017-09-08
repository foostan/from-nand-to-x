package logic_gate

type SingleInputLogicGate interface {
  Calc(a bool) bool
}

type DoubleInputLogicGate interface {
  Calc(a bool, b bool) bool
}
