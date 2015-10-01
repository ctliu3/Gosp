package value

type Symbol struct {
  Value string
}

func NewSymbol(val string) *Symbol {
  return &Symbol{Value: val}
}

func (self *Symbol) String() string {
  return self.Value
}
