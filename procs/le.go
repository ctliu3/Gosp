package procs

import (
  "github.com/ctliu3/gosp/value"
)

type LE struct {
}

func NewLE() *LE {
  return &LE{}
}

func (self *LE) Call(args ...value.Value) value.Value {
  operands, types := getOperands(args...)

  compi := func(lhs int64, rhs int64) bool { return lhs <= rhs }
  compf := func(lhs float64, rhs float64) bool { return lhs <= rhs }
  comps := func(lhs string, rhs string) bool { return lhs <= rhs }

  return check(operands, types, compi, compf, comps)
}
