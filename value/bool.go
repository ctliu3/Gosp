package value

type Bool struct {
  Value bool
}

func NewBool(val bool) *Bool {
  return &Bool{Value: val}
}

func (self *Bool) String() string {
  if self.Value {
    return "#t"
  }
  return "#f"
}
