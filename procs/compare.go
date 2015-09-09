package procs

import (
  //"fmt"
  "github.com/ctliu3/gosp/value"
)

type GT struct {
}

func NewGT() *GT {
  return &GT{}
}

type compiType func(int64, int64) bool
type compfType func(float64, float64) bool
type compsType func(string, string) bool

func (self *GT) Call(args ...value.Value) value.Value {
  operands, types := getOperands(args...)

  compi := func(lhs int64, rhs int64) bool { return lhs > rhs }
  compf := func(lhs float64, rhs float64) bool { return lhs > rhs }
  comps := func(lhs string, rhs string) bool { return lhs > rhs }

  return check(operands, types, compi, compf, comps)
}

func check(operands []interface{}, types []int, compi compiType, compf compfType, comps compsType) value.Value {
  typ := types[0]

  for i, _ := range operands {
    if i == 0 {
      continue
    }
    switch typ {
      case 0: // int
      if !compi(operands[i - 1].(int64), operands[i].(int64)) {
        return value.NewBool(false)
      }
      case 1: // float
      if !compf(operands[i - 1].(float64), operands[i].(float64)) {
        return value.NewBool(false)
      }
      case 2: // string
      if !comps(operands[i - 1].(string), operands[i].(string)) {
        return value.NewBool(false)
      }
    }
  }

  return value.NewBool(true)
}

func getOperands(args ...value.Value) ([]interface{}, []int)  {
  nargs := len(args)
  operands := make([]interface{}, nargs)
  types := make([]int, nargs)

  for i, item := range args {
    switch item.(type) {
    case *value.Int:
      operands[i] = item.(*value.Int).Value
      types[i] = 0
    case *value.Float:
      operands[i] = item.(*value.Float).Value
      types[i] = 1
    case *value.String:
      operands[i] = item.(*value.String).Value
      types[i] = 2
    default:
      panic("unexpeced type")
    }
  }

  if types[0] == 2 {
    for _, typ := range types {
      if typ != types[0] {
        panic("type not match")
      }
    }
  }

  return operands, types
}
