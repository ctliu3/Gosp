package main

import (
  "fmt"
  "os"
  "io/ioutil"
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

func evalFromFile(filename string) {
  file, err := os.Open(filename)
  if err != nil {
    panic(err)
  }
  defer file.Close()
  buffer, err := ioutil.ReadAll(file)
  if err != nil {
    panic(err)
  }

  Eval(string(buffer))
}

func errMsg() {
  fmt.Println(
    "\nUsage:\n" +
    "\t ./Gosp \t\t\t\t interactive mode\n" +
    "\t ./Gosp run [script-name] \t run gosp script")
}

func main() {
  fmt.Printf("Welcome to Gosp v.%s\n", version)

  narg := len(os.Args)
  if narg == 1 {
    run()
  } else if narg == 3 {
    switch os.Args[1] {
    case "run":
      evalFromFile(os.Args[2])
    default:
      errMsg()
    }
  } else {
    errMsg()
  }
}
