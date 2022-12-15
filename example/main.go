package main

import (
	"fmt"

	"github.com/qasemt/ichimoku"
)

func main() {

	// sort bar_h1 [old date to new date] ascending
	bar_h1 := []ichimoku.Bar{
		{L: 9450, H: 9570, C: 9490, O: 9530, V: 1158096.00, T: 1662525000000},
		{L: 9450, H: 9550, C: 9480, O: 9530, V: 1041077.00, T: 1662528600000},
		{L: 9420, H: 9520, C: 9500, O: 9500, V: 1966815.00, T: 1662532200000},
		{L: 9480, H: 9520, C: 9510, O: 9480, V: 748430.00, T: 1662535800000},
		{L: 9250, H: 9500, C: 9390, O: 9500, V: 1961310.00, T: 1662784200000},
		{L: 9340, H: 9570, C: 9520, O: 9350, V: 1881148.00, T: 1662787800000},
		{L: 9300, H: 9520, C: 9320, O: 9520, V: 2174744.00, T: 1662791400000},
		{L: 9320, H: 9440, C: 9440, O: 9320, V: 428193.00, T: 1662795000000},
		{L: 9280, H: 9500, C: 9470, O: 9440, V: 1174617.00, T: 1662870600000},
		{L: 9300, H: 9490, C: 9330, O: 9480, V: 564590.00, T: 1662874200000},
		{L: 9260, H: 9360, C: 9300, O: 9330, V: 1325195.00, T: 1662877800000},
		{L: 9280, H: 9360, C: 9360, O: 9290, V: 725716.00, T: 1662881400000},
		{L: 9290, H: 9470, C: 9310, O: 9310, V: 1156762.00, T: 1662957000000},
		{L: 9210, H: 9330, C: 9250, O: 9310, V: 1463096.00, T: 1662960600000},
		{L: 9240, H: 9440, C: 9380, O: 9240, V: 1350073.00, T: 1662964200000},
		{L: 9360, H: 9400, C: 9390, O: 9380, V: 526677.00, T: 1662967800000},
		{L: 9260, H: 9450, C: 9280, O: 9300, V: 685445.00, T: 1663043400000},
		{L: 9270, H: 9450, C: 9440, O: 9280, V: 957947.00, T: 1663047000000},
		{L: 9340, H: 9440, C: 9340, O: 9440, V: 716902.00, T: 1663050600000},
		{L: 9300, H: 9380, C: 9380, O: 9350, V: 860525.00, T: 1663054200000},
		{L: 9230, H: 9360, C: 9250, O: 9300, V: 1844161.00, T: 1663129800000},
		{L: 9230, H: 9290, C: 9240, O: 9250, V: 876713.00, T: 1663133400000},
		{L: 9210, H: 9280, C: 9260, O: 9240, V: 909086.00, T: 1663137000000},
		{L: 9200, H: 9270, C: 9240, O: 9260, V: 1347750.00, T: 1663140600000},
		{L: 9060, H: 9240, C: 9100, O: 9240, V: 1513067.00, T: 1663475400000},
		{L: 9060, H: 9150, C: 9080, O: 9080, V: 1770320.00, T: 1663479000000},
		{L: 8950, H: 9100, C: 8970, O: 9080, V: 2172300.00, T: 1663482600000},
		{L: 8950, H: 9120, C: 9070, O: 8980, V: 2417598.00, T: 1663486200000},
		{L: 8960, H: 9160, C: 9040, O: 8960, V: 1154404.00, T: 1663561800000},
		{L: 8980, H: 9070, C: 8990, O: 9040, V: 544846.00, T: 1663565400000},
		{L: 8930, H: 9020, C: 8960, O: 8990, V: 901781.00, T: 1663569000000},
		{L: 8910, H: 8970, C: 8920, O: 8970, V: 1021775.00, T: 1663572600000},
		{L: 8900, H: 9060, C: 8960, O: 8900, V: 1163197.00, T: 1663648200000},
		{L: 8860, H: 9000, C: 8870, O: 8950, V: 932584.00, T: 1663651800000},
		{L: 8820, H: 8880, C: 8830, O: 8880, V: 1458995.00, T: 1663655400000},
		{L: 8840, H: 8880, C: 8880, O: 8840, V: 462403.00, T: 1663659000000},
		{L: 8800, H: 8900, C: 8820, O: 8810, V: 787614.00, T: 1663734600000},
		{L: 8800, H: 8870, C: 8810, O: 8810, V: 677890.00, T: 1663738200000},
		{L: 8800, H: 8850, C: 8810, O: 8810, V: 1599221.00, T: 1663741800000},
		{L: 8800, H: 8950, C: 8930, O: 8810, V: 1817537.00, T: 1663745400000},
		{L: 8560, H: 8920, C: 8590, O: 8850, V: 1956165.00, T: 1663997400000},
		{L: 8530, H: 8630, C: 8570, O: 8580, V: 963507.00, T: 1664001000000},
		{L: 8560, H: 8650, C: 8560, O: 8570, V: 1202470.00, T: 1664004600000},
		{L: 8470, H: 8570, C: 8490, O: 8560, V: 1795527.00, T: 1664008200000},
		{L: 8360, H: 8900, C: 8860, O: 8360, V: 2232707.00, T: 1664170200000},
		{L: 8700, H: 8850, C: 8710, O: 8850, V: 804025.00, T: 1664173800000},
		{L: 8710, H: 8750, C: 8740, O: 8720, V: 490415.00, T: 1664177400000},
		{L: 8700, H: 8740, C: 8740, O: 8710, V: 354268.00, T: 1664181000000},
		{L: 8720, H: 9000, C: 8960, O: 8720, V: 1527278.00, T: 1664343000000},
		{L: 8910, H: 9060, C: 9000, O: 8960, V: 1085106.00, T: 1664346600000},
		{L: 8980, H: 9030, C: 9020, O: 9000, V: 473696.00, T: 1664350200000},
		{L: 8980, H: 9020, C: 8990, O: 9010, V: 699219.00, T: 1664353800000},
		{L: 8570, H: 8990, C: 8650, O: 8950, V: 2525245.00, T: 1664602200000},
		{L: 8540, H: 8750, C: 8580, O: 8670, V: 862745.00, T: 1664605800000},
		{L: 8400, H: 8590, C: 8450, O: 8570, V: 2648031.00, T: 1664609400000},
		{L: 8400, H: 8490, C: 8430, O: 8470, V: 880884.00, T: 1664613000000},
		{L: 8380, H: 8610, C: 8450, O: 8460, V: 2047200.00, T: 1664688600000},
		{L: 8410, H: 8490, C: 8460, O: 8470, V: 508220.00, T: 1664692200000},
		{L: 8430, H: 8480, C: 8450, O: 8460, V: 652416.00, T: 1664695800000},
		{L: 8400, H: 8460, C: 8440, O: 8440, V: 906352.00, T: 1664699400000},
		{L: 8420, H: 8600, C: 8480, O: 8490, V: 826543.00, T: 1664775000000},
		{L: 8430, H: 8570, C: 8450, O: 8480, V: 867974.00, T: 1664778600000},
		{L: 8430, H: 8480, C: 8450, O: 8450, V: 580218.00, T: 1664782200000},
		{L: 8440, H: 8540, C: 8500, O: 8450, V: 957990.00, T: 1664785800000},
		{L: 8400, H: 8560, C: 8500, O: 8400, V: 659743.00, T: 1664861400000},
		{L: 8450, H: 8500, C: 8480, O: 8500, V: 1096047.00, T: 1664865000000},
		{L: 8490, H: 8520, C: 8500, O: 8490, V: 390241.00, T: 1664868600000},
		{L: 8480, H: 8530, C: 8520, O: 8510, V: 1052288.00, T: 1664872200000},
		{L: 8400, H: 8650, C: 8620, O: 8400, V: 1420025.00, T: 1665207000000},
		{L: 8580, H: 8630, C: 8580, O: 8630, V: 664976.00, T: 1665210600000},
		{L: 8520, H: 8580, C: 8520, O: 8570, V: 811777.00, T: 1665214200000},
		{L: 8500, H: 8640, C: 8620, O: 8520, V: 1135548.00, T: 1665217800000},
		{L: 8360, H: 8620, C: 8400, O: 8460, V: 1757266.00, T: 1665293400000},
		{L: 8390, H: 8540, C: 8410, O: 8400, V: 2069165.00, T: 1665297000000},
		{L: 8370, H: 8460, C: 8450, O: 8420, V: 1283825.00, T: 1665300600000},
		{L: 8450, H: 8550, C: 8550, O: 8450, V: 1443903.00, T: 1665304200000},
		{L: 8460, H: 8590, C: 8500, O: 8500, V: 792803.00, T: 1665379800000},
		{L: 8420, H: 8520, C: 8470, O: 8510, V: 1158445.00, T: 1665383400000},
		{L: 8460, H: 8500, C: 8490, O: 8480, V: 806689.00, T: 1665387000000},
		{L: 8460, H: 8540, C: 8540, O: 8490, V: 1334939.00, T: 1665390600000},
		{L: 8490, H: 8900, C: 8800, O: 8490, V: 2464838.00, T: 1665466200000},
		{L: 8700, H: 8900, C: 8740, O: 8810, V: 2086815.00, T: 1665469800000},
		{L: 8730, H: 8820, C: 8750, O: 8740, V: 911276.00, T: 1665473400000},
		{L: 8740, H: 8890, C: 8890, O: 8760, V: 1329047.00, T: 1665477000000},
		{L: 8720, H: 8960, C: 8800, O: 8900, V: 1709787.00, T: 1665552600000},
		{L: 8790, H: 8850, C: 8840, O: 8800, V: 1069264.00, T: 1665556200000},
		{L: 8840, H: 8930, C: 8900, O: 8840, V: 1340324.00, T: 1665559800000},
		{L: 8880, H: 8950, C: 8950, O: 8900, V: 1694492.00, T: 1665563400000},
		{L: 8750, H: 8960, C: 8900, O: 8750, V: 1787429.00, T: 1665811800000},
		{L: 8780, H: 8910, C: 8780, O: 8910, V: 1494451.00, T: 1665815400000},
		{L: 8730, H: 8880, C: 8820, O: 8780, V: 924913.00, T: 1665819000000},
		{L: 8790, H: 8900, C: 8860, O: 8820, V: 1028523.00, T: 1665822600000},
		{L: 8640, H: 8880, C: 8650, O: 8870, V: 1524800.00, T: 1665898200000},
		{L: 8620, H: 8700, C: 8690, O: 8680, V: 1513406.00, T: 1665901800000},
		{L: 8570, H: 8690, C: 8600, O: 8640, V: 1445787.00, T: 1665905400000},
		{L: 8490, H: 8690, C: 8500, O: 8590, V: 1722207.00, T: 1665909000000},
		{L: 8500, H: 8650, C: 8500, O: 8500, V: 795140.00, T: 1665984600000},
		{L: 8510, H: 8580, C: 8550, O: 8510, V: 488220.00, T: 1665988200000},
		{L: 8530, H: 8610, C: 8600, O: 8550, V: 501777.00, T: 1665991800000},
		{L: 8540, H: 8700, C: 8690, O: 8590, V: 1029969.00, T: 1665995400000},
		{L: 8600, H: 8750, C: 8680, O: 8600, V: 911243.00, T: 1666071000000},
		{L: 8630, H: 8700, C: 8650, O: 8680, V: 757086.00, T: 1666074600000},
		{L: 8600, H: 8700, C: 8700, O: 8640, V: 425422.00, T: 1666078200000},
		{L: 8630, H: 8730, C: 8700, O: 8640, V: 660782.00, T: 1666081800000},
		{L: 8600, H: 8720, C: 8670, O: 8620, V: 469671.00, T: 1666157400000},
		{L: 8650, H: 8700, C: 8660, O: 8670, V: 546293.00, T: 1666161000000},
		{L: 8620, H: 8670, C: 8670, O: 8650, V: 804228.00, T: 1666164600000},
		{L: 8660, H: 8730, C: 8700, O: 8670, V: 1012749.00, T: 1666168200000},
		{L: 8580, H: 8840, C: 8720, O: 8580, V: 939987.00, T: 1666416600000},
		{L: 8700, H: 8830, C: 8720, O: 8720, V: 1003220.00, T: 1666420200000},
		{L: 8680, H: 8780, C: 8680, O: 8730, V: 767169.00, T: 1666423800000},
		{L: 8610, H: 8770, C: 8610, O: 8680, V: 472382.00, T: 1666427400000},
		{L: 8500, H: 8610, C: 8500, O: 8610, V: 785573.00, T: 1666503000000},
		{L: 8490, H: 8540, C: 8490, O: 8500, V: 1158605.00, T: 1666506600000},
		{L: 8450, H: 8520, C: 8520, O: 8490, V: 797369.00, T: 1666510200000},
		{L: 8450, H: 8530, C: 8450, O: 8510, V: 816922.00, T: 1666513800000},
		{L: 8490, H: 8580, C: 8520, O: 8580, V: 965156.00, T: 1666589400000},
		{L: 8380, H: 8520, C: 8380, O: 8520, V: 1881966.00, T: 1666593000000},
		{L: 8300, H: 8430, C: 8310, O: 8380, V: 1552299.00, T: 1666596600000},
		{L: 8290, H: 8360, C: 8340, O: 8310, V: 1326790.00, T: 1666600200000},
		{L: 8280, H: 8450, C: 8320, O: 8450, V: 1639128.00, T: 1666675800000},
		{L: 8160, H: 8320, C: 8180, O: 8310, V: 2019753.00, T: 1666679400000},
		{L: 8170, H: 8320, C: 8200, O: 8180, V: 1416805.00, T: 1666683000000},
		{L: 8150, H: 8230, C: 8190, O: 8200, V: 1475601.00, T: 1666686600000},
		{L: 8010, H: 8260, C: 8010, O: 8150, V: 1740023.00, T: 1666762200000},
		{L: 7920, H: 8040, C: 7940, O: 8010, V: 4274066.00, T: 1666765800000},
		{L: 7920, H: 7950, C: 7920, O: 7940, V: 3708253.00, T: 1666769400000},
		{L: 7920, H: 8040, C: 8020, O: 7920, V: 2467113.00, T: 1666773000000},
		{L: 7990, H: 8230, C: 7990, O: 8230, V: 3782570.00, T: 1667021400000},
		{L: 7900, H: 8010, C: 7900, O: 7990, V: 4110182.00, T: 1667025000000},
		{L: 7740, H: 7970, C: 7750, O: 7910, V: 1834825.00, T: 1667028600000},
		{L: 7690, H: 7800, C: 7760, O: 7760, V: 2606139.00, T: 1667032200000},
		{L: 7820, H: 7990, C: 7950, O: 7950, V: 4859163.00, T: 1667107800000},
		{L: 7920, H: 8030, C: 8010, O: 7960, V: 4220234.00, T: 1667111400000},
		{L: 8000, H: 8080, C: 8050, O: 8010, V: 1975520.00, T: 1667115000000},
		{L: 8030, H: 8210, C: 8200, O: 8040, V: 2780194.00, T: 1667118600000},
		{L: 8070, H: 8250, C: 8210, O: 8070, V: 1003287.00, T: 1667194200000},
		{L: 8100, H: 8210, C: 8110, O: 8210, V: 642473.00, T: 1667197800000},
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

	driver := ichimoku.NewIchimokuDriver()

	err := driver.MakeIchimokuInPast(&bar_h1, 135)
	if err != nil {
		fmt.Println("error :", err)
	}

	lines_result := make([]ichimoku.IchimokuStatus, 2)
	arr := driver.GetListDay()
	for i := len(arr) - 2; i > 0; i-- {
		lines_result[0] = arr[i]   //current
		lines_result[1] = arr[i+1] // previous

		a, e := driver.PreAnalyseIchimoku(lines_result)

		if e != nil {
			fmt.Println("err", e)
		}
		if a != nil {
			fmt.Printf("____ \r\n Find %v \r\n", a.Print())
			fmt.Print(".")
		}
	}

	fmt.Println(len(bar_h1))

}
