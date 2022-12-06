package ichimoku

import "errors"

var (
	NotEnoughData = errors.New("Not Enough Data")
	DataNotFill   = errors.New("Data not fill")
)

type Point struct {
	X float64
	Y float64
}

func NewPoint(x float64, y float64) Point {
	p := Point{}
	p.X = x
	p.Y = y
	return p
}

type Equation struct {
	Slope     float64
	Intercept float64
}
type Bar struct {
	L float64
	H float64
	C float64
	O float64
	V float64
	T int64
}
type ELine int

const (
	Lin_Tenkan_sen   ELine = 9
	Line_kijon_sen   ELine = 26
	Line_spanPeriod  ELine = 52
	Line_chikoPeriod ELine = 26 //-26
)

type EIchimokuStatus int

const (
	IchimokuStatus_NAN          EIchimokuStatus = 0
	IchimokuStatus_Cross_Inside EIchimokuStatus = 1
	IchimokuStatus_Cross_Below  EIchimokuStatus = 2
	IchimokuStatus_Cross_Above  EIchimokuStatus = 3
	IchimokuStatus_overLab      EIchimokuStatus = 4
)
