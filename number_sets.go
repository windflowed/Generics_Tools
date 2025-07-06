package Generics_Tools

type RealNumber interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64
}

type Number interface {
	~complex64 | ~complex128 | RealNumber
}
