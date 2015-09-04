package ast

import (
  "github.com/ctliu3/gosp/scope"
  "github.com/ctliu3/gosp/value"
  const_ "github.com/ctliu3/gosp/constant"
)

type Char struct {
  value string
  chr byte
}

var chars = map[string]byte {
  "space": ' ',
  "newline": '\n',
}

func NewChar(name string) *Char {
  if len(name) > 3 {
    if chr, ok := chars[name[3:]]; ok {
      return &Char{value: name, chr: chr}
    } else {
      panic("unexcepted char?")
    }
  }
  return &Char{value: name, chr: name[2]}
}

func (self *Char) Type() string {
  return const_.CHAR
}

func (self *Char) Eval(env *scope.Scope) value.Value {
  return value.NewChar(self.value)
}

func (self *Char) String() string {
  return self.value
}
