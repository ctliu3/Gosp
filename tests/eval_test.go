package tests

import (
  "testing"
  . "github.com/ctliu3/gosp/eval"
)

func check(t *testing.T, mp map[string]string) {
  for expr, exprRes := range mp {
    res := Eval(expr)
    if res != exprRes {
      t.Errorf("unexpeced result. expr = %q, return = %q\n", expr, res)
    }
  }
}

func TestDefineValue(t *testing.T) {
  mp := map[string]string {
    //`(define abc 3) (+ abc 0)`: `3`,
    `(define a (+ 8 2)) (+ a 0)`: `10`,
  }
  check(t, mp)
}

func TestAdd(t *testing.T) {
  mp := map[string]string {
    `(define a 3) (define b 4) (+ a b)`: `7`,
  }
  check(t, mp)
}

func TestSub(t *testing.T) {
  mp := map[string]string {
    `(define a 3) (define b 4) (- a b)`: `-1`,
  }
  check(t, mp)
}
