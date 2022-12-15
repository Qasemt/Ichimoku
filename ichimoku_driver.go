package ichimoku

import (
	"fmt"
	"time"
)

type IIchimokuDriver interface {

	//
	// out result array sorte : [new day -to- old day]
	//once call per asset
	MakeIchimokuInPast(bars *[]Bar, numberOfRead int) error

	//once call per asset
	PreAnalyseIchimoku(data []IchimokuStatus) (*IchimokuStatus, error)

	//Analys Trigger Cross tenken & kijon sen
	AnalyseTriggerCross(previous IchimokuStatus, bars_only_25_bars_latest []Bar) (*IchimokuStatus, error)

	GetLastDay() *IchimokuStatus
	GetListDay() []IchimokuStatus
}

type IchimokuDriver struct {
	line_helper    lineHelper
	bars           []Bar
	QIchimokuDays  []IchimokuStatus
	ConversionLine []float64
	BaseLine       []float64
	LeadingSpanA   []float64
	LeadingSpanB   []float64
	laggingSpan    []float64
}

func NewIchimokuDriver() IIchimokuDriver {
	xx := IchimokuDriver{}
	xx.QIchimokuDays = make([]IchimokuStatus, 0)
	xx.line_helper = NewLineHelper()
	return &xx
}

func (xx *IchimokuDriver) loadbars(from int, to int) []Bar {
	if len(xx.bars) == 0 {
		return nil
	}
	if from < 0 {
		from = 0
	}
	if to > len(xx.bars) {
		to = len(xx.bars)
	}

	return xx.bars[from:to]

}

func (o *IchimokuDriver) GetListDay() []IchimokuStatus {
	return o.QIchimokuDays
}

//
// out result array sorte : [new day -to- old day]
func (o *IchimokuDriver) MakeIchimokuInPast(bars *[]Bar, numberOfRead int) error {

	if len(*bars) == 0 {
		return DataNotFill
	}

	if len(*bars) < numberOfRead || len(*bars) < 52 {
		return NotEnoughData
	}

	fmt.Printf("Calc ichi from Last %v days\r\n", numberOfRead)
	o.bars = *bars

	//days:sorted descending
	bars_len := len(o.bars)
	for day := 0; day < numberOfRead; day++ {

		from := bars_len - 52 - day
		to := bars_len - day

		ic, err := BuildIchimokuStatus(o.loadbars(from, to))
		if err != nil {
			return err
		}
		o.Put(ic)

	}

	for day_index := 0; day_index < o.NumberOfIchimoku(); day_index++ {
		item := &(o.QIchimokuDays)[day_index]
		sen_a, sen_b := o.Calc_Cloud_InPast(day_index)
		item.Set_SenCo_A_Past(sen_a)
		item.Set_SenCo_B_Past(sen_b)
	}

	return nil

}

//analyse with two days
func (o *IchimokuDriver) FindStatusin26BarPast() {

	for day_index := 0; day_index < o.NumberOfIchimoku(); day_index++ {
		item := o.QIchimokuDays[day_index]
		sen_a, sen_b := o.Calc_Cloud_InPast(day_index)
		item.Set_SenCo_A_Past(sen_a)
		item.Set_SenCo_B_Past(sen_b)
	}

}

//Analys Trigger Cross tenken & kijon sen
//bars : contain with new bar
func (o *IchimokuDriver) AnalyseTriggerCross(previous IchimokuStatus, _52_bars_latest []Bar) (*IchimokuStatus, error) {

	if len(_52_bars_latest) == 0 {
		return nil, DataNotFill
	}

	if len(_52_bars_latest) < 52 {
		return nil, NotEnoughData
	}

	newIchi, e := BuildIchimokuStatus(_52_bars_latest)

	if e != nil {
		return nil, e
	}
	sen_a_in_26_past, sen_b_in_26_past := o.Calc_Cloud_InPast(0)

	if sen_a_in_26_past.isNil || sen_b_in_26_past.isNil {
		return nil, ChikoStatus26InPastNotMade
	}

	newIchi.Set_SenCo_A_Past(sen_a_in_26_past)
	newIchi.Set_SenCo_B_Past(sen_b_in_26_past)

	if newIchi.SencoA_Past.isNil || newIchi.SencoB_Past.isNil {
		return nil, ChikoStatus26InPastNotMade
	}

	time_current := time.UnixMilli(newIchi.bar.T)
	time_previous := time.UnixMilli(previous.bar.T)

	if !time_current.After(time_previous) {
		return nil, Date_is_not_greater_then_previous
	}
	has_collision1, intersection := o.CrossCheck(previous, *newIchi)
	if has_collision1 == EInterSectionStatus_Find {
		newIchi.SetStatus(newIchi.CloudStatus(intersection))
		newIchi.SetCloudSwitching(o.isSwitchCloud(previous, *newIchi))

		return newIchi, nil
	}
	return nil, nil

}

//analyse with two days
func (o *IchimokuDriver) PreAnalyseIchimoku(data []IchimokuStatus) (*IchimokuStatus, error) {

	if len(data) != 2 {
		return nil, NotEnoughData
	}

	current := data[0]
	previous := data[1]

	latest := IchimokuStatus{}

	line1_point_a := NewPoint(float64(previous.bar.T), previous.TenkenSen.valLine)
	line1_point_b := NewPoint(float64(current.bar.T), current.TenkenSen.valLine)

	line2_point_a := NewPoint(float64(previous.bar.T), previous.KijonSen.valLine)
	line2_point_b := NewPoint(float64(current.bar.T), current.KijonSen.valLine)

	has_collision1, intersection := o.line_helper.GetCollisionDetection(line1_point_a, line1_point_b, line2_point_a, line2_point_b)

	if has_collision1 == EInterSectionStatus(IchimokuStatus_NAN) {
		return nil, nil
	}

	Line_Eq_A := o.getLineEquation(line1_point_a, line1_point_b) // tenken
	Line_Eq_B := o.getLineEquation(line2_point_a, line2_point_b) //kijon

	if line1_point_a == line2_point_a {
		return nil, nil //paraller

	}

	if Line_Eq_A.Slope-Line_Eq_B.Slope == 0 {
		return nil, nil

	}

	if has_collision1 == EInterSectionStatus_Find {

		current.SetStatus(current.CloudStatus(intersection))

		current.SetCloudSwitching(o.isSwitchCloud(previous, current))
		latest = current
		return &latest, nil
	}
	return nil, nil
}
func (o *IchimokuDriver) NumberOfIchimoku() int {
	return len(o.QIchimokuDays)
}
func (o *IchimokuDriver) GetLastDay() *IchimokuStatus {

	if o.QIchimokuDays != nil && len(o.QIchimokuDays) == 0 {
		return nil
	}

	latest := o.QIchimokuDays[0]

	return &latest
}
func (o *IchimokuDriver) Put(v *IchimokuStatus) bool {

	//final := []IchimokuStatus{}
	if len(o.QIchimokuDays) == 0 {
		o.QIchimokuDays = append(o.QIchimokuDays, *v)
	} else {
		o.QIchimokuDays = append(o.QIchimokuDays, *v)
	}

	return true
}

//--------------------------------------------------------------------------------------------------------
func (o *IchimokuDriver) CrossCheck(previous IchimokuStatus, newIchi IchimokuStatus) (EInterSectionStatus, float64) {
	line1_point_a := NewPoint(float64(previous.bar.T), previous.TenkenSen.valLine)
	line1_point_b := NewPoint(float64(newIchi.bar.T), newIchi.TenkenSen.valLine)

	line2_point_a := NewPoint(float64(previous.bar.T), previous.KijonSen.valLine)
	line2_point_b := NewPoint(float64(newIchi.bar.T), newIchi.KijonSen.valLine)

	has_collision1, intersection := o.line_helper.GetCollisionDetection(line1_point_a, line1_point_b, line2_point_a, line2_point_b)

	if has_collision1 == EInterSectionStatus(IchimokuStatus_NAN) {
		return EInterSectionStatus_NAN, 0.0
	}

	Line_Eq_A := o.getLineEquation(line1_point_a, line1_point_b) // tenken
	Line_Eq_B := o.getLineEquation(line2_point_a, line2_point_b) //kijon

	if line1_point_a == line2_point_a {
		return EInterSectionStatus_NAN, intersection //paraller

	}

	if Line_Eq_A.Slope-Line_Eq_B.Slope == 0 {
		return EInterSectionStatus_NAN, intersection

	}
	return has_collision1, intersection
}

//analyse with two days
func (o *IchimokuDriver) isSwitchCloud(previous IchimokuStatus, current IchimokuStatus) bool {

	if previous.SenKoA_Shifted26.valLine <= previous.SenKoB_Shifted26.valLine &&
		current.SenKoA_Shifted26.valLine > current.SenKoB_Shifted26.valLine ||
		previous.SenKoA_Shifted26.valLine < previous.SenKoB_Shifted26.valLine && current.SenKoA_Shifted26.valLine >= current.SenKoB_Shifted26.valLine {
		if current.TenkenSen.valLine >= current.KijonSen.valLine {
			return true
		}
	}
	return false
}

//
// find cloud  Span A,B in Past (26 day )
//
func (o *IchimokuDriver) Calc_Cloud_InPast(current int) (Point, Point) {

	if o.NumberOfIchimoku() < 26 {
		return NewNilPoint(), NewNilPoint()
	}

	rem := o.NumberOfIchimoku() - current
	max := 26 //from 26 bar in past  (find Shift index)

	if rem < max {
		return NewNilPoint(), NewNilPoint()
	}

	index := current + 25
	c := o.QIchimokuDays[index]
	buff_senco_a := NewPoint(float64(c.bar.T/1000), c.SenKoA_Shifted26.valLine)
	buff_senco_b := NewPoint(float64(c.bar.T/1000), c.SenKoB_Shifted26.valLine)

	return buff_senco_a, buff_senco_b
}

// private 1
//
func (o *IchimokuDriver) calcLine(line_type ELine, bars []Bar) ValueLine {
	high := NewValueLineNil()
	low := NewValueLineNil()
	if len(bars) < int(line_type) {
		return NewValueLineNil()
	}
	for _, v := range bars {
		if high.isNil {
			high.SetValue(v.H)
		}
		if low.isNil {
			low.SetValue(v.L)
		}

		if v.H > high.valLine {
			high.SetValue(v.H)
		}

		if v.L < low.valLine {
			low.SetValue(v.L)
		}

	}
	line := (low.valLine + high.valLine) / 2
	return NewValue(line)
}

func (o *IchimokuDriver) calculate_span_a(tenken ValueLine, kijon ValueLine) ValueLine {
	if tenken.isNil == false && kijon.isNil == false {
		v := (tenken.valLine + kijon.valLine) / 2
		return NewValue(v)
	}

	return NewValueLineNil()
}

func (o *IchimokuDriver) getLineEquation(p1 Point, p2 Point) *Equation {
	eq := Equation{}
	eq.Slope = (p2.Y - p1.Y) / (p2.X - p1.X)
	//eq.Intercept = (-1 * eq.Slope * p1.X) + p1.Y
	eq.Intercept = p1.Y - eq.Slope*p1.X
	return &eq
}
