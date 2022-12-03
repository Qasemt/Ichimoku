package ichimoku

import "errors"

var (
	NotEnoughData = errors.New("Not Enough Data")
)

type Equation struct {
	Slope     float64
	Intercept float64
}
type Bar struct {
	Low   float64
	High  float64
	Close float64
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
