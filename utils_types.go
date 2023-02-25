package glog

type Ints interface {
	int64 | int32 | int16 | int8 | int
}

type IntOrInterface interface {
	int64 | int32 | int16 | int8 | int | interface{}
}

type Uints interface {
	uint64 | uint32 | uint16 | uint8 | uint
}

type UintOrInterface interface {
	uint64 | uint32 | uint16 | uint8 | uint | interface{}
}

type IntOrUint interface {
	int64 | int32 | int16 | int8 | int | uint64 | uint32 | uint16 | uint8 | uint
}

type Floats interface {
	float32 | float64
}

type FloatOrInterface interface {
	float32 | float64 | interface{}
}

type Number interface {
	int64 | int32 | int16 | int8 | int | uint64 | uint32 | uint16 | uint8 | uint | float32 | float64
}

type NumberOrInterface interface {
	int64 | int32 | int16 | int8 | int | uint64 | uint32 | uint16 | uint8 | uint | float32 | float64 | interface{}
}

type Durations interface {
	uint32 | uint64 | uint | int32 | int64 | int | float32 | float64
}
