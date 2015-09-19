package ast

import (
  "fmt"
  "github.com/ctliu3/gosp/scope"
  "github.com/ctliu3/gosp/value"
  const_ "github.com/ctliu3/gosp/constant"
)

// (let ⟨bindings⟩ ⟨body⟩)

type LetStar struct {
  bindings Binds
  body []Node
}

func NewLetStar(bindings Binds, body []Node) *LetStar {
  return &LetStar{bindings, body}
}

func (self *LetStar) Type() string {
  return const_.LET_STAR
}

func (self *LetStar) Eval(env *scope.Scope) value.Value {
  fmt.Println("#LetStar#Eval")
  local := scope.NewScope(env)
  for _, bind := range self.bindings.Bindings {
    local.Insert(bind.var_, scope.NewObj(bind.init.Eval(local)))
  }

  var ret value.Value
  for _, node := range self.body {
    ret = node.Eval(local)
  }
  return ret
}

func (self *LetStar) String() string {
  return "let*"
}

func (self *LetStar) ExtRep() string {
  return ""
}
