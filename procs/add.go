package procs

import (
  "github.com/ctliu3/gosp/value"
)

type Add struct {
}

func NewAdd() *Add {
  return &Add{}
}

func (self *Add) Call(args ...value.Value) value.Value {
  var res float64 = 0
  isfloat := false

  for _, item := range args {
    switch item.(type) {
    case *value.Int:
      res += float64(item.(*value.Int).Value)
    case *value.Float:
      isfloat = true
      res += item.(*value.Float).Value
    }
  }
  if isfloat {
    return value.NewFloat(res)
  }
  return value.NewInt(int64(res))
}
