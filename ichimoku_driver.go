package ichimoku

import (
	"fmt"
	"sort"
)

type IIchimokuDriver interface {
	//
	// out result array sorte : [new day -to- old day]
	IchimokuRun(bars []Bar) ([]IchimokuStatus, error)
	AnalyseIchimoku(data []IchimokuStatus) (*IchimokuStatus, error)
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

//
// out result array sorte : [new day -to- old day]
func (o *IchimokuDriver) IchimokuRun(bars []Bar) ([]IchimokuStatus, error) {

	if len(bars) == 0 {
		return nil, DataNotFill
	}
	fmt.Println("calc ichi from Last 100 days")
	o.bars = bars
	days := []IchimokuStatus{}
	bars_len := len(o.bars) - 1
	for day := 0; day < 135; day++ {

		//for day := 3; day >= 0; day-- {
		tenkenLine := o.calcLine(Line_Tenkan_sen, o.loadbars(bars_len-int(Line_Tenkan_sen)-day, bars_len-day))
		kijonLine := o.calcLine(Line_kijon_sen, o.loadbars(bars_len-int(Line_kijon_sen)-day, bars_len-day))

		span_a := o.calculate_span_a(tenkenLine, kijonLine)
		span_b := o.calcLine(Line_spanPeriod, o.loadbars(bars_len-int(Line_spanPeriod)-day, bars_len-day))

		chiko_index := len(o.bars) - int(Line_chikoPeriod) - day
		var chiko ValueLine

		if chiko_index >= 0 && len(o.bars) > chiko_index {
			chiko.SetValue(o.bars[chiko_index].C)
		} else {
			if len(days) == 0 {
				return nil, NotEnoughData
			} else {
				break
			}

		}
		var latestPrice Bar
		latestPriceIndex := bars_len - day
		if len(o.bars) >= latestPriceIndex {
			latestPrice = o.bars[bars_len-day]
		} else {
			if len(days) == 0 {
				return nil, NotEnoughData
			} else {
				break
			}

		}

		if !tenkenLine.isNil && !kijonLine.isNil && !span_a.isNil && !span_b.isNil {

			ichi := NewIchimokuStatus(tenkenLine, kijonLine, span_a, span_b, chiko, latestPrice)
			days = append(days, *ichi)
		} else {
			if len(days) == 0 {
				return nil, NotEnoughData
			} else {
				break
			}
		}
	}

	for i := 0; i < len(days); i++ {
		item := &days[i]
		sen_a, sen_b := o.find_Clouds_InPast(i, &days)
		item.Set_SenCo_A_inPast(sen_a)
		item.Set_SenCo_B_inPast(sen_b)

	}
	return days, nil

}

//analyse with two days

func (o *IchimokuDriver) AnalyseIchimoku(data []IchimokuStatus) (*IchimokuStatus, error) {

	if len(data) != 2 {
		return nil, NotEnoughData
	}

	today := data[0]
	yesterday := data[1]

	latest := IchimokuStatus{}

	line1_point_a := NewPoint(float64(yesterday.bar.T), yesterday.TenkenSen.valLine)
	line1_point_b := NewPoint(float64(today.bar.T), today.TenkenSen.valLine)

	line2_point_a := NewPoint(float64(yesterday.bar.T), yesterday.KijonSen.valLine)
	line2_point_b := NewPoint(float64(today.bar.T), today.KijonSen.valLine)

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

		today.SetStatus(today.CloudStatus(intersection))

		//  else if today.Above(today.bar.C) {
		// 	today.SetStatus(IchimokuStatus_Cross_Above)
		// }
		// else {
		// 	fmt.Printf("TK cross found but not classified for ")
		// }

		if o.price_action_leaving_cloud(today, yesterday) && today.Is_cloud_green() {
			today.SetLeavingCloud(true)
		}

		if yesterday.SencoA.valLine <= yesterday.SencoB.valLine &&
			today.SencoA.valLine > today.SencoB.valLine ||
			yesterday.SencoA.valLine < yesterday.SencoB.valLine && today.SencoA.valLine >= today.SencoB.valLine {
			if today.TenkenSen.valLine >= today.KijonSen.valLine {
				today.SetCloudSwitching(true)
			}
		}
		latest = today
		return &latest, nil
	}
	return nil, nil
}

//
//analyse with 26 day or more
//
func (o *IchimokuDriver) DeepTimeAnalyse(data []IchimokuStatus) (*IchimokuStatus, error) {

	if len(data) != 0 || len(data) < 54 {
		return nil, NotEnoughData
	}
	//first_cross_after_26 := false
	for i := 0; i < len(data); i++ {
		s := data[i]

		if s.bar.C < s.KijonSen.valLine {

		}

		return nil, nil
	}

	return nil, nil
}

//
// find cloud  Span A,B in Past (26 day )
//
func (o *IchimokuDriver) find_Clouds_InPast(current int, days *[]IchimokuStatus) ([]Point, []Point) {

	if len(*days) < 26 {
		return nil, nil
	}
	var buff_senco_a []Point
	var buff_senco_b []Point

	rem := len(*days) - current
	max := 26
	//tail := len(*days) - current
	if rem > max {
		buff_senco_a = make([]Point, max)
		buff_senco_b = make([]Point, max)

	} else {
		buff_senco_a = make([]Point, rem)
		buff_senco_b = make([]Point, rem)
		max = rem
	}
	defer func() {
		buff_senco_a = nil
		buff_senco_b = nil
	}()

	counter := max - 1
	start := current
	i := start
	for {

		if counter < 0 {
			break
		}
		if rem <= 0 {
			break
		}

		buff_senco_a[counter] = NewPoint(float64((*days)[i].bar.T/1000), (*days)[i].SencoA.valLine)
		buff_senco_b[counter] = NewPoint(float64((*days)[i].bar.T/1000), (*days)[i].SencoB.valLine)
		counter--
		i++

	}

	//for i := 25; i > 0; i-- {

	// buff_senco_a = append(buff_senco_a, NewPoint(float64((*days)[i].bar.T/1000), (*days)[i].SencoA.valLine))

	// buff_senco_b = append(buff_senco_b, NewPoint(float64((*days)[i].bar.T/1000), (*days)[i].SencoB.valLine))
	//}
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

func (o *IchimokuDriver) price_action_leaving_cloud(today IchimokuStatus, yesterday IchimokuStatus) bool {
	if o.inside_range([]float64{yesterday.bar.H, yesterday.bar.L}, []float64{yesterday.SencoA.valLine, yesterday.SencoB.valLine}) {
		var comparison float64

		if today.Is_cloud_green() {
			comparison = today.SencoA.valLine
		} else {
			return false
		}
		if today.bar.C > comparison {
			return true
		}
	}
	return false
}

func (o *IchimokuDriver) inside_range(data []float64, range1 []float64) bool {
	return o.minArr(data) > o.minArr(range1) && o.maxArr(data) < o.maxArr(range1)
}

func (o *IchimokuDriver) minArr(v []float64) float64 {
	sort.Float64s(v)
	return v[0]
}

func (o *IchimokuDriver) maxArr(v []float64) float64 {
	sort.Float64s(v)
	return v[len(v)-1]
}
