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

	today := ichimoku.NewIchimokuStatus(ichimoku.NewValue(-1), ichimoku.NewValue(-1), ichimoku.NewValue(8380), ichimoku.NewValue(2), ichimoku.NewValue(2), ichimoku.Bar{L: 8380, H: 8610, C: 8440})

	yesterday := ichimoku.NewIchimokuStatus(ichimoku.NewValue(-0.5), ichimoku.NewValue(-2), ichimoku.NewValue(1.5), ichimoku.NewValue(2), ichimoku.NewValue(8430), ichimoku.Bar{L: 7820, H: 8210, C: 8200})

	lines_result := make([]ichimoku.IchimokuStatus, 2)
	lines_result[0] = *today //today
	lines_result[1] = *yesterday

	a, e := driver.AnalyseIchimoku(lines_result)
	assert.Empty(t, e)
	assert.Equal(t, a.Status, ichimoku.IchimokuStatus_Cross_Below)

}
