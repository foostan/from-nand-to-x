package logic_gate

type LogicGate interface {
  Calc(a bool, b bool) bool
}
