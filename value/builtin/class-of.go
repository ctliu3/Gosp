package builtin

import (
  "fmt"
  "github.com/ctliu3/gosp/value"
  const_ "github.com/ctliu3/gosp/constant"
)

type ClassOf struct {
}

func NewClassOf() *ClassOf {
  return &ClassOf{}
}

func (self *ClassOf) Call(args ...value.Value) value.Value {
  if len(args) != 1 {
    panic(fmt.Errorf(
      "class-of: arguments number does not match, expected: 1 given: %d", len(args)))
  }

  switch args[0].(type) {
  case *value.Bool:
    return value.NewSymbol(const_.BOOL)
  case *value.Int:
    return value.NewSymbol(const_.INT)
  case *value.Float:
    return value.NewSymbol(const_.FLOAT)
  case *value.Char:
    return value.NewSymbol(const_.CHAR)
  case *value.List:
    return value.NewSymbol(const_.LIST)
  case *value.String:
    return value.NewSymbol(const_.STRING)
  case *value.Symbol:
    return value.NewSymbol(const_.SYMBOL)
  case *value.Closure:
    return value.NewSymbol(const_.PROC)
  }
  return nil
}

func (self *ClassOf) String() string {
  return "<#procedurce:class-of>"
}
