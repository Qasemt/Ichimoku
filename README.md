## [![Go](https://github.com/Qasemt/Ichimoku/actions/workflows/go.yml/badge.svg)](https://github.com/Qasemt/Ichimoku/actions/workflows/go.yml)

Source :

- [Intersection Point Of Two Lines][1]

### Calculation

```

There are five plots that make up the Ichimoku Cloud indicator. Their names and calculations are:

TenkanSen (Conversion Line): (High + Low) / 2 default period = 9
KijunSen (Base Line): (High + Low) / 2 default period = 26
Chiku Span (Lagging Span): Price Close shifted back 26 bars
Senkou A (Leading Span A): (TenkanSen + KijunSen) / 2 (Senkou A is shifted forward 26 bars)
Senkou B (Leading Span B): (High + Low) / 2 using period = 52 (Senkou B is shifted forward 26 bars)
```

---

## function Intersection Point Of Two Lines

```golang
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
```

![alt text](./docs/demo_h1.png)

Result :

```console

calc ichi from Last 100 days

 Find ichi 8630|8630|8630|8715|8450|G:false,Chiko UP :true |status : cross below |cloud switching : false|leaving cloud : false |2022 Tue Oct 11 10:00:00 |1665469800000
____
 Find ichi 8685|8690|8687.5|8710|8500|G:false,Chiko UP :true |status : cross inside |cloud switching : false|leaving cloud : false |2022 Tue Oct 18 10:00:00 |1666074600000
____
 Find ichi 8135|8135|8135|8300|8520|G:false,Chiko UP :false |status : cross below |cloud switching : false|leaving cloud : false |2022 Tue Nov 1 10:00:00 |1667284200000
____
 Find ichi 9485|9525|9505|8790|9490|G:true,Chiko UP :true |status : cross above |cloud switching : false|leaving cloud : false |2022 Tue Nov 15 11:00:00 |1668497400000
____
 Find ichi 9570|9545|9557.5|8960|9410|G:true,Chiko UP :true |status : cross above |cloud switching : false|leaving cloud : false |2022 Wed Nov 16 12:00:00 |1668587400000

```

---

# Ichimoku

ichimoku indicator

## sample

```
var (
		bar_h1 := []ichimoku.Bar{
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

[1]: https://web.archive.org/web/20060911055655/http://local.wasp.uwa.edu.au/~pbourke/geometry/lineline2d/
