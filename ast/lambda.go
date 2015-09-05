package ast

import (
  //"fmt"
  //"strconv"
  //"strings"

  "github.com/ctliu3/gosp/scope"
  "github.com/ctliu3/gosp/value"
  const_ "github.com/ctliu3/gosp/constant"
)

// (lambda (⟨formals⟩) ⟨body⟩)

type Lambda struct {
  formals Node
  body Node
}

func NewLambda(formals Node, body Node) *Lambda {
  return &Lambda{formals, body}
}

func (self *Lambda) Type() string {
  return const_.LAMBDA
}

func (self *Lambda) Eval(env *scope.Scope) value.Value {
  return nil
}

func (self *Lambda) String() string {
  return const_.LAMBDA
}
