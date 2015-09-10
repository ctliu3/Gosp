package main

import (
  "fmt"
  "os"
  "bufio"

  . "github.com/ctliu3/gosp/eval"
)

const version = "0.1"

func run() {
  for {
    defer func() {
      if err := recover(); err != nil {
        fmt.Printf("%v\n", err)
      }
      run()
    }()

    fmt.Printf("> ")
    r := bufio.NewReader(os.Stdin)
    expr, _, err := r.ReadLine()
    if err != nil {
      fmt.Println("error, try again")
      continue
    }
    res := Eval(string(expr))
    fmt.Println(res)
  }
}

func main() {
  fmt.Printf("Welcome to Gosp v.%s\n", version)

  run()
}
