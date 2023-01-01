package glog

type Ints interface {
	int64 | int32 | int16 | int8 | int
}

type IntOrInterface interface {
	Ints | interface{}
}

type Uints interface {
	uint64 | uint32 | uint16 | uint8 | uint
}

type UintOrInterface interface {
	Uints | interface{}
}

type IntOrUint interface {
	Ints | Uints
}

type Floats interface {
	float32 | float64
}

type FloatOrInterface interface {
	Floats | interface{}
}

type Number interface {
	IntOrUint | Floats
}

type NumberOrInterface interface {
	Number | interface{}
}

type Durations interface {
	uint32 | uint64 | uint | int32 | int64 | int | Floats
}
