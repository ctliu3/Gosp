package scope

import (
  "fmt"
  "github.com/ctliu3/gosp/value"
  "github.com/ctliu3/gosp/procs"
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

func NewObj(data interface{}) *Object {
  switch data.(type) {
  case value.Value: // implement String() method
    return &Object{Var, data}
  case procs.Proc: // implement Call() method
    return &Object{Proc, data}
  }
  return nil
}

func NewEmtpyScope(outer *Scope) *Scope {
  return &Scope{nil, make(map[string]*Object)}
}

func NewScope(outer *Scope) *Scope {
  objs := map[string]*Object {
    "+": NewObj(procs.NewAdd()),
    "-": NewObj(procs.NewSub()),
    "<": NewObj(procs.NewLT()),
    ">": NewObj(procs.NewGT()),
    "<=": NewObj(procs.NewLE()),
    ">=": NewObj(procs.NewGE()),
    "==": NewObj(procs.NewEQ()),
  }

  return &Scope{outer, objs}
}

func (self *Scope) Insert(name string, obj *Object) {
  self.objects[name] = obj
}

func (self *Scope) Lookup(name string, recur bool) *Object {
  fmt.Println("#Lookup, " + name)
  if !recur {
    return self.objects[name]
  }
  env := self
  for env != nil {
    if alt := env.objects[name]; alt != nil {
      return alt
    }
    env = self.outer
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
