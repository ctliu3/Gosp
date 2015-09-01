package parser

import (
  "fmt"

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
  return parseList(&tuples)
}

func (self *Parser) parse(tuples *[]ast.Node, dep int) {
  for token := self.l.NextToken(); token.Type != lexer.TokenEOF; token = self.l.NextToken() {
    fmt.Println(token.Name)

    switch token.Type {
    case lexer.TokenIdent:
      *tuples = append(*tuples, ast.NewIdent(token.Name))
    case lexer.TokenInt:
      *tuples = append(*tuples, ast.NewInt(token.Name))
    case lexer.TokenFloat:
      *tuples = append(*tuples, ast.NewFloat(token.Name))
    case lexer.TokenLParen:
      sub := make([]ast.Node, 0)
      self.parse(&sub, dep + 1)
      *tuples = append(*tuples, ast.NewTuple(sub))
    case lexer.TokenRParen:
      return
    //default:
      //panic("unexpeced token")
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

func parseNode(tuple ast.Node) ast.Node {
  t := tuple.(*ast.Tuple)
  switch t.Nodes[0].Type() {
  case const_.DEFINE:
    return parseDefine(t)
  case const_.IF:
    return parseIf(t)
  }
  return nil
}

func parseDefine(node *ast.Tuple) ast.Node {
  fmt.Println("define")
  return nil
  //return &NewIf{}
}

func parseIf(node *ast.Tuple) ast.Node {
  fmt.Println("if")
  nexpr := len(node.Nodes)
  if nexpr != 4 {
    panic("unexpeced if expression")
  }
  test := parseNode(node.Nodes[1])
  conseq := parseNode(node.Nodes[2])
  alt := parseNode(node.Nodes[3])

  return ast.NewIf(test, conseq, alt)
}
