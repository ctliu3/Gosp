package ast

import (
  "fmt"
  "github.com/ctliu3/gosp/scope"
  "github.com/ctliu3/gosp/value"
  const_ "github.com/ctliu3/gosp/constant"
)

// (quote ⟨datum⟩)
// ’⟨datum⟩
// ⟨constant⟩

type Quote struct {
  datum Node
}

func NewQuote(datum Node) *Quote {
  return &Quote{datum}
}

func (self *Quote) Type() string {
  return const_.QUOTE
}

func (self *Quote) Eval(env *scope.Scope) value.Value {
  fmt.Println("##")
  fmt.Println(self.ExtRep())
  return value.NewQuote(self.ExtRep())
}

func (self *Quote) String() string {
  return ""
}

func (self *Quote) ExtRep() string {
  _, ok := self.datum.(*Quote)
  if !ok {
    return self.datum.ExtRep()
  }
  return "(quote " + self.datum.ExtRep() +  ")"
}
