package builtin

import (
  //"fmt"
  "runtime"
  "github.com/ctliu3/gosp/value"
)

// (max-procs obj)

type MaxProcs struct {
  value.Proc
}

func NewMaxProcs() *MaxProcs {
  return &MaxProcs{}
}

func (self *MaxProcs) Call(args ...value.Value) value.Value {
  if num, ok := args[0].(*value.Int); !ok {
    panic("max-procs: unexpeced int?")
  } else {
    runtime.GOMAXPROCS(int(num.Value))
  }
  return nil
}

// (>! channel "hello world")
type Send struct {
  value.Proc
}

func NewSend() *Send {
  return &Send{}
}

func (self *Send) Call(args ...value.Value) value.Value {
  if channel, ok := args[0].(*value.Chan); !ok {
    panic(">!: unexpeced chan?")
  } else {
    channel.Send(args[1])
  }
  return nil
}

// (<! channel)
type Recv struct {
  value.Proc
}

func NewRecv() *Recv {
  return &Recv{}
}

func (self *Recv) Call(args ...value.Value) value.Value {
  if channel, ok := args[0].(*value.Chan); !ok {
    panic("<!: unexpeced chan?")
  } else {
    return channel.Recv()
  }
  return nil
}

// (close! channel)
type Close struct {
  value.Proc
}

func NewClose() *Close {
  return &Close{}
}

func (self *Close) Call(args ...value.Value) value.Value {
  if channel, ok := args[0].(*value.Chan); !ok {
    panic("close!: unexpeced chan?")
  } else {
    channel.Close()
  }
  return nil
}
