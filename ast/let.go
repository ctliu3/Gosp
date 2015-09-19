package ast

import (
  "github.com/ctliu3/gosp/scope"
  "github.com/ctliu3/gosp/value"
  const_ "github.com/ctliu3/gosp/constant"
)

// (let ⟨bindings⟩ ⟨body⟩)

type Let struct {
  bindings Binds
  body []Node
}

func NewLet(bindings Binds, body []Node) *Let {
  return &Let{bindings, body}
}

func (self *Let) Type() string {
  return const_.LET
}

func (self *Let) Eval(env *scope.Scope) value.Value {
  local := scope.NewScope(env)
  for _, bind := range self.bindings.Bindings {
    local.Insert(bind.var_, scope.NewObj(bind.init.Eval(env)))
  }

  var ret value.Value
  for _, node := range self.body {
    ret = node.Eval(local)
  }
  return ret
}

func (self *Let) String() string {
  return "let"
}

func (self *Let) ExtRep() string {
  return ""
}
