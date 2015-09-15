package lexer

import (
  "fmt"
  "unicode"
  "strings"
  "unicode/utf8"
)

type TokenType int

const (
  TokenError TokenType = iota

  TokenEOF
  TokenLParen // (
  TokenRParen // )
  TokenLVect // [
  TokenRVect // ]
  TokenIdent // if, define, >, etc
  TokenNumber // only handle int/float, not include complex number
  TokenQuote // '<datum>
  TokenString // "str"
  TokenVect // [1 2 3]
  TokenBool // #t, #f
  TokenChar // #\a, #\space
)

const EOF = -1

type Token struct {
  Type TokenType
  Name string
}

func (self *Token) String() string {
  switch self.Type {
  case TokenEOF:
    return "EOF"
  case TokenError:
    return self.Name
  }
  if len(self.Name) > 10 {
    return fmt.Sprintf("%.10q...", self.Name)
  }
  return fmt.Sprintf("%q", self.Name)
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

func (self *Lexer) emit(typ TokenType) {
  //fmt.Printf("%q, ", self.expr[self.start:self.pos])
  self.tokenChan <- Token{typ, self.expr[self.start:self.pos]}
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

func (self *Lexer) backward() {
  self.pos -= self.width
}

func (self *Lexer) ignore() {
  self.start = self.pos
}

// Return the next rune and not move the self.pos variable.
func (self *Lexer) peek() rune {
  cur := self.next()
  self.backward()
  return cur
}

// Consume the next rune, which is included in a given valid set.
func (self *Lexer) accept (valid string) bool {
  if strings.ContainsRune(valid, self.next()) {
    return true
  }
  self.backward()
  return false
}

// Consume the continuous runes, which are included in a given valid set.
func (self *Lexer) acceptRun(valid string) {
  for strings.ContainsRune(valid, self.next()) {
  }
  self.backward()
}

func (self *Lexer) errorf(format string, args ...interface{}) StateFn {
  self.tokenChan <- Token{TokenError, fmt.Sprintf(format, args...)}
  return nil
}

func lexWhitespace(l *Lexer) StateFn {
  var r rune
  for r = l.next(); isSpace(r); r = l.next() {
  }
  l.backward()
  l.ignore()

  switch r = l.next(); {
  case r == '(':
    return lexLeftParen
  case r == ')':
    return lexRightParen
  case r == ']':
    return lexLeftVect
  case r == '[':
    return lexRightVect
  case r == '"':
    return lexString
  case r == '\'':
    return lexQuote
  case r == '#':
    return lexSharp
  case r == '-' || r == '+' || unicode.IsDigit(r): // number or ident
    next := l.peek()
    if next == EOF {
      panic("expected number or procedure")
    }
    if !unicode.IsDigit(r) && isSpace(next) {
      return lexIdentifier
    }
    if scanNumber(l) {
      return lexNumber
    }

  case isAlphaNumeric(r):
    l.backward()
    return lexIdentifier
  case r == EOF:
    return lexEOF
  }

  return nil
}

func lexError(l *Lexer) StateFn {
  l.emit(TokenError)
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

func lexLeftVect(l *Lexer) StateFn {
  l.emit(TokenLVect)
  return lexWhitespace
}

func lexRightVect(l *Lexer) StateFn {
  l.emit(TokenRVect)
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
  for r := l.next(); isAlphaNumeric(r); r = l.next() {
  }
  l.backward()
  l.emit(TokenIdent)
  return lexWhitespace
}

func lexNumber(l *Lexer) StateFn {
  l.emit(TokenNumber)
  return lexWhitespace
}

func lexChar(l *Lexer) StateFn {
  l.emit(TokenChar)
  return lexWhitespace
}

// "hello", "he\"llo"
func lexString(l *Lexer) StateFn {
Loop:
  for {
    switch l.next() {
    case '\\':
      if r := l.next(); r == '\n' || r == EOF {
        return l.errorf("unexpeced eof of string?")
      }
    case EOF:
      return l.errorf("unexpeced eof of string?")
    case '"':
      break Loop
    }
  }
  if isAlphaNumeric(l.peek()) {
    return l.errorf("unexpeced string?")
  }
  l.emit(TokenString)
  return lexWhitespace
}

func lexQuote(l *Lexer) StateFn {
  for r := l.next(); isAlphaNumeric(r); r = l.next() {
  }
  l.backward()
  l.emit(TokenQuote)
  return lexWhitespace
}

func lexBool(l *Lexer) StateFn {
  l.emit(TokenBool)
  return lexWhitespace
}

func lexSharp(l *Lexer) StateFn {
  r := l.next()
  var digits string

  switch r {
  case 't', 'f':
    if !isAlphaNumeric(l.peek()) {
      return lexBool
    }
  case '\\':
    if isChar(l.peek()) {
      for unicode.IsLetter(l.next()) {
      }
      l.backward()
      if !isAlphaNumeric(l.peek()) {
        return lexChar
      }
    }
  case 'b':
    digits = "01"
    fallthrough
  case 'o':
    digits = "01234567"
    fallthrough
  case 'd':
    digits = "0123456789"
    fallthrough
  case 'x':
    digits = "0123456789ABCFEF"
    if strings.ContainsRune(digits, l.peek()) {
      l.acceptRun(digits)
      if !isAlphaNumeric(l.peek()) {
        return lexNumber
      }
    }
  }

  return l.errorf("unexpeced charactor after #?")
}

func lexEOF(l *Lexer) StateFn {
  l.emit(TokenEOF)
  return nil
}

// 3e3.4 is illegal
func scanNumber(l *Lexer) bool {
  l.accept("+-")
  digits := "0123456789"
  l.acceptRun(digits)
  if l.accept(".") {
    l.acceptRun(digits)
  }
  if l.accept("eE") {
    l.accept("+-")
    l.acceptRun(digits)
  }
  if isAlphaNumeric(l.peek()) {
    l.errorf("unexpeced number?")
    return false
  }

  return true
}

func isSpace(r rune) bool {
  return r == ' ' || r == '\t' || r == '\n';
}

func isAlphaNumeric(r rune) bool {
  return strings.ContainsRune(`!$%&*+-./:<=>?@^_~`, r) || unicode.IsDigit(r) || unicode.IsLetter(r)
}

func isChar(r rune) bool {
  return strings.ContainsRune(` ()[]`, r) || unicode.IsLetter(r)
}
