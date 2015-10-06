package scope

import (
  "fmt"
  "github.com/ctliu3/gosp/value"
  "github.com/ctliu3/gosp/value/builtin"
)

type Scope struct {
  outer *Scope
  objects map[string]*Object
}

type ObjectType int

const (
  Bad ObjectType = iota
  Var
  Proc
)

type Object struct {
  Type ObjectType
  Data interface{}
}

func NewObj(data value.Value) *Object {
  return &Object{Data: data}
}

func NewEmtpyScope(outer *Scope) *Scope {
  return &Scope{nil, make(map[string]*Object)}
}

func NewScope(outer *Scope) *Scope {
  if outer != nil {
    return &Scope{outer, make(map[string]*Object)}
  }

  objs := map[string]*Object {
    "+": NewObj(builtin.NewAdd()),
    "-": NewObj(builtin.NewSub()),
    "*": NewObj(builtin.NewMul()),
    "/": NewObj(builtin.NewDiv()),
    "remainder": NewObj(builtin.NewRemainder()),

    "<": NewObj(builtin.NewLT()),
    ">": NewObj(builtin.NewGT()),
    "<=": NewObj(builtin.NewLE()),
    ">=": NewObj(builtin.NewGE()),
    "==": NewObj(builtin.NewEQ()),

    ">!": NewObj(builtin.NewSend()),
    "<!": NewObj(builtin.NewRecv()),
    "close!": NewObj(builtin.NewClose()),
    "max-procs": NewObj(builtin.NewMaxProcs()),

    "vector-length": NewObj(builtin.NewVectLen()),
    "vector-ref": NewObj(builtin.NewVectRef()),

    "display": NewObj(builtin.NewDisplay()),
    "newline": NewObj(builtin.NewNewline()),

    "sleep": NewObj(builtin.NewSleep()),
    "eqv?": NewObj(builtin.NewIsEqv()),
    "class-of": NewObj(builtin.NewClassOf()),
  }
  return &Scope{nil, objs}
}

func (self *Scope) Insert(name string, obj *Object) {
  self.objects[name] = obj
}

func (self *Scope) Lookup(name string, recur bool) *Object {
  //fmt.Println("#Lookup, " + name)
  if !recur {
    return self.objects[name]
  }
  env := self
  for env != nil {
    if alt := env.objects[name]; alt != nil {
      return alt
    }
    env = env.outer
  }
  return nil
}

func (self *Scope) Size() int {
  size := len(self.objects)
  if self.outer != nil {
    size += self.outer.Size()
  }
  return size
}

func (self *Scope) DisplayDefine() {
  if self == nil {
    return
  }

  self.outer.DisplayDefine()
  for key := range self.objects {
    fmt.Printf("%v, ", key)
  }
  fmt.Println()
}
