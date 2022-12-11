package tests

import (
	"testing"

	"github.com/qasemt/ichimoku"
	"github.com/stretchr/testify/assert"
)

func TestInside(t *testing.T) {

	//fmt.Println("bars ", bars)
	driver := ichimoku.NewIchimokuDriver()
	//ctl := gomock.NewController(t)
	//	d := mock_ichimoku.NewMockIIchimokuDriver(ctl)
	//	fmt.Println("a", d)

	today := ichimoku.NewIchimokuStatus(ichimoku.NewValue(8705), ichimoku.NewValue(8710), ichimoku.NewValue(8707), ichimoku.NewValue(8930), ichimoku.NewValue(8830), ichimoku.Bar{L: 8400, H: 8460, C: 8440, O: 8440, V: 906352, T: 1664699400000})

	yesterday := ichimoku.NewIchimokuStatus(ichimoku.NewValue(8720), ichimoku.NewValue(8710), ichimoku.NewValue(8715), ichimoku.NewValue(8940), ichimoku.NewValue(8870), ichimoku.Bar{
		L: 8430, H: 8480, C: 8450, O: 8460, V: 652416, T: 1664695800000})

	lines_result := make([]ichimoku.IchimokuStatus, 2)
	lines_result[0] = *today //today
	lines_result[1] = *yesterday

	a, e := driver.AnalyseIchimoku(lines_result)
	assert.Empty(t, e)
	assert.Equal(t, a.Status, ichimoku.IchimokuStatus_Cross_Above)

}

func TestCheckCloud_Above(t *testing.T) {

	h1_above_ := []ichimoku.Point{
		{X: 1667802600, Y: 8762},
		{X: 1667806200, Y: 8882},
		{X: 1667809800, Y: 8908},
		{X: 1667885400, Y: 8928},
		{X: 1667889000, Y: 9152},
		{X: 1667892600, Y: 9222},
		{X: 1667896200, Y: 9238},
		{X: 1667971800, Y: 9238},
		{X: 1667975400, Y: 9260},
		{X: 1667979000, Y: 9285},
		{X: 1667982600, Y: 9318},
		{X: 1668231000, Y: 9318},
		{X: 1668234600, Y: 9318},
		{X: 1668238200, Y: 9320},
		{X: 1668241800, Y: 9320},
		{X: 1668317400, Y: 9358},
		{X: 1668321000, Y: 9358},
		{X: 1668324600, Y: 9410},
		{X: 1668328200, Y: 9485},
		{X: 1668403800, Y: 9485},
		{X: 1668407400, Y: 9435},
		{X: 1668411000, Y: 9430},
		{X: 1668414600, Y: 9490},
		{X: 1668490200, Y: 9498},
		{X: 1668493800, Y: 9482},
	}
	line_helper := ichimoku.NewLineHelper()
	point_from_price := ichimoku.NewPoint(float64(1668497400/1000), 9670)
	res_senko_a, err := line_helper.GetCollisionWithLine(point_from_price, h1_above_)

	//fmt.Println("senko A", res_senko_a, "err:", err)
	assert.Equal(t, res_senko_a, ichimoku.EPointLocation_above)
	assert.Empty(t, err)

}
func TestCheckCloud_below(t *testing.T) {

	h1_above_ := []ichimoku.Point{
		{X: 1666589400, Y: 8660},
		{X: 1666593000, Y: 8660},
		{X: 1666596600, Y: 8618},
		{X: 1666600200, Y: 8555},
		{X: 1666675800, Y: 8548},
		{X: 1666679400, Y: 8502},
		{X: 1666683000, Y: 8435},
		{X: 1666686600, Y: 8435},
		{X: 1666762200, Y: 8430},
		{X: 1666765800, Y: 8360},
		{X: 1666769400, Y: 8300},
		{X: 1666773000, Y: 8282},
		{X: 1667021400, Y: 8282},
		{X: 1667025000, Y: 8282},
		{X: 1667028600, Y: 8240},
		{X: 1667032200, Y: 8160},
		{X: 1667107800, Y: 8120},
		{X: 1667111400, Y: 8120},
		{X: 1667115000, Y: 8112},
		{X: 1667118600, Y: 8110},
		{X: 1667194200, Y: 8098},
		{X: 1667197800, Y: 8100},
		{X: 1667201400, Y: 8060},
		{X: 1667205000, Y: 8052},
		{X: 1667280600, Y: 8055},
	}
	line_helper := ichimoku.NewLineHelper()
	point_from_price := ichimoku.NewPoint(float64(1667284200/1000), 8360)
	res_senko_a, err := line_helper.GetCollisionWithLine(point_from_price, h1_above_)

	assert.Equal(t, res_senko_a, ichimoku.EPointLocation_below)
	assert.Empty(t, err)

}
func TestCheckCloud_below1(t *testing.T) {
	// h1 shegoya Tue Oct  2022 18 10:00:00

	h1_above_ := []ichimoku.Point{
		{X: 1665819000, Y: 8745},
		{X: 1665822600, Y: 8750},
		{X: 1665898200, Y: 8750},
		{X: 1665901800, Y: 8730},
		{X: 1665905400, Y: 8725},
		{X: 1665909000, Y: 8712},
		{X: 1665984600, Y: 8692},
		{X: 1665988200, Y: 8692},
		{X: 1665991800, Y: 8680},
		{X: 1665995400, Y: 8680},
		{X: 1666071000, Y: 8680},
	}
	line_helper := ichimoku.NewLineHelper()
	point_from_price := ichimoku.NewPoint(float64(1666074600/1000), 8650)
	res_senko_a, err := line_helper.GetCollisionWithLine(point_from_price, h1_above_)

	assert.Equal(t, res_senko_a, ichimoku.EPointLocation_below)
	assert.Empty(t, err)

}
func TestCheckCloud_below3(t *testing.T) {
	//|2022 Tue Nov 1 10:00:00
	h1_above_ := []ichimoku.Point{
		{X: 1666589400, Y: 8660},
		{X: 1666593000, Y: 8660},
		{X: 1666596600, Y: 8618},
		{X: 1666600200, Y: 8555},
		{X: 1666675800, Y: 8548},
		{X: 1666679400, Y: 8502},
		{X: 1666683000, Y: 8435},
		{X: 1666686600, Y: 8435},
		{X: 1666762200, Y: 8430},
		{X: 1666765800, Y: 8360},
		{X: 1666769400, Y: 8300},
		{X: 1666773000, Y: 8282},
		{X: 1667021400, Y: 8282},
		{X: 1667025000, Y: 8282},
		{X: 1667028600, Y: 8240},
		{X: 1667032200, Y: 8160},
		{X: 1667107800, Y: 8120},
		{X: 1667111400, Y: 8120},
		{X: 1667115000, Y: 8112},
		{X: 1667118600, Y: 8110},
		{X: 1667194200, Y: 8098},
		{X: 1667197800, Y: 8100},
		{X: 1667201400, Y: 8060},
		{X: 1667205000, Y: 8052},
		{X: 1667280600, Y: 8055},
	}
	line_helper := ichimoku.NewLineHelper()
	point_from_price := ichimoku.NewPoint(float64(1667284200/1000), 8360)
	res_senko_a, err := line_helper.GetCollisionWithLine(point_from_price, h1_above_)

	//fmt.Println("senko A", res_senko_a, "err:", err)
	assert.Equal(t, res_senko_a, ichimoku.EPointLocation_below)
	assert.Empty(t, err)

}
