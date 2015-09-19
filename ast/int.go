package ast

import (
  "fmt"
  "strconv"
  "strings"

  "github.com/ctliu3/gosp/scope"
  "github.com/ctliu3/gosp/value"
  const_ "github.com/ctliu3/gosp/constant"
)

type Int struct {
  value int64
  base int
}

func NewInt(name string) *Int {
  base := 10
  if strings.ContainsRune(name, '#') {
    switch name[1] {
    case 'b':
      base = 2
    case 'd':
      base = 10
    case 'x':
      base = 16
    }
  }

  val, err := strconv.ParseInt(name, base, 64)
  if err != nil {
    panic(err)
  }

  return &Int{value: val, base: base}
}

func (self *Int) Type() string {
  return const_.INT
}

func (self *Int) Eval(env *scope.Scope) value.Value {
  return value.NewInt(self.value)
}

func (self *Int) String() string {
  switch self.base {
  case 2:
    return fmt.Sprintf("%b", self.value)
  case 8:
    return fmt.Sprintf("%o", self.value)
  case 10:
    return fmt.Sprintf("%d", self.value)
  case 16:
    return fmt.Sprintf("%X", self.value)
  }
  return "error while handling String() of Int"
}

func (self *Int) ExtRep() string {
  return self.String()
}
