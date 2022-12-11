package ichimoku

import (
	"fmt"
	"time"
)

type IchimokuStatus struct {
	TenkenSen ValueLine
	KijonSen  ValueLine
	//in the future
	SencoA ValueLine
	//in the future
	SencoB ValueLine

	//SencoA 26 candle in the past
	SencoALineInPast Point
	//SencoB 26 candle in the past
	SencoBLineInPast Point

	Chiko ValueLine
	bar   Bar
	//-----
	Status         EIchimokuStatus
	cloudSwitching bool
	leavingCloud   bool
	line_helper    lineHelper
}

func NewIchimokuStatus(tenken ValueLine, kijon ValueLine, sencoA ValueLine, sencoB ValueLine, chiko ValueLine, bar Bar) *IchimokuStatus {

	o := IchimokuStatus{}
	o.TenkenSen = tenken
	o.KijonSen = kijon
	o.SencoA = sencoA
	o.SencoB = sencoB
	o.Chiko = chiko
	o.bar = bar
	o.Status = IchimokuStatus_NAN
	o.line_helper = NewLineHelper()
	// if !o.KijonSen.isNil && o.bar.C > o.KijonSen.valLine {
	// 	o.CrossKijonAndPrice = true
	// }

	return &o
}

func (o *IchimokuStatus) Set_SenCo_A_inPast(p Point) {
	o.SencoALineInPast = p
}
func (o *IchimokuStatus) Set_SenCo_B_inPast(p Point) {
	o.SencoBLineInPast = p
}
func (o *IchimokuStatus) SetStatus(status EIchimokuStatus) {
	o.Status = status
}
func (o *IchimokuStatus) GetStatus() EIchimokuStatus {
	return o.Status
}

func (o *IchimokuStatus) SetLeavingCloud(v bool) {
	o.leavingCloud = v
}
func (o *IchimokuStatus) GetLeavingCloud() bool {
	return o.leavingCloud
}
func (o *IchimokuStatus) SetCloudSwitching(v bool) {
	o.cloudSwitching = v
}
func (o *IchimokuStatus) GetCloudSwitching() bool {
	return o.cloudSwitching
}
func (o *IchimokuStatus) Is_cloud_green() bool {
	return o.SencoA.valLine > o.SencoB.valLine
}
func (o *IchimokuStatus) IsChikoAbovePrice() bool {
	return o.bar.C > o.Chiko.valLine
}
func (o *IchimokuStatus) CloudStatus(intersection float64) EIchimokuStatus {
	if o.SencoA.isNil || o.SencoB.isNil {
		return IchimokuStatus_NAN
	}
	if o.SencoALineInPast.isNil || o.SencoBLineInPast.isNil {
		return IchimokuStatus_NAN
	}
	// point_from_price := NewPoint(float64(o.bar.T/1000), o.bar.C)
	sen_B := o.SencoBLineInPast //Senko B in_26_candle_pass
	sen_A := o.SencoALineInPast //Senko A in_26_candle_pass
	if sen_A.Y > intersection && sen_B.Y > intersection {
		return IchimokuStatus_Cross_Below
	} else if sen_A.Y < intersection && sen_B.Y < intersection {
		return IchimokuStatus_Cross_Above
	} else if sen_A.Y < intersection && sen_B.Y > intersection || sen_A.Y > intersection && sen_B.Y < intersection {
		return IchimokuStatus_Cross_Inside
	}

	return IchimokuStatus_NAN
	// res_senko_a, err := o.line_helper.GetCollisionWithLine(point_from_price, o.SencoALineInPast)
	// if err != nil {
	// 	return false
	// }

	// res_senko_b, err := o.line_helper.GetCollisionWithLine(point_from_price, o.SencoBLineInPast)
	// if err != nil {
	// 	return false
	// }

	//return res_senko_a == EPointLocation_below
	//return intersection < o.SencoA.valLine && intersection < o.SencoB.valLine
}

func (o *IchimokuStatus) GetStatusString() string {
	result := ""
	switch o.Status {
	case IchimokuStatus_NAN:
		result = "nan"

	case IchimokuStatus_Cross_Below:
		result = "cross below"
	case IchimokuStatus_Cross_Above:
		result = "cross above"
	case IchimokuStatus_Cross_Inside:
		result = "cross inside"
	}

	return result
}
func (o *IchimokuStatus) Print() string {
	d := time.UnixMilli(o.bar.T).Local().Format("2006 Mon Jan 2 15:04:05 ")
	return fmt.Sprintf("ichi %v|%v|%v|%v|%v|G:%v,Chiko UP :%v |status : %v |cloud switching : %v|leaving cloud : %v |%v|%v", o.TenkenSen.Value(), o.KijonSen.Value(), o.SencoA.Value(), o.SencoB.Value(), o.Chiko.Value(), o.Is_cloud_green(), o.IsChikoAbovePrice(), o.GetStatusString(), o.GetCloudSwitching(), o.leavingCloud, d, o.bar.T)

}
