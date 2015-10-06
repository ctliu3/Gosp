package builtin

import (
  "fmt"
  "github.com/ctliu3/gosp/value"
)

// (vector-length vector)

type VectLen struct {
  value.Proc
}

func NewVectLen() *VectLen {
  return &VectLen{}
}

func (self *VectLen) Call(args ...value.Value) value.Value {
  if vect, ok := args[0].(*value.Vect); !ok {
    panic(fmt.Errorf("vector-length: contract violation, expected: vector?"))
  } else {
    return vect.Len()
  }
  return nil
}

// (vector-ref vector)

type VectRef struct {
  value.Proc
}

func NewVectRef() *VectRef {
  return &VectRef{}
}

func (self *VectRef) Call(args ...value.Value) value.Value {
  if len(args) != 2 {
    panic(fmt.Errorf("vector-ref: arity mismatch, expected 2, given %v\n", len(args)))
  }

  var vect *value.Vect
  var index int64

  if arg, ok := args[0].(*value.Vect); !ok {
    panic(fmt.Errorf("vector-ref: contract violation, expected: vector?"))
  } else {
    vect = arg
  }
  if arg, ok := args[1].(*value.Int); !ok {
    panic(fmt.Errorf("vector-ref: contract violation, expected: int?"))
  } else {
    index = arg.Value
  }

  return vect.At(int(index))
}
