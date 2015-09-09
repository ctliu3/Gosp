package parser

import (
  "fmt"
  "strings"

  "github.com/ctliu3/gosp/ast"
  "github.com/ctliu3/gosp/lexer"
  const_ "github.com/ctliu3/gosp/constant"
)

type Parser struct {
  l *lexer.Lexer
}

// Syntactic parsing in two phases:
// 1. Parse the tokens into a set of tuples;
// 2. Parse the tuples into abstract syntactic tree (AST).
func ParseFromString(expr string) []ast.Node {
  parser := &Parser{lexer.Lex("gosp", expr)}
  tuples := make([]ast.Node, 0)
  parser.parse(&tuples, 0)

  // For debugging.
  for _, node := range tuples {
    printType(node)
    fmt.Println()
  }

  return parseList(&tuples)
}

func (self *Parser) parse(nodes *[]ast.Node, dep int) {
  for token := self.l.NextToken(); token.Type != lexer.TokenEOF; token = self.l.NextToken() {

    switch token.Type {
    case lexer.TokenBool:
      *nodes = append(*nodes, ast.NewBool(token.Name))
    case lexer.TokenChar:
      *nodes = append(*nodes, ast.NewChar(token.Name))
    case lexer.TokenString:
      *nodes = append(*nodes, ast.NewString(token.Name))
    case lexer.TokenQuote:
      *nodes = append(*nodes, ast.NewQuote(token.Name))
    case lexer.TokenIdent:
      *nodes = append(*nodes, ast.NewIdent(token.Name))
    case lexer.TokenNumber:
      if strings.ContainsAny(token.Name, "Ee.") {
        *nodes = append(*nodes, ast.NewFloat(token.Name))
      } else {
        *nodes = append(*nodes, ast.NewInt(token.Name))
      }
    case lexer.TokenLParen:
      sub := make([]ast.Node, 0)
      self.parse(&sub, dep + 1)
      *nodes = append(*nodes, ast.NewTuple(sub))
    case lexer.TokenRParen:
      return
    default:
      panic("unexpeced token")
    }
  }
}

func parseList(tuples *[]ast.Node) []ast.Node {
  var nodes []ast.Node
  for _, tuple := range *tuples {
    nodes = append(nodes, parseNode(tuple))
  }
  return nodes
}

// Each node has differnt type.
// 1. Tuple, the first element in tuple is the identifier (such as `if',
// `begin', `lambda', etc) of the node.
// 2. Var or ..
func parseNode(node ast.Node) ast.Node {
  fmt.Println("#parseNode")
  switch node.(type) {
  case *ast.Tuple:
    t := node.(*ast.Tuple)
    fmt.Println(t.Nodes[0].Type())
    switch t.Nodes[0].Type() {
    case const_.DEFINE:
      return parseDefine(t)
    case const_.LAMBDA:
      return parseLambda(t)
    case const_.IF:
      return parseIf(t)
    default:
      return parseProc(t)
    }
  case *ast.Ident:
    return parseIdent(&node)
  }

  return node
}

func parseDefine(node *ast.Tuple) ast.Node {
  fmt.Println("#parseDefine()")
  nNode := len(node.Nodes)
  fmt.Println(nNode)
  if nNode != 3 {
    panic("unexpeced define expression")
  }
  //formals := parseNode(node.Nodes[1])
  formals := node.Nodes[1]
  body := parseNode(node.Nodes[2])

  return ast.NewDefine(formals, body)
}

func parseIf(node *ast.Tuple) ast.Node {
  fmt.Println("#parseIf()")
  nNode := len(node.Nodes)
  if nNode != 4 {
    panic("unexpeced if expression")
  }
  test := parseNode(node.Nodes[1])
  conseq := parseNode(node.Nodes[2])
  alt := parseNode(node.Nodes[3])

  return ast.NewIf(test, conseq, alt)
}

func parseLambda(node *ast.Tuple) ast.Node {
  fmt.Println("#parseLambda()")
  nNode := len(node.Nodes)
  if nNode != 3 {
    panic("unexpeced lambda procedure")
  }

  formals := node.Nodes[1]
  body := parseNode(node.Nodes[2])

  return ast.NewLambda(formals, body, nil)
}

func parseIdent(node *ast.Node) ast.Node {
  return nil
  //return (*node).(*ast.Ident)
}

func parseProc(node *ast.Tuple) ast.Node {
  fmt.Println("#parseProc()")
  switch node.Nodes[0].(type) {
  case *ast.Ident:
    name := node.Nodes[0].(*ast.Ident).String()
    args := node.Nodes[1:]
    return ast.NewProc(name, args)

  case *ast.Tuple:
    // ((lambda (x) (+ x x)) 4)
    t := node.Nodes[0].(*ast.Tuple)
    if t.Nodes[0].Type() != const_.LAMBDA {
      panic("unexpeced procedure?")
    }
    args := node.Nodes[1:]
    for _, arg := range args {
      if arg.Type() == const_.TUPLE {
        panic("unexpeced args in lambda proc")
      }
    }
    lambdaNode := parseLambda(t) // ast.Node
    lambda := lambdaNode.(*ast.Lambda)
    return ast.NewLambda(lambda.Formals, lambda.Body, args)

  default:
    panic("unexpeced procedure?")
  }
  return nil
}

func printType(node ast.Node) {
  switch node.(type) {
  case *ast.Tuple:
    t := node.(*ast.Tuple)
    fmt.Printf("( ")
    for _, node := range t.Nodes {
      printType(node)
    }
    fmt.Printf(")")
  default:
    fmt.Printf("%q ", node.Type())
  }
}
