package value

type String struct {
  value string
}

func (self *String) String() string {
  return self.value
}
