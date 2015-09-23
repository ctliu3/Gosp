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
    return value.NewQuote(const_.BOOL)
  case *value.Int:
    return value.NewQuote(const_.INT)
  case *value.Float:
    return value.NewQuote(const_.FLOAT)
  case *value.Char:
    return value.NewQuote(const_.CHAR)
  case *value.List:
    return value.NewQuote(const_.LIST)
  case *value.String:
    return value.NewQuote(const_.STRING)
  case *value.Quote:
    return value.NewQuote(const_.QUOTE)
  case *value.Closure:
    return value.NewQuote(const_.PROC)
  }
  return nil
}

func (self *ClassOf) String() string {
  return "<#procedurce:class-of>"
}
