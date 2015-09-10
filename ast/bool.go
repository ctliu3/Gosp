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
  if len(name) != 2 || (name[1] != 't' && name[1] != 'f') {
    panic(fmt.Errorf("bad syntax `%q", name))
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
  return fmt.Sprintf("%v", self.value)
}
