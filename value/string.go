package value

type String struct {
  Value string
}

func NewString(val string) *String {
  return &String{Value: val}
}

func (self *String) String() string {
  return self.Value
}
