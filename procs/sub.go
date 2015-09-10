package procs

import (
  "github.com/ctliu3/gosp/value"
)

type Sub struct {
}

func NewSub() *Sub {
  return &Sub{}
}

func (self *Sub) Call(args ...value.Value) value.Value {
  var result float64 = 0
  isfloat := false

  for i, item := range args {
    switch item.(type) {
    case *value.Int:
      result -= float64(item.(*value.Int).Value)
    case *value.Float:
      isfloat = true
      result -= item.(*value.Float).Value
    default:
      panic("expected: number?")
    }
    if i == 0 {
      result = -result
    }
  }
  if isfloat {
    return value.NewFloat(result)
  }
  return value.NewInt(int64(result))
}
