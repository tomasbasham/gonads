package gonads

// Func is a function that takes a value and returns a value.
type Func[T, U any] func(T) U

// Unit is a type that has only one value.
type Unit struct{}

// UnitValue is the only value of type Unit.
var UnitValue = struct{}{}
