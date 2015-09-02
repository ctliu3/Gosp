package value

import (
  "fmt"
)

type Int struct {
  Value int64
}

func NewInt(val int64) *Int {
  return &Int{Value: val}
}

func (self *Int) String() string {
  return fmt.Sprintf("%d", self.Value)
}
