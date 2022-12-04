package main

import (
	"fmt"

	"github.com/qasemt/ichimoku"
)

func main() {

	bars := []ichimoku.Bar{

		{Low: 9250, High: 9520, Close: 9420},
		{Low: 9230, High: 9550, Close: 9420},
		{Low: 9350, High: 9710, Close: 9520},
		{Low: 9390, High: 9680, Close: 9560},
		{Low: 9510, High: 9640, Close: 9600},
		{Low: 9600, High: 9770, Close: 9740},
		{Low: 9400, High: 9740, Close: 9590},
		{Low: 9500, High: 9640, Close: 9580},
		{Low: 9500, High: 10080, Close: 9900},
		{Low: 9660, High: 10040, Close: 9760},
		{Low: 9700, High: 9870, Close: 9790},
		{Low: 9550, High: 9830, Close: 9650},
		{Low: 9520, High: 9770, Close: 9560},
		{Low: 9500, High: 9700, Close: 9610},
		{Low: 9460, High: 9650, Close: 9470},
		{Low: 9440, High: 9600, Close: 9600},
		{Low: 9490, High: 9690, Close: 9690},
		{Low: 9620, High: 9990, Close: 9810},
		{Low: 9750, High: 10040, Close: 9920},
		{Low: 9700, High: 9980, Close: 9770},
		{Low: 9690, High: 9920, Close: 9840},
		{Low: 9700, High: 9880, Close: 9760},
		{Low: 9670, High: 9840, Close: 9770},
		{Low: 9540, High: 9790, Close: 9630},
		{Low: 9530, High: 9830, Close: 9650},
		{Low: 9600, High: 9780, Close: 9640},
		{Low: 9500, High: 9760, Close: 9510},
		{Low: 9420, High: 9570, Close: 9510},
		{Low: 9250, High: 9570, Close: 9440},
		{Low: 9260, High: 9500, Close: 9360},
		{Low: 9210, High: 9470, Close: 9390},
		{Low: 9260, High: 9450, Close: 9380},
		{Low: 9200, High: 9360, Close: 9240},
		{Low: 8950, High: 9240, Close: 9070},
		{Low: 8910, High: 9160, Close: 8920},
		{Low: 8820, High: 9060, Close: 8880},
		{Low: 8800, High: 8950, Close: 8930},
		{Low: 8470, High: 8920, Close: 8490},
		{Low: 8360, High: 8900, Close: 8740},
		{Low: 8720, High: 9060, Close: 8990},
		{Low: 8400, High: 8990, Close: 8430},
		{Low: 8380, High: 8610, Close: 8440},
		{Low: 8420, High: 8600, Close: 8500},
		{Low: 8400, High: 8560, Close: 8520},
		{Low: 8400, High: 8650, Close: 8620},
		{Low: 8360, High: 8620, Close: 8550},
		{Low: 8420, High: 8590, Close: 8540},
		{Low: 8490, High: 8900, Close: 8890},
		{Low: 8720, High: 8960, Close: 8950},
		{Low: 8730, High: 8960, Close: 8860},
		{Low: 8490, High: 8880, Close: 8500},
		{Low: 8500, High: 8700, Close: 8690},
		{Low: 8600, High: 8750, Close: 8700},
		{Low: 8600, High: 8730, Close: 8700},
		{Low: 8580, High: 8840, Close: 8610},
		{Low: 8450, High: 8610, Close: 8450},
		{Low: 8290, High: 8580, Close: 8340},
		{Low: 8150, High: 8450, Close: 8190},
		{Low: 7920, High: 8260, Close: 8020},
		{Low: 7690, High: 8230, Close: 7760},
		{Low: 7820, High: 8210, Close: 8200},
		{Low: 8070, High: 8260, Close: 8200},
		{Low: 8110, High: 8450, Close: 8440},
		{Low: 8270, High: 8540, Close: 8440},
		{Low: 8480, High: 8730, Close: 8730},
		{Low: 8800, High: 9070, Close: 9070},
		{Low: 9160, High: 9490, Close: 9490},
		{Low: 9300, High: 9890, Close: 9450},
		{Low: 9250, High: 9860, Close: 9680},
		{Low: 9520, High: 9870, Close: 9740},
		{Low: 9450, High: 9720, Close: 9470},
		{Low: 9220, High: 9620, Close: 9370},
		{Low: 9380, High: 9800, Close: 9730},
		{Low: 9580, High: 9780, Close: 9630},
		{Low: 9580, High: 9780, Close: 9630},
		{Low: 9300, High: 9600, Close: 9480},
		{Low: 9320, High: 9780, Close: 9630},
		{Low: 9500, High: 9700, Close: 9510},
		{Low: 9400, High: 9520, Close: 9450},
		{Low: 9430, High: 9580, Close: 9480},
		{Low: 9330, High: 9560, Close: 9400},
		{Low: 9300, High: 9520, Close: 9430},
		{Low: 9360, High: 9550, Close: 9390},
		{Low: 9360, High: 9600, Close: 9430},
		{Low: 9350, High: 9590, Close: 9450},
	}

	fmt.Println(len(bars))

	driver := ichimoku.NewIchimokuDriver()

	arr, err := driver.IchimokuRun(bars)
	if err != nil {
		fmt.Println("error :", err)
	}
	for _, it := range arr {
		fmt.Printf("%v\r\n", it.Print())
	}

	lines_result := make([]ichimoku.IchimokuStatus, 2)

	for i := len(arr) - 2; i > 0; i-- {
		today := arr[i]
		pre := arr[i+1]
		lines_result[0] = today //today
		lines_result[1] = pre

		a, e := driver.AnalyseIchimoku(lines_result)

		if e != nil {
			fmt.Println("err", e)
		}
		if a != nil {
			fmt.Printf("____ \r\n Find %v \r\n", a.Print())
		}
	}

}
