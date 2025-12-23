package typing

// +Foo=true
// +Bar=123
type StringType string

// +Foo=true
// +Bar=123
type IntType int

// +Foo=true
// +Bar=123
type AStruct struct{}

// +Foo=true
// +Bar=123
type AStructType AStruct

// +Foo=true
// +Bar=123
type SliceType []AStruct

// +Foo=true
// +Bar=123
type AliasStringType = string
