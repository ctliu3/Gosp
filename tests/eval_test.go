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
    `(define abc 3) (+ abc 0)`: `3`,
    `(define a (+ 8 2)) (+ a 0)`: `10`,
  }
  check(t, mp)
}

//func TestDefineProc(t *testing.T) {
  //mp := map[string]string {
    //`(define (f x  y) (+ x y)) (f 1 2)`: `3`,
  //}
  //check(t, mp)
//}

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

func TestGT(t *testing.T) {
  mp := map[string]string {
    `(> 2 3)`: `#f`,
  }
  check(t, mp)
}

func TestIf(t *testing.T) {
  mp := map[string]string {
    `(if 1 0 1)`: `0`,
    `(if (> 2 3) 0 1)`: `1`,
    `(if (> 2 3 1) 0 1)`: `1`,
    `(if (if (> 2 3) 0 1) 3 4)`: `3`,
  }
  check(t, mp)
}

func TestLambda(t *testing.T) {
  mp := map[string]string {
    `(define f (lambda (x y) (+ x y))) (f 3 4)`: `7`,
    //`((lambda (x) (+ x x)) 4)`,
  }
  check(t, mp)
}

func TestLet(t *testing.T) {
  mp := map[string]string {
    `(let ((x 2) (y 3)) (+ x y))`: `5`,
    `(let ((x (+ 1 2)) (y 3)) (+ x y))`: `6`,
    `(let ((x 2) (y 3)) (let ((x 7) (z (+ x y))) (- z x)))`: `-2`,
  }
  check(t, mp)
}

func TestLetStar(t *testing.T) {
  mp := map[string]string {
    `(let ((x 2) (y 3)) (let* ((x 7) (z (+ x y))) (- z x)))`: `3`,
  }
  check(t, mp)
}
