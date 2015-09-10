package value

import (
  "fmt"
)

type Char struct {
  Value string
}

func NewChar(val string) *Char {
  return &Char{Value: val}
}

func (self *Char) String() string {
  return fmt.Sprintf("%v", self.Value)
}
