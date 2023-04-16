package climacell

import "time"

type FloatValue struct {
	Value *float64
	Units string
}

type NonNullableTimeValue struct{ Value time.Time }

type Weather struct {
	Lat             float64
	Lon             float64
	Temp            *FloatValue
	ObservationTime NonNullableTimeValue
}
