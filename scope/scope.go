package scope

type Scope struct {
  outer *Scope
  objects map[string]*Object
}

type ObjectType int

const (
  Bad ObjectType = iota
  Fun
  Var
)

type Object struct {
  Type ObjectType
  Data interface{}
}

func NewScope() *Scope {
  return &Scope{nil, make(map[string]*Object)}
}

func (self *Scope) Lookup(ident string) bool {
  if _, ok := self.objects[ident]; ok {
    return true
  }
  return false
}

// Insert
