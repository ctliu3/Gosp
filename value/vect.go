package value

import (
  "fmt"
  "bytes"
)

type Vect struct {
  elements []Value
  len_ int
}

func NewVect(nodes []Value) *Vect {
  return &Vect{elements: nodes, len_: len(nodes)};
}

func (self *Vect) At(index int) Value {
  if index > self.len_ {
    panic(fmt.Errorf("vector-ref: index is out of range, index: %v, valid range: [0, %v]", index, self.len_ - 1))
  }
  return self.elements[index]
}

func (self *Vect) Len() Value {
  return NewInt(int64(self.len_))
}

func (self *Vect) String() string {
  var buf bytes.Buffer

  buf.WriteRune('[')
  for i, element := range self.elements {
    if i > 0 {
      buf.WriteRune(' ')
    }
    buf.WriteString(element.String())
  }
  buf.WriteRune(']')

  return buf.String()
}
