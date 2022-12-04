package ichimoku

import (
	"fmt"
	"sort"
)

type IIchimokuDriver interface {
	IchimokuRun(bars []Bar) ([]IchimokuStatus, error)
	PrintResult()
	AnalyseIchimoku(lines_ichi []IchimokuStatus) (*IchimokuStatus, error)
}

type IchimokuDriver struct {
	bars           []Bar
	ConversionLine []float64
	BaseLine       []float64
	LeadingSpanA   []float64
	LeadingSpanB   []float64
	laggingSpan    []float64
}

func NewIchimokuDriver() IIchimokuDriver {
	xx := IchimokuDriver{}

	return &xx
}

func (xx *IchimokuDriver) load(from int, to int) []Bar {
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
func (o *IchimokuDriver) AnalyseIchimoku(data []IchimokuStatus) (*IchimokuStatus, error) {

	if len(data) != 2 {
		return nil, NotEnoughData
	}

	today := data[0]
	yesterday := data[1]
	var intersection float64
	latest := IchimokuStatus{}
	if (yesterday.TenkenSen.valLine <= yesterday.KijonSen.valLine && today.TenkenSen.valLine > today.KijonSen.valLine) || (yesterday.TenkenSen.valLine < yesterday.KijonSen.valLine && today.TenkenSen.valLine >= today.KijonSen.valLine) {

		if today.KijonSen.valLine == today.TenkenSen.valLine {
			intersection = today.KijonSen.valLine
		} else {
			intersection = o.get_intersection_point(yesterday.TenkenSen.valLine, today.TenkenSen.valLine, yesterday.KijonSen.valLine, today.KijonSen.valLine)
		}
		if today.Below(intersection) {
			today.SetStatus(IchimokuStatus_Cross_Below)
		} else if today.inside(intersection) {
			today.SetStatus(IchimokuStatus_Cross_Inside)
		} else if today.Above(intersection) {
			today.SetStatus(IchimokuStatus_Cross_Above)
		} else {
			fmt.Printf("TK cross found but not classified for ")
		}

		if o.price_action_leaving_cloud(today, yesterday) && today.Is_cloud_green() {
			today.SetLeavingCloud(true)
		}

		//    if yesterday_ichi.leading_span_a <= yesterday_ichi.leading_span_b and today_ichi.leading_span_a > today_ichi.leading_span_b or yesterday_ichi.leading_span_a < yesterday_ichi.leading_span_b and today_ichi.leading_span_a >= today_ichi.leading_span_b:

		if yesterday.SencoA.valLine <= yesterday.SencoB.valLine && today.SencoA.valLine > today.SencoB.valLine || yesterday.SencoA.valLine < yesterday.SencoB.valLine && today.SencoA.valLine >= today.SencoB.valLine {
			if today.TenkenSen.valLine >= today.KijonSen.valLine {
				today.SetFolding(true)
			}
		}
		latest = today
		return &latest, nil
	} else {
		return nil, nil
	}

}
func (xx *IchimokuDriver) IchimokuRun(bars []Bar) ([]IchimokuStatus, error) {

	if len(bars) == 0 {
		return nil, DataNotFill
	}

	xx.bars = bars
	// TurningLine
	days := []IchimokuStatus{}

	for day := 0; day < 100; day++ {
		lenx := len(xx.bars) - 1
		//for day := 3; day >= 0; day-- {
		tenkenLine := xx.calcLine(Lin_Tenkan_sen, xx.load(lenx-int(Lin_Tenkan_sen)-day, lenx-day))
		kijonLine := xx.calcLine(Line_kijon_sen, xx.load(lenx-int(Line_kijon_sen)-day, lenx-day))
		span_a := xx.calculate_span_a(tenkenLine, kijonLine)
		span_b := xx.calcLine(Line_spanPeriod, xx.load(lenx-int(Line_spanPeriod)-day, lenx-day))

		chiko_index := len(xx.bars) - int(Line_chikoPeriod) - day
		var chiko ValueLine

		if chiko_index >= 0 && len(xx.bars) > chiko_index {
			chiko.SetValue(xx.bars[chiko_index].Close)
		} else {
			if len(days) == 0 {
				return nil, NotEnoughData
			} else {
				break
			}

		}
		var latestPrice Bar
		latestPriceIndex := lenx - day
		if len(xx.bars) >= latestPriceIndex {
			latestPrice = xx.bars[lenx-day]
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
	return days, nil

}
func (xx *IchimokuDriver) PrintResult() {
	// fmt.Printf("ichi %v|%v|%v|%v|%v|G:%v,Chiko UP :%v \r\n", it.TenkenSen.Value(), it.KijonSen.Value(), it.SencoA.Value(), it.SencoB.Value(), it.Chiko.Value(), it.Is_cloud_green(), it.IsChikoAbovePrice())
}

// private 1
func (o *IchimokuDriver) calcLine(line_type ELine, bars []Bar) ValueLine {
	high := NewValueLineNil()
	low := NewValueLineNil()
	if len(bars) < int(line_type) {
		return NewValueLineNil()
	}
	for _, v := range bars {
		if high.isNil {
			high.SetValue(v.High)
		}
		if low.isNil {
			low.SetValue(v.Low)
		}

		if v.High > high.valLine {
			high.SetValue(v.High)
		}

		if v.Low < low.valLine {
			low.SetValue(v.Low)
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

func (o *IchimokuDriver) get_intersection_point(a float64, b float64, x float64, y float64) float64 {

	conversion := o.get_line_equation([]float64{0, a}, []float64{1, b})
	base := o.get_line_equation([]float64{0, x}, []float64{1, y})
	x_intersection := (base.Intercept - conversion.Intercept) / (conversion.Slope - base.Slope)
	y_intersection := (base.Slope * x_intersection) + base.Intercept
	return y_intersection
}

func (o *IchimokuDriver) get_line_equation(p1 []float64, p2 []float64) *Equation {
	eq := Equation{}
	eq.Slope = (p2[1] - p1[1]) / (p2[0] - p1[0])
	eq.Intercept = (-1 * eq.Slope * p1[0]) + p1[1]
	return &eq
}
func (o *IchimokuDriver) price_action_leaving_cloud(today IchimokuStatus, yesterday IchimokuStatus) bool {
	if o.inside_range([]float64{yesterday.bar.High, yesterday.bar.Low}, []float64{yesterday.SencoA.valLine, yesterday.SencoB.valLine}) {
		var comparison float64

		if today.Is_cloud_green() {
			comparison = today.SencoA.valLine
		} else {
			return false
		}
		if today.bar.Close > comparison {
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
