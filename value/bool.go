package value

import (
  const_ "github.com/ctliu3/gosp/constant"
)

type Bool struct {
  Value bool
}

func NewBool(val bool) *Bool {
  return &Bool{Value: val}
}

func (self *Bool) String() string {
  if self.Value {
    return const_.TRUE
  }
  return const_.FALSE
}
