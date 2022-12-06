package ichimoku

import (
	"fmt"
	"sort"
)

type IIchimokuDriver interface {
	IchimokuRun(bars []Bar) ([]IchimokuStatus, error)
	AnalyseIchimoku(lines_ichi []IchimokuStatus) (*IchimokuStatus, error)
	GetIntersectionPoint(line1_pointA Point, line1_pointB Point, line2_pointA Point, line2_pointB Point) float64
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

func (xx *IchimokuDriver) IchimokuRun(bars []Bar) ([]IchimokuStatus, error) {

	if len(bars) == 0 {
		return nil, DataNotFill
	}

	//previous_cross_with_kijon := false
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
			chiko.SetValue(xx.bars[chiko_index].C)
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

			if kijonLine.isNil == false && latestPrice.O < kijonLine.valLine && latestPrice.C > kijonLine.valLine {
				ichi.SetCrossKijonAndPrice(true)
			}
			if kijonLine.isNil == false && latestPrice.C < kijonLine.valLine && latestPrice.O > kijonLine.valLine {
				ichi.SetCrossKijonAndPrice(true)
			}
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
func (o *IchimokuDriver) AnalyseIchimoku(data []IchimokuStatus) (*IchimokuStatus, error) {

	if len(data) != 2 {
		return nil, NotEnoughData
	}

	today := data[0]
	yesterday := data[1]
	var intersection float64
	latest := IchimokuStatus{}
	// if yesterday.TenkenSen.valLine == 9515 && yesterday.KijonSen.valLine == 9480 {
	// 	fmt.Println("a")
	// }

	G := yesterday.TenkenSen.valLine <= yesterday.KijonSen.valLine && today.TenkenSen.valLine > today.KijonSen.valLine
	R := yesterday.TenkenSen.valLine < yesterday.KijonSen.valLine && today.TenkenSen.valLine >= today.KijonSen.valLine

	// R := yesterday.TenkenSen.valLine > today.TenkenSen.valLine && today.KijonSen.valLine >= yesterday.KijonSen.valLine
	// G := yesterday.TenkenSen.valLine < today.TenkenSen.valLine && yesterday.KijonSen.valLine <= today.KijonSen.valLine
	if G || R {

		// if yesterday.bar.T == 1668490200000 {
		// 	fmt.Println("a")
		// }
		if today.KijonSen.valLine == today.TenkenSen.valLine {
			intersection = today.KijonSen.valLine
		} else {
			line1_point_a := NewPoint(0, yesterday.TenkenSen.valLine)
			line1_point_b := NewPoint(1, today.TenkenSen.valLine)
			line2_point_a := NewPoint(0, yesterday.KijonSen.valLine)
			line2_point_b := NewPoint(1, today.KijonSen.valLine)
			intersection = o.get_intersection_point(line1_point_a, line1_point_b, line2_point_a, line2_point_b)
		}
		if today.Below(intersection) {
			today.SetStatus(IchimokuStatus_Cross_Below)
		} else if today.Above(intersection) {
			today.SetStatus(IchimokuStatus_Cross_Above)
		} else {
			fmt.Printf("TK cross found but not classified for ")
		}

		if o.price_action_leaving_cloud(today, yesterday) && today.Is_cloud_green() {
			today.SetLeavingCloud(true)
		}

		if yesterday.SencoA.valLine <= yesterday.SencoB.valLine && today.SencoA.valLine > today.SencoB.valLine || yesterday.SencoA.valLine < yesterday.SencoB.valLine && today.SencoA.valLine >= today.SencoB.valLine {
			if today.TenkenSen.valLine >= today.KijonSen.valLine {
				today.SetCloudSwitching(true)
			}
		}
		latest = today
		return &latest, nil
	} else {
		return nil, nil
	}

}

//analyse with 26 day or more
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

// private 1
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

func (o *IchimokuDriver) get_intersection_point(line1_pointA Point, line1_pointB Point, line2_pointA Point, line2_pointB Point) float64 {

	tenken := o.getLineEquation(line1_pointA, line1_pointB)
	kijon := o.getLineEquation(line2_pointA, line2_pointB)
	x_intersection := (kijon.Intercept - tenken.Intercept) / (tenken.Slope - kijon.Slope)
	y_intersection := (kijon.Slope * x_intersection) + kijon.Intercept
	return y_intersection
}
func (o *IchimokuDriver) GetIntersectionPoint(line1_pointA Point, line1_pointB Point, line2_pointA Point, line2_pointB Point) float64 {

	var m1, m2 float64
	m1 = (line1_pointB.Y-line1_pointA.Y)/(line1_pointB.X) - line1_pointA.X
	m2 = (line2_pointB.Y-line2_pointA.Y)/(line2_pointB.X) - line2_pointA.X

	if m1 == m2 {
		fmt.Println("Lines are parallel")
	} else if m1 == -(1 / m2) {
		var b1, b2, x1, y1, inx, inb float64

		b1 = line1_pointA.Y - m1*line1_pointA.X
		b2 = line2_pointA.Y - m2*line2_pointA.X

		inx = (m1 - m2)
		inb = (b2 - b1)
		x1 = inb / inx
		y1 = m1*x1 + b1

		fmt.Println("Point of intersection is", x1, ",", y1)
	} else {
		fmt.Println("Lines are neither parallel or  perpendicular")

	}
	return 0
}
func (o *IchimokuDriver) getLineEquation(p1 Point, p2 Point) *Equation {
	eq := Equation{}
	eq.Slope = (p1.Y - p2.Y) / (p1.X - p2.X)
	eq.Intercept = (-1 * eq.Slope * p1.X) + p1.Y
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
