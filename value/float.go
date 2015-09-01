package value

import (
  "fmt"
)

type Float struct {
  Value float64
}

func (self Float) String() string {
  return fmt.Sprintf("%d", self.Value)
}
