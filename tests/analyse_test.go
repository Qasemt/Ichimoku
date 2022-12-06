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

	today := ichimoku.NewIchimokuStatus(ichimoku.NewValue(8380), ichimoku.NewValue(8380), ichimoku.NewValue(8380), ichimoku.NewValue(8865), ichimoku.NewValue(8440), ichimoku.Bar{L: 8380, H: 8610, C: 8440})

	yesterday := ichimoku.NewIchimokuStatus(ichimoku.NewValue(8210), ichimoku.NewValue(8375), ichimoku.NewValue(8292.5), ichimoku.NewValue(8865), ichimoku.NewValue(8430), ichimoku.Bar{L: 7820, H: 8210, C: 8200})

	lines_result := make([]ichimoku.IchimokuStatus, 2)
	lines_result[0] = *today //today
	lines_result[1] = *yesterday

	a, e := driver.AnalyseIchimoku(lines_result)
	assert.Empty(t, e)
	assert.Equal(t, a.Status, ichimoku.IchimokuStatus_Cross_Inside)

}
func TestCrossBelow(t *testing.T) {

	//fmt.Println("bars ", bars)
	driver := ichimoku.NewIchimokuDriver()
	//ctl := gomock.NewController(t)
	//	d := mock_ichimoku.NewMockIIchimokuDriver(ctl)
	//	fmt.Println("a", d)

	today := ichimoku.NewIchimokuStatus(ichimoku.NewValue(8380), ichimoku.NewValue(8380), ichimoku.NewValue(8380), ichimoku.NewValue(8865), ichimoku.NewValue(8440), ichimoku.Bar{L: 8380, H: 8610, C: 8440})

	yesterday := ichimoku.NewIchimokuStatus(ichimoku.NewValue(8210), ichimoku.NewValue(8375), ichimoku.NewValue(8292.5), ichimoku.NewValue(8865), ichimoku.NewValue(8430), ichimoku.Bar{L: 7820, H: 8210, C: 8200})

	lines_result := make([]ichimoku.IchimokuStatus, 2)
	lines_result[1] = *today //today
	lines_result[0] = *yesterday

	a, e := driver.AnalyseIchimoku(lines_result)
	assert.Empty(t, e)
	assert.Equal(t, a.Status, ichimoku.IchimokuStatus_Cross_Below)

}

var (
	bars = []ichimoku.Bar{

		{L: 9250, H: 9520, C: 9420, O: 9440, V: 7302087.00},
		{L: 9230, H: 9550, C: 9420, O: 9230, V: 6538362.00},
		{L: 9350, H: 9710, C: 9520, O: 9350, V: 5471852.00},
		{L: 9390, H: 9680, C: 9560, O: 9390, V: 2749313.00},
		{L: 9510, H: 9640, C: 9600, O: 9510, V: 2589092.00},
		{L: 9600, H: 9770, C: 9740, O: 9610, V: 4428946.00},
		{L: 9400, H: 9740, C: 9590, O: 9740, V: 3276725.00},
		{L: 9500, H: 9640, C: 9580, O: 9500, V: 2932533.00},
		{L: 9500, H: 10080, C: 9900, O: 9510, V: 6478283.00},
		{L: 9660, H: 10040, C: 9760, O: 9950, V: 5556492.00},
		{L: 9700, H: 9870, C: 9790, O: 9710, V: 3083011.00},
		{L: 9550, H: 9830, C: 9650, O: 9790, V: 4739627.00},
		{L: 9520, H: 9770, C: 9560, O: 9770, V: 5348730.00},
		{L: 9500, H: 9700, C: 9610, O: 9700, V: 4583396.00},
		{L: 9460, H: 9650, C: 9470, O: 9540, V: 6816560.00},
		{L: 9440, H: 9600, C: 9600, O: 9480, V: 5141584.00},
		{L: 9490, H: 9690, C: 9690, O: 9530, V: 4459637.00},
		{L: 9620, H: 9990, C: 9810, O: 9700, V: 7028425.00},
		{L: 9750, H: 10040, C: 9920, O: 9820, V: 7785084.00},
		{L: 9700, H: 9980, C: 9770, O: 9880, V: 4293447.00},
		{L: 9690, H: 9920, C: 9840, O: 9850, V: 5501951.00},
		{L: 9700, H: 9880, C: 9760, O: 9800, V: 4744251.00},
		{L: 9670, H: 9840, C: 9770, O: 9800, V: 4551177.00},
		{L: 9540, H: 9790, C: 9630, O: 9780, V: 4848313.00},
		{L: 9530, H: 9830, C: 9650, O: 9530, V: 5265632.00},
		{L: 9600, H: 9780, C: 9640, O: 9630, V: 3783557.00},
		{L: 9500, H: 9760, C: 9510, O: 9760, V: 5197515.00},
		{L: 9420, H: 9570, C: 9510, O: 9530, V: 4914418.00},
		{L: 9250, H: 9570, C: 9440, O: 9500, V: 6445395.00},
		{L: 9260, H: 9500, C: 9360, O: 9440, V: 3790118.00},
		{L: 9210, H: 9470, C: 9390, O: 9310, V: 4496608.00},
		{L: 9260, H: 9450, C: 9380, O: 9300, V: 3220819.00},
		{L: 9200, H: 9360, C: 9240, O: 9300, V: 4977710.00},
		{L: 8950, H: 9240, C: 9070, O: 9240, V: 7873285.00},
		{L: 8910, H: 9160, C: 8920, O: 8960, V: 3622806.00},
		{L: 8820, H: 9060, C: 8880, O: 8900, V: 4017179.00},
		{L: 8800, H: 8950, C: 8930, O: 8810, V: 4882262.00},
		{L: 8470, H: 8920, C: 8490, O: 8850, V: 5917669.00},
		{L: 8360, H: 8900, C: 8740, O: 8360, V: 3881415.00},
		{L: 8720, H: 9060, C: 8990, O: 8720, V: 3785299.00},
		{L: 8400, H: 8990, C: 8430, O: 8950, V: 6916905.00},
		{L: 8380, H: 8610, C: 8440, O: 8460, V: 4114188.00},
		{L: 8420, H: 8600, C: 8500, O: 8490, V: 3232725.00},
		{L: 8400, H: 8560, C: 8520, O: 8400, V: 3198319.00},
		{L: 8400, H: 8650, C: 8620, O: 8400, V: 4032326.00},
		{L: 8360, H: 8620, C: 8550, O: 8460, V: 6554159.00},
		{L: 8420, H: 8590, C: 8540, O: 8500, V: 4092876.00},
		{L: 8490, H: 8900, C: 8890, O: 8490, V: 6791976.00},
		{L: 8720, H: 8960, C: 8950, O: 8900, V: 5813867.00},
		{L: 8730, H: 8960, C: 8860, O: 8750, V: 5235316.00},
		{L: 8490, H: 8880, C: 8500, O: 8870, V: 6206200.00},
		{L: 8500, H: 8700, C: 8690, O: 8500, V: 2815106.00},
		{L: 8600, H: 8750, C: 8700, O: 8600, V: 2754533.00},
		{L: 8600, H: 8730, C: 8700, O: 8620, V: 2832941.00},
		{L: 8580, H: 8840, C: 8610, O: 8580, V: 3182758.00},
		{L: 8450, H: 8610, C: 8450, O: 8610, V: 3558469.00},
		{L: 8290, H: 8580, C: 8340, O: 8580, V: 5726211.00},
		{L: 8150, H: 8450, C: 8190, O: 8450, V: 6551287.00},
		{L: 7920, H: 8260, C: 8020, O: 8150, V: 12189455.00},
		{L: 7690, H: 8230, C: 7760, O: 8230, V: 12333716.00},
		{L: 7820, H: 8210, C: 8200, O: 7950, V: 13835111.00},
		{L: 8070, H: 8260, C: 8200, O: 8070, V: 3551433.00},
		{L: 8110, H: 8450, C: 8440, O: 8170, V: 6160240.00},
		{L: 8270, H: 8540, C: 8440, O: 8440, V: 4192221.00},
		{L: 8480, H: 8730, C: 8730, O: 8550, V: 7895026.00},
		{L: 8800, H: 9070, C: 9070, O: 8800, V: 14698643.00},
		{L: 9160, H: 9490, C: 9490, O: 9290, V: 13115250.00},
		{L: 9300, H: 9890, C: 9450, O: 9890, V: 16722411.00},
		{L: 9250, H: 9860, C: 9680, O: 9350, V: 14876492.00},
		{L: 9520, H: 9870, C: 9740, O: 9750, V: 10704587.00},
		{L: 9450, H: 9720, C: 9470, O: 9710, V: 6626602.00},
		{L: 9220, H: 9620, C: 9370, O: 9510, V: 8234037.00},
		{L: 9380, H: 9800, C: 9730, O: 9450, V: 8455091.00},
		{L: 9580, H: 9780, C: 9630, O: 9750, V: 3622640.00},
		{L: 9580, H: 9780, C: 9630, O: 9750, V: 3622640.00},
		{L: 9300, H: 9600, C: 9480, O: 9510, V: 7048062.00},
		{L: 9320, H: 9780, C: 9630, O: 9510, V: 4334180.00},
		{L: 9500, H: 9700, C: 9510, O: 9600, V: 3777194.00},
		{L: 9400, H: 9520, C: 9450, O: 9520, V: 2425998.00},
		{L: 9430, H: 9580, C: 9480, O: 9430, V: 4214377.00},
		{L: 9330, H: 9560, C: 9400, O: 9560, V: 2692258.00},
		{L: 9300, H: 9520, C: 9430, O: 9400, V: 5115817.00},
		{L: 9360, H: 9550, C: 9390, O: 9430, V: 5320065.00},
		{L: 9360, H: 9600, C: 9430, O: 9470, V: 10541962.00},
		{L: 9350, H: 9590, C: 9450, O: 9450, V: 6573233.00},
	}
)
