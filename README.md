## [![Go](https://github.com/Qasemt/Ichimoku/actions/workflows/go.yml/badge.svg)](https://github.com/Qasemt/Ichimoku/actions/workflows/go.yml)

# Ichimoku

ichimoku indicator

## sample

```
var (
	bars = []ichimoku.Bar{
        .
        .
        .
        old bars
        .
		{Low: 9250,H: 9520, Close: 9420},
		{Low: 9230,H: 9550, Close: 9420},
		{Low: 9350,H: 9710, Close: 9520},
		{Low: 9390,H: 9680, Close: 9560}
        .
        .
        .
        new bars  -> over 52 candles
        }
    )

	driver := ichimoku.NewIchimokuDriver()

	arr, err := driver.IchimokuRun(bars)
	if err != nil {
		fmt.Println("error :", err)
	}
	for _, it := range arr {
		fmt.Printf("%v\r\n", it.Print())
	}

   today := arr[0]
   yesterday := arr[1]

   lines_result := make([]ichimoku.IchimokuStatus, 2)
   lines_result[0] = today //today
   lines_result[1] = yesterday

    a, e := driver.AnalyseIchimoku(lines_result)

```
