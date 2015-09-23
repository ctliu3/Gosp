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
  parser.parse(&tuples, 0, -1)

  // For debugging.
  for _, node := range tuples {
    printType(node)
    fmt.Println()
  }

  return parseNodes(tuples)
}

func (self *Parser) parse(nodes *[]ast.Node, dep int, typ lexer.TokenType) {
  for token := self.l.NextToken(); token.Type != lexer.TokenEOF; token = self.l.NextToken() {

    var node ast.Node
    switch token.Type {
    case lexer.TokenBool:
      node = ast.NewBool(token.Name)

    case lexer.TokenChar:
      node = ast.NewChar(token.Name)

    case lexer.TokenString:
      node = ast.NewString(token.Name)

    case lexer.TokenQuote:
      quoteExpr := []ast.Node{ast.NewIdent("quote")}
      self.parse(&quoteExpr, dep + 1, lexer.TokenQuote)
      node = ast.NewTuple(quoteExpr)

    case lexer.TokenQuasiQuote:
      qqExpr := []ast.Node{ast.NewIdent("quasiquote")}
      self.parse(&qqExpr, dep + 1, lexer.TokenQuasiQuote)
      node = ast.NewTuple(qqExpr)

    case lexer.TokenComma:

    case lexer.TokenIdent:
      node = ast.NewIdent(token.Name)

    case lexer.TokenNumber:
      if strings.ContainsAny(token.Name, "Ee.") {
        node = ast.NewFloat(token.Name)
      } else {
        node = ast.NewInt(token.Name)
      }

    case lexer.TokenLParen:
      sub := make([]ast.Node, 0)
      self.parse(&sub, dep + 1, -1)
      node = ast.NewTuple(sub)

    case lexer.TokenRParen:
      return

    default:
      panic("unexpeced token")
    }
    *nodes = append(*nodes, node)
    if typ == lexer.TokenQuote || typ == lexer.TokenQuasiQuote {
      break
    }
  }
}

func parseNodes(tuples []ast.Node) []ast.Node {
  var nodes []ast.Node
  for _, tuple := range tuples {
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

    if len(t.Nodes) == 0 {
      return parseList(t)
    }

    fmt.Println("#parseNode#" + t.Nodes[0].Type())
    switch t.Nodes[0].Type() {
    case const_.DEFINE:
      return parseDefine(t)
    case const_.SET:
      return parseSet(t)
    case const_.DISPLAY:
      return parseDisplay(t)
    case const_.BEGIN:
      return parseBegin(t)
    case const_.COND:
      return parseCond(t)
    case const_.QUOTE:
      return parseQuote(t)
    case const_.LAMBDA:
      return parseLambda(t)
    case const_.IF:
      return parseIf(t)
    case const_.LET, const_.LET_STAR, const_.LETREC:
      return parseLets(t)
    case const_.IMPORT:
      return parseImport(t)
    case const_.LIST:
      return parseList(t)
    default:
      return parseCall(t)
    }
  case *ast.Ident:
    return parseIdent(&node)
  }

  return node
}

func parseDefine(node *ast.Tuple) ast.Node {
  fmt.Println("#parseDefine")
  nNode := len(node.Nodes)
  if nNode != 3 {
    panic("unexpeced define expression")
  }
  // implicit lambda expressions, e.g., (define (add x y) (+ x y))
  var vars ast.Node
  var expr ast.Node

  if node.Nodes[1].Type() == const_.TUPLE {
    t := node.Nodes[1].(*ast.Tuple)
    vars = t.Nodes[0]

    var args []ast.Node
    for _, node := range t.Nodes[1:] {
      args = append(args, node)
    }

    body := parseNode(node.Nodes[2])
    expr = ast.NewLambda(ast.NewTuple(args), body)
  } else {

    // 1. normal definition, e.g. (defien a 1)
    // 2. explicit lambda expressions, e.g. (define inc (lambda (x) (+ x 1)))
    vars = node.Nodes[1]
    expr = parseNode(node.Nodes[2])
  }

  return ast.NewDefine(vars, expr)
}

func parseSet(node *ast.Tuple) ast.Node {
  fmt.Println("#parseSet")
  nNode := len(node.Nodes)
  if nNode != 3 {
    panic("unexpeced define expression")
  }
  formals := node.Nodes[1]
  body := parseNode(node.Nodes[2])

  return ast.NewSet(formals, body)
}

func parseList(node *ast.Tuple) ast.Node {
  fmt.Println("#parseList")
  nNode := len(node.Nodes)
  if nNode == 0 {
    return ast.NewList(make([]ast.Node, 0))
  }
  var nodes []ast.Node
  for _, node := range node.Nodes[1:] {
    nodes = append(nodes, parseNode(node))
    //if i > 0 && nodes[i].Type() != nodes[0].Type() {
      //panic("list: type not match")
    //}
  }
  return ast.NewList(nodes)
}

func parseDisplay(node *ast.Tuple) ast.Node {
  fmt.Println("#parseDisplay")
  nNode := len(node.Nodes)
  if nNode != 2 {
    panic("unexpeced display expression")
  }
  obj := parseNode(node.Nodes[1])

  return ast.NewDisplay(obj)
}

func parseBegin(node *ast.Tuple) ast.Node {
  fmt.Println("#parseBegin")
  nNode := len(node.Nodes)
  if nNode < 2 {
    panic("unexpeced begin expression")
  }
  exprs := make([]ast.Node, len(node.Nodes) - 1)
  for i, _ := range exprs {
    exprs[i] = parseNode(node.Nodes[i + 1])
  }

  return ast.NewBegin(exprs)
}

func parseCond(node *ast.Tuple) ast.Node {
  var clause []ast.Tuple
  for _, node := range node.Nodes[1:] {
    fmt.Printf("len = %v\n", len(node.(*ast.Tuple).Nodes))
    if node.Type() != const_.TUPLE || len(node.(*ast.Tuple).Nodes) != 2 {
      panic("cond: bad syntax")
    }
    var t ast.Tuple
    for _, item := range node.(*ast.Tuple).Nodes {
      t.Nodes = append(t.Nodes, parseNode(item))
    }
    clause = append(clause, t)
  }
  return ast.NewCond(clause)
}

func parseQuote(node *ast.Tuple) ast.Node {
  fmt.Printf("#parseQuote, %v\n", len(node.Nodes))
  nNode := len(node.Nodes)
  if nNode < 2 {
    panic("unexpeced quote expression")
  }
  datum := parseNode(node.Nodes[1])

  return ast.NewQuote(datum)
}

func parseIf(node *ast.Tuple) ast.Node {
  fmt.Println("#parseIf")
  nNode := len(node.Nodes)
  if nNode != 4 {
    panic("unexpeced if expression")
  }
  test := parseNode(node.Nodes[1])
  conseq := parseNode(node.Nodes[2])
  alt := parseNode(node.Nodes[3])

  return ast.NewIf(test, conseq, alt)
}

func parseLets(node *ast.Tuple) ast.Node {
  fmt.Println("#parseLet")
  nNode := len(node.Nodes)
  if nNode < 3 {
    panic("let: bad syntax")
  }
  bindings := parseBinds(node.Nodes[1])
  body := parseNodes(node.Nodes[2:])

  switch node.Nodes[0].Type() {
  case const_.LET:
    return ast.NewLet(bindings, body)
  case const_.LET_STAR:
    return ast.NewLetStar(bindings, body)
  case const_.LETREC:
    return ast.NewLetrec(bindings, body)
  default:
    panic("unexpeced let expression?")
  }
  return nil
}

func parseBinds(node ast.Node) ast.Binds {
  if node.Type() != const_.TUPLE {
    panic("let: bad syntax")
  }

  nodes := node.(*ast.Tuple).Nodes
  binds := ast.NewBinds(nil)
  for _, bind := range nodes {
    if bind.Type() != const_.TUPLE {
      panic("bad syntax")
    }
    t := bind.(*ast.Tuple)
    if len(t.Nodes) != 2 && t.Nodes[0].Type() != const_.IDENT {
      panic("bad syntax")
    }
    binds.Bindings = append(
      binds.Bindings, *ast.NewBind(t.Nodes[0].(*ast.Ident).Name, parseNode(t.Nodes[1])))
  }
  return *binds
}

func parseImport(node *ast.Tuple) ast.Node {
  // (import "lib.gosp")

  if len(node.Nodes) != 2 {
    panic("import: bad syntax")
  }
  filename := node.Nodes[1].(*ast.Ident).Name
  return ast.NewImport(filename)
}

func parseLambda(node *ast.Tuple) ast.Node {
  fmt.Println("#parseLambda")
  nNode := len(node.Nodes)
  if nNode != 3 {
    panic("unexpeced lambda procedure")
  }

  formals := node.Nodes[1]
  body := parseNode(node.Nodes[2])

  return ast.NewLambda(formals, body)
}

func parseIdent(node *ast.Node) ast.Node {
  return *node;
}

func parseCall(node *ast.Tuple) ast.Node {
  fmt.Println("#parseCall")
  switch node.Nodes[0].(type) {
  case *ast.Ident:
    // (f x y)
    operator := node.Nodes[0].(*ast.Ident)
    args := make([]ast.Node, len(node.Nodes) - 1)
    for i, _ := range args {
      args[i] = parseNode(node.Nodes[i + 1])
    }
    return ast.NewCall(operator, args)

  case *ast.Tuple:
    // ((lambda (x) (+ x x)) 4)
    // ((if #f + *)) 3 4)
    callee := node.Nodes[0].(*ast.Tuple)
    args := node.Nodes[1:]
    for _, arg := range args {
      if arg.Type() == const_.TUPLE {
        panic("unexpeced args in lambda proc")
      }
    }
    return ast.NewCall(parseNode(callee), args)

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
    fmt.Printf(") ")
  default:
    fmt.Printf("%v ", node.Type())
  }
}
