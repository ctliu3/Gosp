package builtin

import (
  "fmt"
  "github.com/ctliu3/gosp/value"
)

// (car list)

type Car struct {
  value.Proc
}

func NewCar() *Car {
  return &Car{}
}

func (self *Car) Call(args ...value.Value) value.Value {
  if len(args) != 1 {
    panic(fmt.Errorf("car: arity mismatch, expected 1, given %v\n", len(args)))
  }

  if arg, ok := args[0].(*value.List); !ok {
    panic(fmt.Errorf("car: contract violation, expected: list?"))
  } else {
    return arg.Car()
  }

  return nil
}

// (cdr list)

type Cdr struct {
  value.Proc
}

func NewCdr() *Cdr {
  return &Cdr{}
}

func (self *Cdr) Call(args ...value.Value) value.Value {
  if len(args) != 1 {
    panic(fmt.Errorf("cdr: arity mismatch, expected 1, given %v\n", len(args)))
  }

  if arg, ok := args[0].(*value.List); !ok {
    panic(fmt.Errorf("cdr: contract violation, expected: list?"))
  } else {
    return arg.Cdr()
  }

  return nil
}

// (null? obj)

type IsNull struct {
  value.Proc
}

func NewIsNull() *IsNull {
  return &IsNull{}
}

func (self *IsNull) Call(args ...value.Value) value.Value {
  if len(args) != 1 {
    panic(fmt.Errorf("null?: arity mismatch, expected 1, given %v\n", len(args)))
  }

  if arg, ok := args[0].(*value.List); !ok {
    panic(fmt.Errorf("null?: contract violation, expected: list?"))
  } else {
    if arg.Empty() {
      return value.NewBool(true)
    } else {
      return value.NewBool(false)
    }
  }

  return nil
}
