package ast

import (
  "fmt"

  "github.com/ctliu3/gosp/scope"
  "github.com/ctliu3/gosp/value"
  const_ "github.com/ctliu3/gosp/constant"
)

type Bool struct {
  value bool
}

func NewBool(name string) *Bool {
  if len(name) != 2 {
    panic("bool parse error")
  }
  if name[1] == 't' {
    return &Bool{value: true}
  }
  return &Bool{value: false}
}

func (self *Bool) Type() string {
  return const_.BOOL
}

func (self *Bool) Eval(env *scope.Scope) value.Value {
  return value.NewBool(self.value)
}

func (self *Bool) String() string {
  return fmt.Sprintf("%d", self.value)
}
