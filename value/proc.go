package value

import (
  "fmt"
)

type Builtin interface {
  Call(...Value) Value
}

type Proc struct {
  name string
  //Builtin
}

func NewProc(name string) *Proc {
  return &Proc{name: name}
}

func (self *Proc) String() string {
  return fmt.Sprintf("<#procedure: %v>", self.name)
}
