package builtin

import (
  "fmt"
  "github.com/ctliu3/gosp/value"
)

// Add operation.
type Add struct {
  value.Proc
}

func NewAdd() *Add {
  return &Add{*value.NewProc("+")}
}

func (self *Add) Call(args ...value.Value) value.Value {
  var result float64 = 0
  isfloat := false

  for _, item := range args {
    switch operand := item.(type) {
    case *value.Int:
      result += float64(operand.Value)
    case *value.Float:
      isfloat = true
      result += operand.Value
    default:
      panic(fmt.Errorf("%v: unexpected number", self.String()))
    }
  }
  if isfloat {
    return value.NewFloat(result)
  }
  return value.NewInt(int64(result))
}

// Sub operation.
type Sub struct {
  value.Proc
}

func NewSub() *Sub {
  return &Sub{*value.NewProc("-")}
}

func (self *Sub) Call(args ...value.Value) value.Value {
  var result float64 = 0
  isfloat := false

  for i, item := range args {
    switch operand := item.(type) {
    case *value.Int:
      result -= float64(operand.Value)
    case *value.Float:
      isfloat = true
      result -= operand.Value
    default:
      panic(fmt.Errorf("%v: unexpected number", self.String()))
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

// Mult operation.
type Mul struct {
  value.Proc
}

func NewMul() *Mul {
  return &Mul{*value.NewProc("*")}
}

func (self *Mul) Call(args ...value.Value) value.Value {
  var result float64 = 1
  isfloat := false

  for _, item := range args {
    switch operand := item.(type) {
    case *value.Int:
      result *= float64(operand.Value)
    case *value.Float:
      isfloat = true
      result *= operand.Value
    default:
      panic(fmt.Errorf("%v: unexpected number", self.String()))
    }
  }
  if isfloat {
    return value.NewFloat(result)
  }
  return value.NewInt(int64(result))
}

// Div operation.
type Div struct {
  value.Proc
}

func NewDiv() *Div {
  return &Div{*value.NewProc("/")}
}

func (self *Div) Call(args ...value.Value) value.Value {
  var result float64

  for i, item := range args {
    switch operand := item.(type) {
    case *value.Int:
      if operand.Value == 0 {
        panic(fmt.Errorf("%v: division by zero", self.String()))
      }
      if i == 0 {
        result = float64(operand.Value)
      } else {
        result /= float64(operand.Value)
      }

    case *value.Float:
      if operand.Value == 0 {
        panic(fmt.Errorf("%v: division by zero", self.String()))
      }
      if i == 0 {
        result = operand.Value
      } else {
        result /= operand.Value
      }

    default:
      panic(fmt.Errorf("%v: unexpected number", self.String()))
    }
  }
  return value.NewFloat(result)
}

// remainder
type Remainder struct {
  value.Proc
}

func NewRemainder() *Remainder {
  return &Remainder{*value.NewProc("remainder")}
}

func (self *Remainder) Call(args ...value.Value) value.Value {
  if len(args) != 2 {
    panic(fmt.Errorf("remainder: arity mismatch, expected 2, given %s\n", len(args)))
  }

  var n1, n2 int64
  if arg1, ok := args[0].(*value.Int); !ok {
    panic(fmt.Errorf("remainder: contract violation, expected: integer?"))
  } else {
    n1 = arg1.Value
  }

  if arg2, ok := args[1].(*value.Int); !ok {
    panic(fmt.Errorf("remainder: contract violation, expected: integer?"))
  } else {
    n2 = arg2.Value
  }

  return value.NewInt(n1 % n2)
}
