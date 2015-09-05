package value

// The String() method is not for debug.
// Proc -> Proc name
// Var -> Var
// Ident -> Ident name
type Value interface {
  String() string
}
