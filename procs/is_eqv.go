package procs

import (
  "fmt"
  "github.com/ctliu3/gosp/value"
)

type IsEqv struct {
}

func NewIsEqv() *IsEqv {
  return &IsEqv{}
}

func (self *IsEqv) Call(args ...value.Value) value.Value {
  if len(args) != 2 {
    panic(fmt.Errorf("eqv?: unexpeced arguments number, expected: 2, given %v", len(args)))
  }

  if args[0].String() != args[1].String() {
    return value.NewBool(false)
  }

  var ok bool
  switch args[0].(type) {
  case *value.Bool:
    _, ok = args[1].(*value.Bool)
  case *value.Int:
    _, ok = args[1].(*value.Int)
  case *value.Char:
    _, ok = args[1].(*value.Char)
  case *value.String:
    _, ok = args[1].(*value.String)
  case *value.Float:
    _, ok = args[1].(*value.Float)
  case *value.Quote:
    _, ok = args[1].(*value.Quote)
  case *value.List:
    _, ok = args[1].(*value.List)
  case *value.Closure:
    _, ok = args[1].(*value.Closure)
  }
  return value.NewBool(ok)
}
