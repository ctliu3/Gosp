package value

import (
  "bytes"
)

type List struct {
  head *ListNode
  len_ int
}

type ListNode struct {
  val Value
  next *ListNode
}

func NewList(nodes []Value) *List {
  if len(nodes) == 0 {
    return &List{nil, 0}
  }

  var dummy ListNode
  cur := &dummy
  for _, node := range nodes {
    cur.next = &ListNode{node, nil}
    cur = cur.next;
  }

  return &List{dummy.next, len(nodes)}
}

func (self *List) Car() Value {
  if self.head != nil {
    return self.head.val
  }
  return nil
}

func (self *List) Cdr() Value {
  if self.head.next != nil {
    return &List{self.head.next, self.len_ - 1}
  }
  return &List{nil, 0}
}

func (self *List) Empty() bool {
  if self.head == nil {
    return true
  }
  return false
}

func (self *List) String() string {
  var buf bytes.Buffer

  buf.WriteRune('(')
  for cur := self.head; cur != nil; cur = cur.next {
    if cur != self.head {
      buf.WriteRune(' ')
    }
    buf.WriteString(cur.val.String())
  }
  buf.WriteRune(')')

  return buf.String()
}
