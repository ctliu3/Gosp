package ast

import (
  "fmt"
  "github.com/ctliu3/gosp/scope"
  "github.com/ctliu3/gosp/value"
  const_ "github.com/ctliu3/gosp/constant"
)

// (set! ⟨variable⟩ ⟨expression⟩) 

type Set struct {
  var_ Node
  expr Node // expr can be int/float or lambda
}

func NewSet(var_ Node, expr Node) *Set {
  return &Set{var_, expr}
}

func (self *Set) Type() string {
  return const_.SET
}

// This is a set manipulation.
func (self *Set) Eval(env *scope.Scope) value.Value {
  fmt.Println("Set#Eval")
  var_ := self.var_.(*Ident)
  val := self.expr.Eval(env)
  if obj := env.Lookup(var_.Name, true); obj == nil {
    panic("set!: assignment disallowed.")
  }
  env.Insert(var_.Name, scope.NewObj(val))

  // Define does not have return val.
  return nil
}

func (self *Set) ExtRep() string {
  return ""
}
