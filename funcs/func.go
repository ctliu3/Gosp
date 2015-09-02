package funcs

import (
  "github.com/ctliu3/gosp/value"
)

type Func interface {
  Call(...value.Value) value.Value
}
