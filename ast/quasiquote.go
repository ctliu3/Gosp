package ast

import (
  "fmt"
  "reflect"

  "github.com/ctliu3/gosp/scope"
  "github.com/ctliu3/gosp/value"
  const_ "github.com/ctliu3/gosp/constant"
)

// (quasiquote ⟨qq template⟩)
// ⟨qq template⟩

type QuasiQuote struct {
  template Node
}

func NewQuasiQuote(template Node) *QuasiQuote {
  return &QuasiQuote{template}
}

func (self *QuasiQuote) Type() string {
  return const_.QUASIQUOTE
}

func (self *QuasiQuote) Eval(env *scope.Scope) value.Value {
  //self.evalQuasiQuote(env, 0)
  fmt.Println(reflect.TypeOf(self.template))

  switch val := self.template.(type) {
  case *Ident:
    return value.NewSymbol(val.ExtRep())
  case *List:
    return self.template.Eval(env)
  default:
    panic("unexpected expression")
  }
}

func (self *QuasiQuote) String() string {
  return ""
}

func (self *QuasiQuote) ExtRep() string {
  return ""
  //if self.template.Type() != const_.QUOTE {
    //return self.template.ExtRep()
  //}
  //return "(quote " + self.template.ExtRep() +  ")"
}

//func (self *QuasiQuote) SyntaxTree() {
//}
