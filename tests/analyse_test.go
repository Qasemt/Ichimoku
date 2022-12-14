package tests

import (
	"testing"

	"github.com/qasemt/ichimoku"
	"github.com/stretchr/testify/assert"
)

var (
	//52 bar
	bar_h1 = []ichimoku.Bar{
		{L: 8110, H: 8180, C: 8160, O: 8110, V: 664372.00, T: 1667201400000},
		{L: 8100, H: 8260, C: 8200, O: 8150, V: 1241301.00, T: 1667205000000},
		{L: 8110, H: 8450, C: 8440, O: 8170, V: 2909458.00, T: 1667280600000},
		{L: 8310, H: 8450, C: 8360, O: 8440, V: 778238.00, T: 1667284200000},
		{L: 8240, H: 8370, C: 8260, O: 8360, V: 658420.00, T: 1667287800000},
		{L: 8240, H: 8450, C: 8440, O: 8260, V: 1814124.00, T: 1667291400000},
		{L: 8270, H: 8440, C: 8300, O: 8440, V: 1267103.00, T: 1667367000000},
		{L: 8270, H: 8510, C: 8510, O: 8300, V: 1821017.00, T: 1667370600000},
		{L: 8430, H: 8540, C: 8440, O: 8510, V: 559250.00, T: 1667374200000},
		{L: 8420, H: 8470, C: 8440, O: 8440, V: 544851.00, T: 1667377800000},
		{L: 8480, H: 8730, C: 8730, O: 8550, V: 4284720.00, T: 1667626200000},
		{L: 8730, H: 8730, C: 8730, O: 8730, V: 1382828.00, T: 1667629800000},
		{L: 8730, H: 8730, C: 8730, O: 8730, V: 1678201.00, T: 1667633400000},
		{L: 8730, H: 8730, C: 8730, O: 8730, V: 549277.00, T: 1667637000000},
		{L: 8800, H: 9070, C: 9060, O: 8800, V: 5342062.00, T: 1667712600000},
		{L: 9040, H: 9070, C: 9070, O: 9060, V: 8126959.00, T: 1667716200000},
		{L: 9070, H: 9070, C: 9070, O: 9070, V: 527101.00, T: 1667719800000},
		{L: 9070, H: 9070, C: 9070, O: 9070, V: 702521.00, T: 1667723400000},
		{L: 9160, H: 9440, C: 9430, O: 9290, V: 4409696.00, T: 1667799000000},
		{L: 9410, H: 9490, C: 9490, O: 9420, V: 7522839.00, T: 1667802600000},
		{L: 9490, H: 9490, C: 9490, O: 9490, V: 777299.00, T: 1667806200000},
		{L: 9490, H: 9490, C: 9490, O: 9490, V: 405416.00, T: 1667809800000},
		{L: 9300, H: 9890, C: 9530, O: 9890, V: 7097789.00, T: 1667885400000},
		{L: 9460, H: 9570, C: 9470, O: 9520, V: 3033312.00, T: 1667889000000},
		{L: 9380, H: 9490, C: 9410, O: 9470, V: 2714433.00, T: 1667892600000},
		{L: 9390, H: 9490, C: 9450, O: 9420, V: 3876877.00, T: 1667896200000},
		{L: 9250, H: 9540, C: 9410, O: 9350, V: 3448605.00, T: 1667971800000},
		{L: 9400, H: 9840, C: 9800, O: 9410, V: 6547559.00, T: 1667975400000},
		{L: 9640, H: 9830, C: 9650, O: 9800, V: 2416825.00, T: 1667979000000},
		{L: 9650, H: 9860, C: 9680, O: 9700, V: 2463503.00, T: 1667982600000},
		{L: 9640, H: 9870, C: 9800, O: 9750, V: 2000789.00, T: 1668231000000},
		{L: 9520, H: 9800, C: 9520, O: 9780, V: 3214849.00, T: 1668234600000},
		{L: 9520, H: 9680, C: 9620, O: 9550, V: 3019512.00, T: 1668238200000},
		{L: 9610, H: 9810, C: 9740, O: 9640, V: 2473212.00, T: 1668241800000},
		{L: 9450, H: 9710, C: 9530, O: 9710, V: 1455003.00, T: 1668317400000},
		{L: 9510, H: 9700, C: 9700, O: 9520, V: 1341450.00, T: 1668321000000},
		{L: 9520, H: 9720, C: 9650, O: 9700, V: 2922575.00, T: 1668324600000},
		{L: 9470, H: 9650, C: 9470, O: 9650, V: 907574.00, T: 1668328200000},
		{L: 9250, H: 9620, C: 9250, O: 9510, V: 1573592.00, T: 1668403800000},
		{L: 9220, H: 9420, C: 9380, O: 9270, V: 1372258.00, T: 1668407400000},
		{L: 9340, H: 9530, C: 9490, O: 9380, V: 3147032.00, T: 1668411000000},
		{L: 9370, H: 9550, C: 9370, O: 9490, V: 2153637.00, T: 1668414600000},
		{L: 9380, H: 9750, C: 9670, O: 9450, V: 1861478.00, T: 1668490200000},
		{L: 9580, H: 9700, C: 9650, O: 9670, V: 2890813.00, T: 1668493800000},
		{L: 9610, H: 9700, C: 9670, O: 9610, V: 1288957.00, T: 1668497400000},
		{L: 9630, H: 9800, C: 9730, O: 9650, V: 2413843.00, T: 1668501000000},
		{L: 9580, H: 9780, C: 9630, O: 9750, V: 803830.00, T: 1668576600000},
		{L: 9630, H: 9720, C: 9670, O: 9650, V: 699785.00, T: 1668580200000},
		{L: 9640, H: 9700, C: 9640, O: 9700, V: 393592.00, T: 1668583800000},
		{L: 9580, H: 9660, C: 9630, O: 9640, V: 1443871.00, T: 1668587400000},
		{L: 9300, H: 9600, C: 9370, O: 9510, V: 3845936.00, T: 1668835800000},
		{L: 9310, H: 9380, C: 9330, O: 9380, V: 1380628.00, T: 1668839400000},
	}
)

func Test_26DayinThePastNotExist(t *testing.T) {

	//fmt.Println("bars ", bars)
	driver := ichimoku.NewIchimokuDriver()

	driver.MakeIchimokuInPast(&bar_h1, 1)

	_, e1 := driver.AnalyseTriggerCross(*driver.GetLastDay(), bar_h1)

	assert.Equal(t, e1, ichimoku.ChikoStatus26InPastNotMade)

}
func TestInside(t *testing.T) {

	//fmt.Println("bars ", bars)
	driver := ichimoku.NewIchimokuDriver()
	//ctl := gomock.NewController(t)
	//	d := mock_ichimoku.NewMockIIchimokuDriver(ctl)
	//	fmt.Println("a", d)

	today := ichimoku.NewIchimokuStatus(ichimoku.NewValue(8705), ichimoku.NewValue(8710), ichimoku.NewValue(8707), ichimoku.NewValue(8930), ichimoku.Bar{}, ichimoku.Bar{L: 8400, H: 8460, C: 8440, O: 8440, V: 906352, T: 1664699400000})

	yesterday := ichimoku.NewIchimokuStatus(ichimoku.NewValue(8720), ichimoku.NewValue(8710), ichimoku.NewValue(8715), ichimoku.NewValue(8940), ichimoku.Bar{}, ichimoku.Bar{
		L: 8430, H: 8480, C: 8450, O: 8460, V: 652416, T: 1664695800000})

	lines_result := make([]ichimoku.IchimokuStatus, 2)
	lines_result[0] = *today //today
	lines_result[1] = *yesterday

	a, e := driver.PreAnalyseIchimoku(lines_result)
	assert.Empty(t, e)
	assert.Equal(t, a.Status, ichimoku.IchimokuStatus_Cross_Above)

}
