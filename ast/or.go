package ast

import (
  "github.com/ctliu3/gosp/scope"
  "github.com/ctliu3/gosp/value"
  const_ "github.com/ctliu3/gosp/constant"
)

// (or ⟨test1⟩ ...)

type Or struct {
  tests []Node
}

func NewOr(tests []Node) *Or {
  return &Or{tests}
}

func (self *Or) Type() string {
  return const_.OR
}

func (self *Or) Eval(env *scope.Scope) value.Value {
  var res value.Value

  for _, test := range self.tests {
    ret := test.Eval(env)
    if val, ok := ret.(*value.Bool); ok {
      if val.Value {
        return val
      }
    }
    res = ret
  }

  return res
}

func (self *Or) String() string {
  return "or"
}

func (self *Or) ExtRep() string {
  return ""
}
