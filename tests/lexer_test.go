package tests

import (
  //"fmt"
  "testing"
  "strings"
  "github.com/ctliu3/gosp/lexer"
)

func lex(l *lexer.Lexer) string {
  var res []string
  for token := l.NextToken(); token.Type != lexer.TokenEOF; token = l.NextToken() {
    res = append(res, token.Name)
  }
  return strings.Join(res, " ")
}

func TestLexer(t *testing.T) {
  mp := map[string]string {
    "(define a 10)": "( define a 10 )",
    "(define a \"hello\")": "( define a \"hello\" )",
    `(define a "he\"llo")`: `( define a "he\"llo" )`,
    "(define a 3.4)": "( define a 3.4 )",
    "(+ 3 -4000)": "( + 3 -4000 )",
    "(- 30e10 -4000)": "( - 30e10 -4000 )",
    "(- 3.e10 -4E-2)": "( - 3.e10 -4E-2 )",
    "(if (> a 3) (display a) (newline))": "( if ( > a 3 ) ( display a ) ( newline ) )",
  }

  for raw, slex := range mp {
    res := lex(lexer.Lex("test", raw))
    if res != slex {
      t.Errorf("fail in case: %q, err: %q", raw, res)
    }
  }
}
