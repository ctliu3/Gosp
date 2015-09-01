package main

import (
  "fmt"
  "os"
  "bufio"

  . "github.com/ctliu3/gosp/repl"
)

const version = "0.1"

func main() {
  fmt.Printf("Welcome to Gosp v.%s\n", version)

  for {
    fmt.Printf("> ")
    r := bufio.NewReader(os.Stdin)
    expr, _, err := r.ReadLine()
    if err != nil {
      fmt.Println("error, try again")
      continue
    }
    res := REPL(string(expr))
    fmt.Println(res)

    defer func() {
      if err := recover(); err != nil {
        fmt.Println("panic")
      }
    }()
  }
}
