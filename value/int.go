package value

import (
  "fmt"
)

type Int struct {
  Value int64
}

func (self Int) String() string {
  return fmt.Sprintf("%d", self.Value)
}
