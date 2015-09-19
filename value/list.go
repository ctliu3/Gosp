package value

import (
  "bytes"
)

type List struct {
  Value []Value
}

func NewList(val []Value) *List {
  return &List{Value: val}
}

func (self *List) String() string {
  var buf bytes.Buffer

  buf.WriteRune('(')
  for i, val := range self.Value {
    if i > 0 {
      buf.WriteRune(' ')
    }
    buf.WriteString(val.String())
  }
  buf.WriteRune(')')

  return buf.String()
}
