package ichimoku

import "fmt"

type Ichimoku struct {
	TenkenSen ValueLine
	KijonSen  ValueLine
	SencoA    ValueLine
	SencoB    ValueLine
	Chiko     ValueLine
	bar       Bar
	//-----
	Status       EIchimokuStatus
	Folding      bool
	LeavingCloud bool
}

func NewIchimoku(tenken ValueLine, kijon ValueLine, sencoA ValueLine, sencoB ValueLine, chiko ValueLine, bar Bar) *Ichimoku {

	o := Ichimoku{}
	o.TenkenSen = tenken
	o.KijonSen = kijon
	o.SencoA = sencoA
	o.SencoB = sencoB
	o.Chiko = chiko
	o.bar = bar
	o.Status = IchimokuStatus_NAN
	return &o
}
func (o *Ichimoku) SetStatus(status EIchimokuStatus) {
	o.Status = status
}
func (o *Ichimoku) GetStatus() EIchimokuStatus {
	return o.Status
}

func (o *Ichimoku) SetLeavingCloud(v bool) {
	o.LeavingCloud = v
}
func (o *Ichimoku) GetLeavingCloud() bool {
	return o.LeavingCloud
}
func (o *Ichimoku) SetFolding(v bool) {
	o.Folding = v
}
func (o *Ichimoku) GetFolding() bool {
	return o.Folding
}
func (o *Ichimoku) Is_cloud_green() bool {
	return o.SencoA.valLine > o.SencoB.valLine
}
func (o *Ichimoku) IsChikoAbovePrice() bool {
	return o.bar.Close > o.Chiko.valLine
}
func (o *Ichimoku) Below(intersection float64) bool {
	if o.SencoA.isNil || o.SencoB.isNil {
		return false
	}
	return intersection < o.SencoA.valLine && intersection < o.SencoB.valLine
}
func (o *Ichimoku) Above(intersection float64) bool {
	if o.SencoA.isNil || o.SencoB.isNil {
		return false
	}
	return intersection > o.SencoA.valLine && intersection > o.SencoB.valLine
}

func (o *Ichimoku) inside(intersection float64) bool {
	return o.in_float_range(intersection, o.SencoA.valLine, o.SencoB.valLine)
}

func (o *Ichimoku) in_float_range(num float64, range_a float64, range_b float64) bool {
	if range_a > range_b {
		return num >= range_b && num <= range_a
	} else {
		return num <= range_b && num >= range_a
	}
}
func (o *Ichimoku) GetStatusString() string {
	result := ""
	switch o.Status {
	case IchimokuStatus_NAN:
		result = "nan"
	case IchimokuStatus_Cross_Inside:
		result = "cross inside"
	case IchimokuStatus_Cross_Below:
		result = "cross below"
	case IchimokuStatus_Cross_Above:
		result = "cross above"
	}

	return result
}
func (o *Ichimoku) Print() string {
	return fmt.Sprintf("ichi %v|%v|%v|%v|%v|G:%v,Chiko UP :%v |status : %v |Folding : %v|leaving cloud : %v ", o.TenkenSen.Value(), o.KijonSen.Value(), o.SencoA.Value(), o.SencoB.Value(), o.Chiko.Value(), o.Is_cloud_green(), o.IsChikoAbovePrice(), o.GetStatusString(), o.Folding, o.LeavingCloud)
}
