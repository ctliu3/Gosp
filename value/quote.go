package value

import (
  "fmt"
)

type Quote struct {
  Value string
}

func NewQuote(val string) *Quote {
  return &Quote{Value: val}
}

func (self *Quote) String() string {
  return fmt.Sprintf("%q", self.Value)
}
