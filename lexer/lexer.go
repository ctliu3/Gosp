package lexer

import (
  "unicode"
  "strings"
  "unicode/utf8"
)

type TokenType int

const (
  TokenError TokenType = iota

  TokenEOF
  TokenIdent // letters, >, etc
  TokenLParen // (
  TokenRParen // )
  TokenNumber
  TokenInt
  TokenFloat
)

const EOF = -1

type Token struct {
  Type TokenType
  Name string
}

func (self *Token) String() string {
  return self.Name
  //switch self.Type {
  //case TokenEOF:
    //return "EOF"
  //case TokenError:
    //return self.Name
  //}
  //return "nil"
  //if len(self.Name) > 10 {
    //return fmt.Sprintf("%.10q...", self.Name)
  //}
  //return fmt.Sprintf("%q", self.Name)
}

type Lexer struct {
  name string
  expr string
  start int
  pos int
  width int
  tokenChan chan Token
}

type StateFn func(*Lexer) StateFn

func Lex(name string, expr string) *Lexer {
  l := &Lexer {
    name: name,
    expr: expr,
    tokenChan: make(chan Token),
  }
  go l.run()
  return l
}

func (self *Lexer) NextToken() Token {
  token, ok := <-self.tokenChan
  if ok {
    return token
  } else {
    return Token{TokenEOF, ""}
  }
}

// Do the lexer async
func (self *Lexer) run() {
  for state := lexWhitespace; state != nil; {
    state = state(self)
  }

  close(self.tokenChan)
}

func (self *Lexer) emit(type_ TokenType) {
  self.tokenChan <- Token{type_, self.expr[self.start:self.pos]}
  self.start = self.pos
}

func (self *Lexer) next() rune {
  if self.pos >= len(self.expr) {
    self.width = 0
    return EOF
  }
  cur, size := utf8.DecodeRuneInString(self.expr[self.pos:])
  self.width = size
  self.pos += self.width
  return cur
}

func (self *Lexer) ignore() {
  self.start = self.pos
}

func (self *Lexer) backward() {
  self.pos -= self.width
}

// Return the next rune and not move the self.pos variable.
func (self *Lexer) peek() rune {
  cur := self.next()
  self.backward()
  return cur
}

func lexWhitespace(l *Lexer) StateFn {
  var r rune
  for r = l.next(); isSpace(r); r = l.next() {
  }
  if l.pos >= len(l.expr) {
    return lexEof
  }
  l.backward()
  l.ignore()

  switch r = l.next(); {
  case r == '(':
    return lexLeftParen
  case r == ')':
    return lexRightParen
  case r == '-' || r == '+' || unicode.IsDigit(r):
    next := l.peek()
    if next == EOF {
      panic("expected number or procedure")
    }
    if unicode.IsDigit(r) && (isSpace(next) || unicode.IsDigit(next)) {
      return lexNumber
    }
    fallthrough
  case isAlphaNumeric(r):
    l.backward()
    return lexIdentifier
  case r == EOF:
    return lexEof
  }

  return nil
}

func lexLeftParen(l *Lexer) StateFn {
  l.emit(TokenLParen)
  return lexWhitespace
}

func lexRightParen(l *Lexer) StateFn {
  l.emit(TokenRParen)
  return lexWhitespace
}

func lexAlphaNumber(l *Lexer) StateFn {
  r := l.peek()
  if unicode.IsLetter(r) {
    return lexIdentifier
  } else {
    return lexNumber
  }
}

func lexIdentifier(l *Lexer) StateFn {
  for r := l.next(); r != EOF && isAlphaNumeric(r); r = l.next() {
  }
  l.backward()
  l.emit(TokenIdent)
  return lexWhitespace
}

func lexNumber(l *Lexer) StateFn {
  for r := l.next(); r != EOF && unicode.IsDigit(r); r = l.next() {
  }
  l.backward()
  l.emit(TokenInt)
  return lexWhitespace
}

func lexEof(l *Lexer) StateFn {
  l.emit(TokenEOF)
  return nil
}

func isSpace(r rune) bool {
  return r == ' ' || r == '\t';
}

func isAlphaNumeric(r rune) bool {
  return strings.ContainsRune("+-*/<>=-?", r) || unicode.IsDigit(r) || unicode.IsLetter(r)
}
