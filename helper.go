package ichimoku

type lineHelper struct {
}
type EPointLocation int

const (
	EPointLocation_NAN     EPointLocation = 1
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

	var checkAboveOrBelow = func(p Point, line_p_a Point, line_p_b Point) EPointLocation {
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
	var lastest_location EPointLocation = EPointLocation_above
	for i := 1; i < len(line); i++ {

		lastest_location = checkAboveOrBelow(point, line[i-1], line[i])
	}

	return lastest_location

}
func (o *lineHelper) GetCollisionDetection(a Point, b Point, c Point, d Point) EInterSectionStatus {

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

//line senco A or B
func (o *lineHelper) GetCollisionWithLine(price_point Point, line_clouds []Point) (bool, error) {

	len_line_clouds := len(line_clouds)
	if len_line_clouds < 1 {
		return false, NotEnoughData
	}
	// Create a point at infinity, y is same as point p
	//x := time.Now().AddDate(0, -1, 0).Unix()
	//line1_a := NewPoint(float64(x), price_point.Y)
	//	line1_b := price_point
	//  line_b := NewPoint(0,price_point)
	//var ps EPointLocation = EPointLocation_NAN
	below := 0
	above := 0
	for i := 1; i < len_line_clouds; i++ {
		line2_a := line_clouds[i-1]
		line2_b := line_clouds[i]
		//fmt.Printf("Cloud :x:%.0f y:%.0f \r\n", line2_a.X, line2_a.Y)
		//res := o.GetCollisionDetection(line1_a, line1_b, line2_a, line2_b)
		buff := []Point{line2_a, line2_b}
		res := o.isAboveLine(price_point, buff)

		switch res {
		case EPointLocation_below:
			below++
		case EPointLocation_above:
			above++
		}

	}
	if below > 0 {
		return true, nil
	}
	// return EInterSectionStatus_NAN
	return false, nil
}
