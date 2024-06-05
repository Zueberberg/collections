package collections

type signednumbers interface {
	~int8 | ~int16 | ~int32 | ~int64 | ~int
}

type unsignednumbers interface {
	~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uint | ~uintptr
}

type floats interface {
	~float32 | ~float64
}

type specialnumbers interface {
	~byte | ~rune
}

type integers interface {
	signednumbers | unsignednumbers | specialnumbers
}

/*
The ordered interface is a set of next types:

Integers:

	~int8 | ~int16 | ~int32 | ~int64 | ~int

	~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uint | ~uintptr

	~byte | ~rune

Floats:

	~float32 | ~float64

Strings:

	~string
*/
type ordered interface {
	integers | floats | ~string
}
