package ichimoku

import (
	"fmt"
	"time"
)

type IchimokuStatus struct {
	TenkenSen ValueLine

	//_______________

	KijonSen ValueLine

	//in the future
	SenKoA_Shifted26 ValueLine

	//in the future
	SenKoB_Shifted26 ValueLine

	//extract value sen A & B from 26 candle past (26 shift forward in calc ichimoku)
	//SencoA 26 candle in the past (for check)
	SencoA_Past Point
	//extract value sen A & B from 26 candle past (26 shift forward in calc ichimoku)
	//SencoB 26 candle in the past (for check)
	SencoB_Past Point

	ChikoSpan Bar //close bar

	bar Bar
	//-----
	Status         EIchimokuStatus
	cloudSwitching bool

	line_helper lineHelper
}

func NewIchimokuStatus(tenken ValueLine, kijon ValueLine, senKoA_Shifted26 ValueLine, senKoB_Shifted52 ValueLine, chiko_span Bar, bar Bar) *IchimokuStatus {

	o := IchimokuStatus{}

	o.TenkenSen = tenken

	o.KijonSen = kijon

	o.SenKoA_Shifted26 = senKoA_Shifted26
	o.SenKoB_Shifted26 = senKoB_Shifted52

	o.ChikoSpan = chiko_span
	o.bar = bar
	o.Status = IchimokuStatus_NAN
	o.line_helper = NewLineHelper()

	return &o
}

//----------------------------------------GET SET

func (o *IchimokuStatus) SetChikoSpan(v Bar) {
	o.ChikoSpan = v
}
func (o *IchimokuStatus) Set_SenCo_A_Past(p Point) {
	o.SencoA_Past = p
}
func (o *IchimokuStatus) Set_SenCo_B_Past(p Point) {
	o.SencoB_Past = p
}
func (o *IchimokuStatus) SetStatus(status EIchimokuStatus) {
	o.Status = status
}
func (o *IchimokuStatus) GetStatus() EIchimokuStatus {
	return o.Status
}

func (o *IchimokuStatus) SetCloudSwitching(v bool) {
	o.cloudSwitching = v
}
func (o *IchimokuStatus) GetCloudSwitching() bool {
	return o.cloudSwitching
}
func (o *IchimokuStatus) Is_cloud_green() bool {
	return o.SenKoA_Shifted26.valLine > o.SenKoB_Shifted26.valLine
}
func (o *IchimokuStatus) IsChikoAbovePrice() bool {
	return o.ChikoSpan.H > o.bar.C
}
func (o *IchimokuStatus) CloudStatus(intersection float64) EIchimokuStatus {
	if o.SenKoA_Shifted26.isNil || o.SenKoB_Shifted26.isNil {
		return IchimokuStatus_NAN
	}
	if o.SencoA_Past.isNil || o.SencoB_Past.isNil {
		return IchimokuStatus_NAN
	}

	sen_B := o.SencoB_Past //Senko B in_26_candle_pass
	sen_A := o.SencoA_Past //Senko A in_26_candle_pass
	if sen_A.Y > intersection && sen_B.Y > intersection {
		return IchimokuStatus_Cross_Below
	} else if sen_A.Y < intersection && sen_B.Y < intersection {
		return IchimokuStatus_Cross_Above
	} else if sen_A.Y < intersection && sen_B.Y > intersection || sen_A.Y > intersection && sen_B.Y < intersection {
		return IchimokuStatus_Cross_Inside
	}

	return IchimokuStatus_NAN

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
	return fmt.Sprintf("ichi %v|%v|%v|%v|%v|G:%v,Chiko UP :%v |status : %v |%v|%v", o.TenkenSen.Value(), o.KijonSen.Value(), o.SenKoA_Shifted26.Value(), o.SenKoB_Shifted26.Value(), o.ChikoSpan.C, o.Is_cloud_green(), o.IsChikoAbovePrice(), o.GetStatusString(), d, o.bar.T)

}
