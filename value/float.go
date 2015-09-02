package value

import (
  "fmt"
)

type Float struct {
  Value float64
}

func NewFloat(val float64) *Float {
  return &Float{Value: val}
}

func (self *Float) String() string {
  return fmt.Sprintf("%d", self.Value)
}
