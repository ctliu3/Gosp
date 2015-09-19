package ast

import (
  "github.com/ctliu3/gosp/scope"
  "github.com/ctliu3/gosp/value"
  const_ "github.com/ctliu3/gosp/constant"
)

// (let ⟨bindings⟩ ⟨body⟩)

type Letrec struct {
  bindings Binds
  body []Node
}

func NewLetrec(bindings Binds, body []Node) *Letrec {
  return &Letrec{bindings, body}
}

func (self *Letrec) Type() string {
  return const_.LETREC
}

func (self *Letrec) Eval(env *scope.Scope) value.Value {
  local := scope.NewScope(env)
  nbind := len(self.bindings.Bindings)

  for i := 0; i < nbind; i++ {
    bind := self.bindings.Bindings[i]
    local.Insert(bind.var_, scope.NewObj(bind.init.Eval(local)))
  }
  for i := nbind - 1; i >= 0; i-- {
    bind := self.bindings.Bindings[i]
    local.Insert(bind.var_, scope.NewObj(bind.init.Eval(local)))
  }

  var ret value.Value
  for _, node := range self.body {
    ret = node.Eval(local)
  }
  return ret
}

func (self *Letrec) String() string {
  return "letrec"
}

func (self *Letrec) ExtRep() string {
  return ""
}
