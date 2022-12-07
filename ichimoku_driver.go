package ichimoku

import (
	"sort"
)

type IIchimokuDriver interface {
	IchimokuRun(bars []Bar) ([]IchimokuStatus, error)
	AnalyseIchimoku(data []IchimokuStatus) (*IchimokuStatus, error)
	GetIntersectionPoint(line1_pointA Point, line1_pointB Point, line2_pointA Point, line2_pointB Point) (EInterSectionStatus, float64)
	GetCollisionDetection(line1_pointA Point, line1_pointB Point, line2_pointA Point, line2_pointB Point) EInterSectionStatus
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

	latest := IchimokuStatus{}

	line1_point_a := NewPoint(1.0, yesterday.TenkenSen.valLine)
	line1_point_b := NewPoint(2.0, today.TenkenSen.valLine)

	line2_point_a := NewPoint(1.0, yesterday.KijonSen.valLine)
	line2_point_b := NewPoint(2.0, today.KijonSen.valLine)

	has_collision1 := o.GetCollisionDetection(line1_point_a, line1_point_b, line2_point_a, line2_point_b)

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
		if today.Below(today.TenkenSen.valLine) {
			today.SetStatus(IchimokuStatus_Cross_Below)
		} else if today.Above(today.TenkenSen.valLine) {
			today.SetStatus(IchimokuStatus_Cross_Above)
		}
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
func (o *IchimokuDriver) GetIntersectionPoint(line1_pointA Point, line1_pointB Point, line2_pointA Point, line2_pointB Point) (EInterSectionStatus, float64) {

	Line_Eq_A := o.getLineEquation(line1_pointA, line1_pointB) // tenken
	Line_Eq_B := o.getLineEquation(line2_pointA, line2_pointB) //kijon

	if line1_pointA == line2_pointA {
		return EInterSectionStatus_Parallel, -1

	}

	if Line_Eq_A.Slope-Line_Eq_B.Slope == 0 {
		return EInterSectionStatus_Parallel, -1

	}
	aa := (Line_Eq_A.Intercept - Line_Eq_B.Intercept)
	bb := (Line_Eq_B.Slope - Line_Eq_A.Slope)
	x_intersection := aa / bb
	y_intersection := Line_Eq_A.Slope*x_intersection + Line_Eq_A.Intercept

	round_intersection_point := (x_intersection + y_intersection) / 2

	round_point_b := (line2_pointB.X + line2_pointB.Y) / 2
	if Line_Eq_A.Slope == 0 && Line_Eq_B.Slope == 0 {
		return EInterSectionStatus_Parallel, 0
	} else if x_intersection > 0 && round_point_b > round_intersection_point {
		return EInterSectionStatus_Find, y_intersection
	} else if x_intersection < 0 && round_point_b < round_intersection_point {
		return EInterSectionStatus_Find, y_intersection
	}
	//fmt.Printf("Point of intersection is x:%.2f , y:%.2f ", x_intersection, y_intersection)

	return EInterSectionStatus_NAN, 0
}
func (o *IchimokuDriver) GetCollisionDetection(a Point, b Point, c Point, d Point) EInterSectionStatus {

	denominator := ((b.X - a.X) * (d.Y - c.Y)) - ((b.Y - a.Y) * (d.X - c.X))
	numerator1 := ((a.Y - c.Y) * (d.X - c.X)) - ((a.X - c.X) * (d.Y - c.Y))
	numerator2 := ((a.Y - c.Y) * (b.X - a.X)) - ((a.X - c.X) * (b.Y - a.Y))

	// Detect coincident lines (has a problem, read below)
	if denominator == 0 {
		return EInterSectionStatus_NAN
	}
	r := numerator1 / denominator
	s := numerator2 / denominator

	if (r >= 0 && r <= 1) && (s >= 0 && s <= 1) {
		//	fmt.Printf("collision detec : a:%v , b:%v, c:%v ,d:%v ,r %v s %v\r\n", a, b, c, d, r, s)
		return EInterSectionStatus_Collision_Find
	}
	return EInterSectionStatus_NAN
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
