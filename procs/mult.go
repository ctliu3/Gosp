package procs

import (
  "github.com/ctliu3/gosp/value"
)

type Mult struct {
}

func NewMult() *Mult {
  return &Mult{}
}

func (self *Mult) Call(args ...value.Value) value.Value {
  var result float64 = 1
  isfloat := false

  for _, item := range args {
    switch item.(type) {
    case *value.Int:
      result *= float64(item.(*value.Int).Value)
    case *value.Float:
      isfloat = true
      result *= item.(*value.Float).Value
    default:
      panic("expected: number?")
    }
  }
  if isfloat {
    return value.NewFloat(result)
  }
  return value.NewInt(int64(result))
}
