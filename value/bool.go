package value

import (
  "fmt"
)

type Bool struct {
  Value bool
}

func NewBool(val bool) *Bool {
  return &Bool{Value: val}
}

func (self *Bool) String() string {
  return fmt.Sprintf("%t", self.Value)
}
