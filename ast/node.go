package ast

import (
  "github.com/ctliu3/gosp/scope"
  "github.com/ctliu3/gosp/value"
)

type Node interface {
  Type() string
  Eval(env *scope.Scope) value.Value
}
