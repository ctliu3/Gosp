package value

import (
  "fmt"
)

type Chan struct {
  channel chan Value
  size int
}

func NewChan(size int) *Chan {
  return &Chan{
    size: size,
    channel: make(chan Value),
  }
}

func (self *Chan) Send(val Value) {
  self.channel <- val
}

func (self *Chan) Recv() Value {
  return <- self.channel
}

func (self *Chan) Close() {
  close(self.channel)
}

func (self *Chan) String() string {
  return fmt.Sprintf("#<channel>")
}
