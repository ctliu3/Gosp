package ast

import (
  "github.com/ctliu3/gosp/scope"
  "github.com/ctliu3/gosp/value"
  const_ "github.com/ctliu3/gosp/constant"
)

// (quote ⟨datum⟩)
// ’⟨datum⟩
// ⟨constant⟩

type Quote struct {
  datum string
}

func NewQuote(datum string) *Quote {
  return &Quote{datum}
}

func (self *Quote) Type() string {
  return const_.QUOTE
}

func (self *Quote) Eval(env *scope.Scope) value.Value {
  if len(self.datum) > 0 && self.datum[0] == '\'' {
    quote := NewQuote(self.datum[1:])
    return value.NewQuote("(quote " + quote.Eval(env).String() +  ")")
  }
  return value.NewQuote(self.datum)
}

func (self *Quote) String() string {
  return self.datum
}

func (self *Quote) ExtRep() string {
  return ""
}
