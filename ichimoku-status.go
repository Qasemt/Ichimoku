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
	SencoALineInPast []Point
	//SencoB 26 candle in the past
	SencoBLineInPast []Point

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

func (o *IchimokuStatus) Set_SenCo_A_inPast(buff []Point) {
	o.SencoALineInPast = buff
}
func (o *IchimokuStatus) Set_SenCo_B_inPast(buff []Point) {
	o.SencoBLineInPast = buff
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
func (o *IchimokuStatus) Below(intersection float64) bool {
	if o.SencoA.isNil || o.SencoB.isNil {
		return false
	}
	point_from_price := NewPoint(float64(o.bar.T/1000), o.bar.C)
	res_senko_a, err := o.line_helper.GetCollisionWithLine(point_from_price, o.SencoALineInPast)
	if err != nil {
		return false
	}

	// res_senko_b, err := o.line_helper.GetCollisionWithLine(point_from_price, o.SencoBLineInPast)
	// if err != nil {
	// 	return false
	// }

	return res_senko_a
	//return intersection < o.SencoA.valLine && intersection < o.SencoB.valLine
}
func (o *IchimokuStatus) Above(intersection float64) bool {
	if o.SencoA.isNil || o.SencoB.isNil {
		return false
	}
	point_from_price := NewPoint(float64(o.bar.T/1000), o.bar.C)
	res_senko_a, err := o.line_helper.GetCollisionWithLine(point_from_price, o.SencoALineInPast)
	if err != nil {
		return false
	}

	// res_senko_b, err := o.line_helper.GetCollisionWithLine(point_from_price, o.SencoBLineInPast)
	// if err != nil {
	// 	return false
	// }

	return res_senko_a == false
	//return intersection > o.SencoA.valLine && intersection > o.SencoB.valLine
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
	}

	return result
}
func (o *IchimokuStatus) Print() string {
	d := time.UnixMilli(o.bar.T).Local().Format("2006 Mon Jan 2 15:04:05 ")
	return fmt.Sprintf("ichi %v|%v|%v|%v|%v|G:%v,Chiko UP :%v |status : %v |cloud switching : %v|leaving cloud : %v |%v|%v", o.TenkenSen.Value(), o.KijonSen.Value(), o.SencoA.Value(), o.SencoB.Value(), o.Chiko.Value(), o.Is_cloud_green(), o.IsChikoAbovePrice(), o.GetStatusString(), o.GetCloudSwitching(), o.leavingCloud, d, o.bar.T)

}
