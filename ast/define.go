package ast

import (
  "fmt"
  "github.com/ctliu3/gosp/scope"
  "github.com/ctliu3/gosp/value"
  const_ "github.com/ctliu3/gosp/constant"
)

// 1. (define ⟨variable⟩ ⟨expression⟩)
// 2. (define (⟨variable⟩ ⟨formals⟩) ⟨body⟩)
// where 2. equals to
// (define ⟨variable⟩
//    (lambda (⟨formals⟩) ⟨body⟩))

type Define struct {
  var_ Node
  expr Node // expr can be int/float or lambda
}

func NewDefine(var_ Node, expr Node) *Define {
  return &Define{var_, expr}
}

func (self *Define) Type() string {
  return const_.DEFINE
}

// This is a set manipulation.
func (self *Define) Eval(env *scope.Scope) value.Value {
  fmt.Println("Define#Eval")
  var_ := self.var_.(*Ident)
  val := self.expr.Eval(env)
  env.Insert(var_.Name, scope.NewObj(val))

  // Define does not have return val.
  return nil
}

func (self *Define) ExtRep() string {
  return ""
}
