package ast

import (
  "os"
  "io/ioutil"
  "fmt"

  "github.com/ctliu3/gosp/scope"
  "github.com/ctliu3/gosp/value"
  const_ "github.com/ctliu3/gosp/constant"
)

type Import struct {
  filename string
}

func NewImport(filename string) *Import{
  if _, err := os.Stat(filename); os.IsNotExist(err) {
    panic(fmt.Errorf("%v is not a valid file", filename))
  }
  return &Import{filename}
}

func (self *Import) Type() string {
  return const_.IMPORT
}

func (self *Import) Eval(env *scope.Scope) value.Value {
  file, err := os.Open(self.filename)
  if err != nil {
    panic(err)
  }
  defer file.Close()
  buffer, err := ioutil.ReadAll(file)
  if err != nil {
    panic(err)
  }

  //fmt.Println(string(buffer))
  return NewString(string(buffer))
}

func (self* Import) ExtRep() string {
  return ""
}
