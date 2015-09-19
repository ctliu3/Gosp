package value

// Use interface to avoid import cycles.
type Closure struct {
  Env interface{} // *scope.Scope
  Lambda interface{} // *ast.Lambda
}

func NewClosure(env interface{}, lambda interface{}) *Closure {
  return &Closure{env, lambda}
}

func (self *Closure) String() string {
  // TODO
  return "closure"
}
