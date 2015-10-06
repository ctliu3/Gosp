package builtin

import (
  "fmt"
  "time"
  "github.com/ctliu3/gosp/value"
)

// (sleep [secs])

type Sleep struct {
  value.Proc
}

func NewSleep() *Sleep {
  return &Sleep{}
}

func (self *Sleep) Call(args ...value.Value) value.Value {
  if len(args) != 1 {
    panic(fmt.Errorf("sleep: arity mismatch, expected 1, given %s\n", len(args)))
  }
  if sec, ok := args[0].(*value.Int); ok {
    time.Sleep(time.Duration(sec.Value) * time.Second)
  }
  return nil
}
