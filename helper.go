package ichimoku

import (
	"fmt"
)

type lineHelper struct {
}
type EPointLocation int

const (
	EPointLocation_NAN     EPointLocation = 0
	EPointLocation_above   EPointLocation = 1
	EPointLocation_below   EPointLocation = 2
	EPointLocation_overlap EPointLocation = 3
)

func NewLineHelper() lineHelper {
	f := lineHelper{}
	return f
}

//check Point above or below line
func (n *lineHelper) isAboveLine(point Point, line []Point) EPointLocation {

	if len(line) < 2 {
		return EPointLocation_NAN
	}

	checkAboveOrBelow := func(p Point, line_p_a Point, line_p_b Point) EPointLocation {
		//var d = (_p_x- _x1) *(_y2-_y1)-(_p_y-_y1)*(_x2-_x1)
		v := (p.X-line_p_a.X)*(line_p_b.Y-line_p_a.Y) - (p.Y-line_p_a.Y)*(line_p_b.X-line_p_a.X)
		if v == 0 {
			return EPointLocation_overlap
		} else if v > 0 {
			return EPointLocation_above
		} else {
			return EPointLocation_below
		}

	}
	//lastest_location := EPointLocation_NAN

	//for i := 1; i < len(line); i++ {

	//res := checkAboveOrBelow(point, line[i-1], line[i])
	res := checkAboveOrBelow(point, line[0], line[1])
	//}

	return res

}
func (n *lineHelper) positionPointInline(point Point, line []Point) float64 {

	if len(line) < 2 {
		return 0
	}

	checkAboveOrBelow := func(p Point, line_p_a Point, line_p_b Point) float64 {
		//var d = (_p_x- _x1) *(_y2-_y1)-(_p_y-_y1)*(_x2-_x1)
		v := (p.X-line_p_a.X)*(line_p_b.Y-line_p_a.Y) - (p.Y-line_p_a.Y)*(line_p_b.X-line_p_a.X)
		return v

	}

	res := checkAboveOrBelow(point, line[0], line[1])

	return res

}
func (o *lineHelper) GetCollisionDetection(a Point, b Point, c Point, d Point) (EInterSectionStatus, float64) {

	denominator := ((b.X - a.X) * (d.Y - c.Y)) - ((b.Y - a.Y) * (d.X - c.X))
	numerator1 := ((a.Y - c.Y) * (d.X - c.X)) - ((a.X - c.X) * (d.Y - c.Y))
	numerator2 := ((a.Y - c.Y) * (b.X - a.X)) - ((a.X - c.X) * (b.Y - a.Y))

	// Detect coincident lines (has a problem, read below)
	if denominator == 0 {
		return EInterSectionStatus_NAN, 0
	}
	r := numerator1 / denominator
	s := numerator2 / denominator

	if (r >= 0 && r <= 1) && (s >= 0 && s <= 1) {
		//	fmt.Printf("collision detec : a:%v , b:%v, c:%v ,d:%v ,r %v s %v\r\n", a, b, c, d, r, s)
		intersection := o.get_intersection_point(a, b, c, d)

		return EInterSectionStatus_Collision_Find, intersection
	}
	return EInterSectionStatus_NAN, 0
}

//line senco A or B
func (o *lineHelper) GetCollisionWithLine(price_point Point, line_clouds []Point) (EPointLocation, error) {

	len_line_clouds := len(line_clouds)
	if len_line_clouds < 1 {
		return EPointLocation_NAN, NotEnoughData
	}
	// Create a point at infinity, y is same as point p
	//x := time.Now().AddDate(0, -1, 0).Unix()
	//line1_a := NewPoint(float64(x), price_point.Y)
	//	line1_b := price_point
	//  line_b := NewPoint(0,price_point)
	//var ps EPointLocation = EPointLocation_NAN
	below := 0
	above := 0
	fmt.Println("___")
	fmt.Printf("Cloud : check point :x:%.0f y:%.0f \r\n", price_point.X, price_point.Y)
	sum := 0.0
	for i := 1; i < len_line_clouds; i++ {
		line2_a := line_clouds[i-1]
		line2_b := line_clouds[i]
		//fmt.Printf("Cloud :x:%.0f y:%.0f ,\r\n", line2_a.X, line2_a.Y)
		fmt.Printf("%.0f,%.0f,\r\n", line2_a.X, line2_a.Y)
		//res := o.GetCollisionDetection(line1_a, line1_b, line2_a, line2_b)
		buff := []Point{line2_a, line2_b}
		// res := o.isAboveLine(price_point, buff)
		c := o.positionPointInline(price_point, buff)
		if c > 0 { //below
			below++
		} else {
			above++
		}

		sum += c

	}
	v := sum / float64(len_line_clouds)

	if v > 0 {
		return EPointLocation_below, nil // above
	} else {
		return EPointLocation_above, nil //below
	}

}
func (o *lineHelper) getLineEquation(p1 Point, p2 Point) *Equation {
	eq := Equation{}
	eq.Slope = (p1.Y - p2.Y) / (p1.X - p2.X)
	eq.Intercept = (-1 * eq.Slope * p1.X) + p1.Y
	return &eq
}

func (o *lineHelper) get_intersection_point(line1_pointA Point, line1_pointB Point, line2_pointA Point, line2_pointB Point) float64 {

	tenken := o.getLineEquation(line1_pointA, line1_pointB)
	kijon := o.getLineEquation(line2_pointA, line2_pointB)
	x_intersection := (kijon.Intercept - tenken.Intercept) / (tenken.Slope - kijon.Slope)
	y_intersection := (kijon.Slope * x_intersection) + kijon.Intercept
	return y_intersection
}
