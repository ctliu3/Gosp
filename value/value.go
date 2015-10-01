package value

// The String() method is not for debug.
// Very value (include precedure) should have a value output.
type Value interface {
  String() string
}
