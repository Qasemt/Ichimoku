package ichimoku

import (
	"fmt"
	"time"
)

type IIchimokuDriver interface {
	Init(bars *[]Bar, numberOfRead int) ([]IchimokuStatus, error)
	//
	// out result array sorte : [new day -to- old day]
	//once call per asset
	MakeIchimokuInPast(bars *[]Bar, numberOfRead int) ([]IchimokuStatus, error)

	//once call per asset
	PreAnalyseIchimoku(data []IchimokuStatus) (*IchimokuStatus, error)

	//Analys Trigger Cross tenken & kijon sen
	AnalyseTriggerCross(previous IchimokuStatus, bar26FromLatest IchimokuStatus, bars []Bar) (*IchimokuStatus, error)
}

type IchimokuDriver struct {
	line_helper    lineHelper
	bars           []Bar
	ConversionLine []float64
	BaseLine       []float64
	LeadingSpanA   []float64
	LeadingSpanB   []float64
	laggingSpan    []float64
}

func NewIchimokuDriver() IIchimokuDriver {
	xx := IchimokuDriver{}
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
func (o *IchimokuDriver) Init(bars *[]Bar, numberOfRead int) ([]IchimokuStatus, error) {

	a, e := o.MakeIchimokuInPast(bars, numberOfRead)
	if e != nil {
		return nil, e
	}

	o.ReViewBar(&a)

	return a, nil
}

//
// out result array sorte : [new day -to- old day]
func (o *IchimokuDriver) MakeIchimokuInPast(bars *[]Bar, numberOfRead int) ([]IchimokuStatus, error) {

	if len(*bars) == 0 {
		return nil, DataNotFill
	}
	if len(*bars) < numberOfRead {
		return nil, NotEnoughData
	}
	fmt.Printf("Calc ichi from Last %v days", numberOfRead)
	o.bars = *bars

	//days:sorted descending
	days := []IchimokuStatus{}
	bars_index_with_zero := len(o.bars)
	for day := 0; day < numberOfRead; day++ {

		//for day := 3; day >= 0; day-- {
		tenkenLine := o.calcLine(Line_Tenkan_sen, o.loadbars(bars_index_with_zero-int(Line_Tenkan_sen)-day, bars_index_with_zero-day))
		kijonLine := o.calcLine(Line_kijon_sen, o.loadbars(bars_index_with_zero-int(Line_kijon_sen)-day, bars_index_with_zero-day))

		span_a := o.calculate_span_a(tenkenLine, kijonLine)
		span_b := o.calcLine(Line_spanPeriod, o.loadbars(bars_index_with_zero-int(Line_spanPeriod)-day, bars_index_with_zero-day))
		chiko_index := (bars_index_with_zero - int(Line_chikoPeriod) - day) - 1
		cheko_span := o.bars[chiko_index]

		var latestPrice Bar
		latestPriceIndex := bars_index_with_zero - day
		if len(o.bars) >= latestPriceIndex {
			latestPrice = o.bars[(bars_index_with_zero-1)-day]
		} else {
			if len(days) == 0 {
				return nil, NotEnoughData
			} else {
				break
			}

		}

		if !tenkenLine.isNil && !kijonLine.isNil && !span_a.isNil && !span_b.isNil {

			ichi := NewIchimokuStatus(tenkenLine, kijonLine, span_a, span_b, cheko_span, latestPrice)
			days = append(days, *ichi)
		} else {
			if len(days) == 0 {
				return nil, NotEnoughData
			} else {
				break
			}
		}
	}

	return days, nil

}

//analyse with two days
func (o *IchimokuDriver) ReViewBar(days *[]IchimokuStatus) {

	for day_index := 0; day_index < len(*days); day_index++ {
		item := &(*days)[day_index]
		sen_a, sen_b := o.Calc_Cloud_InPast(day_index, days)
		item.Set_SenCo_A_Past(sen_a)
		item.Set_SenCo_B_Past(sen_b)
	}

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
func (o *IchimokuDriver) Calc_Cloud_InPast(current int, days *[]IchimokuStatus) (Point, Point) {

	if len(*days) < 26 {
		return NewNilPoint(), NewNilPoint()
	}

	rem := len(*days) - current
	max := 26 //from 26 bar in past  (find Shift index)

	if rem < max {
		return NewNilPoint(), NewNilPoint()
	}

	index := current + 25
	c := (*days)[index]
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

//Analys Trigger Cross tenken & kijon sen
func (o *IchimokuDriver) AnalyseTriggerCross(previous IchimokuStatus, bar26FromLatest IchimokuStatus, bars []Bar) (*IchimokuStatus, error) {

	find, e := o.MakeIchimokuInPast(&o.bars, 52)

	if e != nil {
		return nil, e
	}

	if bar26FromLatest.SencoA_Past.isNil || bar26FromLatest.SencoB_Past.isNil {
		return nil, DataNotFill
	}

	if len(find) == 0 {
		return nil, nil
	}
	current := find[0]

	time_current := time.UnixMilli(current.bar.T)
	time_previous := time.UnixMilli(previous.bar.T)

	if !time_current.After(time_previous) {
		return nil, Date_is_not_greater_then_previous
	}

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
