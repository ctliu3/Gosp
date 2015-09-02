package funcs

import (
  "github.com/ctliu3/gosp/value"
)

type Sub struct {
}

func NewSub() *Sub {
  return &Sub{}
}

func (self *Sub) Call(args ...value.Value) value.Value {
  var res float64 = 0
  isfloat := false

  for i, item := range args {
    switch item.(type) {
    case *value.Int:
      res -= float64(item.(*value.Int).Value)
    case *value.Float:
      isfloat = true
      res -= item.(*value.Float).Value
    }
    if i == 0 {
      res = -res
    }
  }
  if isfloat {
    return value.NewFloat(res)
  }
  return value.NewInt(int64(res))
}
