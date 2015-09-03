package procs

import (
  "github.com/ctliu3/gosp/value"
)

type Proc interface {
  Call(...value.Value) value.Value
}
