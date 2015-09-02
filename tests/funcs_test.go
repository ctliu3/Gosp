package tests

import (
  "testing"
  . "github.com/ctliu3/gosp/funcs"
  . "github.com/ctliu3/gosp/value"
)

func TestAdd(t *testing.T) {
  args := []Value{NewInt(5), NewFloat(6.0)}
  fun := NewAdd()

  res := fun.Call(args...)
  val, ok := res.(*Float)
  if !ok {
    t.Error("add function: return type error")
  } else if val.Value != 11.0 {
    t.Error("add function: return value error")
  }
}

func TestSub(t *testing.T) {
  args := []Value{NewInt(-9), NewFloat(6.4)}
  fun := NewSub()

  res := fun.Call(args...)
  val, ok := res.(*Float)
  if !ok {
    t.Error("add function: return type error")
  } else if val.Value != -15.4 {
    t.Error("add function: return value error")
  }
}
