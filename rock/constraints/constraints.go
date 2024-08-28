// Package constraints The stereotype extension type definition is used for function /type annotations
package constraints

type SignedInteger interface {
	~int8 | ~int16 | ~int32 | ~int | ~int64 | ~byte
}

type UnsignedInteger interface {
	~uint8 | ~uint16 | ~uint32 | ~uint | ~uint64 | ~uintptr
}

type Integer interface {
	SignedInteger | UnsignedInteger
}

type Float interface {
	~float32 | ~float64
}

// Equable Simple assignable type, excluding resurrection type, and can judge whether it is equal to.
type Equable interface {
	Integer | Float | ~string | ~bool | ~rune
	// @tip: error is fail
	//Integer | Float | ~string | ~bool | error
}

type KeyIterable interface {
	Integer | Float | ~string
}

// ValueIterable support basic types and composite lists
type ValueIterable interface {
	any
}
