package shingo

import (
	"testing"
)

// Raw data REQ/USD one minute interval
// from: 2018-04-29 13:32:00
// to: 2018-04-29 14:17:00
// 0.27190569, 0.26964718, 0.26981776, 0.2745, 0
// 0.27010248, 0.26894660, 0.26894943, 0.2732, 0
// 0.27046785, 0.26770696, 0.26929180, 0.2732, 0
// 0.26917804, 0.26672307, 0.26804067, 0.2725, 0
// 0.26941121, 0.26859641, 0.26922168, 0.2725, 0
// 0.26968698, 0.26941121, 0.26968698, 0.2725, 0
// 0.27000002, 0.26968414, 0.26981028, 0.2725, 0
// 0.26983872, 0.26941496, 0.26950313, 0.2725, 0
// 0.26958276, 0.26943487, 0.26958276, 0.2725, 0
// 0.26976193, 0.26957992, 0.26976193, 0.2725, 0
// 0.27004934, 0.26963858, 0.26987567, 0.2725, 0
// 0.27020565, 0.26988706, 0.27018000, 0.2725, 0
// 0.27039812, 0.26961120, 0.27039812, 0.2725, 0
// 0.26995636, 0.26954199, 0.26976658, 0.2724, 0
// 0.27054357, 0.27023204, 0.27032682, 0.2724, 0
// 0.27032682, 0.26987050, 0.26993324, 0.2724, 0
// 0.27022004, 0.26993324, 0.27022004, 0.2724, 0
// 0.26946286, 0.26744416, 0.26744416, 0.2714, 0
// 0.26744416, 0.26740177, 0.26740177, 0.27, 0
// 0.26995657, 0.26916830, 0.26992804, 0.27, 0
// 0.27003074, 0.26980250, 0.26997939, 0.27, 0
// 0.27040362, 0.26994844, 0.27026899, 0.2674, 1
// 0.27026614, 0.26994912, 0.27015761, 0.2676, 1
// 0.27014678, 0.27007764, 0.27008621, 0.2678, 1
// 0.27012619, 0.26989486, 0.26992913, 0.2679, 1
// 0.27023186, 0.26992627, 0.27011477, 0.268, 1
// 0.27018331, 0.27008906, 0.27015761, 0.2683, 1
// 0.27026649, 0.27021506, 0.27021506, 0.2685, 1
// 0.27056686, 0.27021506, 0.27053256, 0.2687, 1
// 0.27151356, 0.27030964, 0.27131854, 0.2691, 1
// 0.27158236, 0.27138737, 0.27157949, 0.2698, 1
// 0.27182906, 0.27148199, 0.27149921, 0.27, 1
// 0.27178310, 0.27159068, 0.27178310, 0.2701, 1
// 0.27223088, 0.27178023, 0.27222800, 0.2704, 1
// 0.27423544, 0.27213331, 0.27423544, 0.2711, 1
// 0.27521870, 0.27404592, 0.27461455, 0.2724, 1
// 0.27518390, 0.27428464, 0.27485330, 0.2725, 1
// 0.27518390, 0.27441917, 0.27470223, 0.2726, 1
// 0.27493740, 0.27481560, 0.27481850, 0.2728, 1
// 0.27512880, 0.27359810, 0.27380308, 0.2728, 1
// 0.27423210, 0.27209321, 0.27370323, 0.2728, 1
// 0.27341333, 0.27178580, 0.27311309, 0.2728, 1
// 0.27311309, 0.27310154, 0.27310443, 0.2728, 1
// 0.27293831, 0.27206385, 0.27237067, 0.2752, 0
// 0.27364682, 0.27246528, 0.27251136, 0.2752, 0

func TestAppendSuperTrend(t *testing.T) {

	superTrendTests := []struct {
		title    string
		arg      IndicatorInputArg
		candles  []*Candlestick
		expected []*SuperTrendDelta
	}{
		{
			title: "Provide Super Trend values from candles",
			arg: IndicatorInputArg{
				Period:     10,
				Multiplier: 3,
			},
			candles: []*Candlestick{
				&Candlestick{High: 0.27190569, Low: 0.26964718, Close: 0.26981776},
				&Candlestick{High: 0.27010248, Low: 0.26894660, Close: 0.26894943},
				&Candlestick{High: 0.27046785, Low: 0.26770696, Close: 0.26929180},
				&Candlestick{High: 0.26917804, Low: 0.26672307, Close: 0.26804067},
				&Candlestick{High: 0.26941121, Low: 0.26859641, Close: 0.26922168},
				&Candlestick{High: 0.26968698, Low: 0.26941121, Close: 0.26968698},
				&Candlestick{High: 0.27000002, Low: 0.26968414, Close: 0.26981028},
				&Candlestick{High: 0.26983872, Low: 0.26941496, Close: 0.26950313},
				&Candlestick{High: 0.26958276, Low: 0.26943487, Close: 0.26958276},
				&Candlestick{High: 0.26976193, Low: 0.26957992, Close: 0.26976193},
				&Candlestick{High: 0.27004934, Low: 0.26963858, Close: 0.26987567},
				&Candlestick{High: 0.27020565, Low: 0.26988706, Close: 0.27018000},
				&Candlestick{High: 0.27039812, Low: 0.26961120, Close: 0.27039812},
				&Candlestick{High: 0.26995636, Low: 0.26954199, Close: 0.26976658},
				&Candlestick{High: 0.27054357, Low: 0.27023204, Close: 0.27032682},
				&Candlestick{High: 0.27032682, Low: 0.26987050, Close: 0.26993324},
				&Candlestick{High: 0.27022004, Low: 0.26993324, Close: 0.27022004},
				&Candlestick{High: 0.26946286, Low: 0.26744416, Close: 0.26744416},
				&Candlestick{High: 0.26744416, Low: 0.26740177, Close: 0.26740177},
				&Candlestick{High: 0.26995657, Low: 0.26916830, Close: 0.26992804},
				&Candlestick{High: 0.27003074, Low: 0.26980250, Close: 0.26997939},
				&Candlestick{High: 0.27040362, Low: 0.26994844, Close: 0.27026899},
				&Candlestick{High: 0.27026614, Low: 0.26994912, Close: 0.27015761},
				&Candlestick{High: 0.27014678, Low: 0.27007764, Close: 0.27008621},
				&Candlestick{High: 0.27012619, Low: 0.26989486, Close: 0.26992913},
				&Candlestick{High: 0.27023186, Low: 0.26992627, Close: 0.27011477},
				&Candlestick{High: 0.27018331, Low: 0.27008906, Close: 0.27015761},
				&Candlestick{High: 0.27026649, Low: 0.27021506, Close: 0.27021506},
				&Candlestick{High: 0.27056686, Low: 0.27021506, Close: 0.27053256},
				&Candlestick{High: 0.27151356, Low: 0.27030964, Close: 0.27131854},
				&Candlestick{High: 0.27158236, Low: 0.27138737, Close: 0.27157949},
				&Candlestick{High: 0.27182906, Low: 0.27148199, Close: 0.27149921},
				&Candlestick{High: 0.27178310, Low: 0.27159068, Close: 0.27178310},
				&Candlestick{High: 0.27223088, Low: 0.27178023, Close: 0.27222800},
				&Candlestick{High: 0.27423544, Low: 0.27213331, Close: 0.27423544},
				&Candlestick{High: 0.27521870, Low: 0.27404592, Close: 0.27461455},
				&Candlestick{High: 0.27518390, Low: 0.27428464, Close: 0.27485330},
				&Candlestick{High: 0.27518390, Low: 0.27441917, Close: 0.27470223},
				&Candlestick{High: 0.27493740, Low: 0.27481560, Close: 0.27481850},
				&Candlestick{High: 0.27512880, Low: 0.27359810, Close: 0.27380308},
				&Candlestick{High: 0.27423210, Low: 0.27209321, Close: 0.27370323},
				&Candlestick{High: 0.27341333, Low: 0.27178580, Close: 0.27311309},
				&Candlestick{High: 0.27311309, Low: 0.27310154, Close: 0.27310443},
				&Candlestick{High: 0.27293831, Low: 0.27206385, Close: 0.27237067},
				&Candlestick{High: 0.27364682, Low: 0.27246528, Close: 0.27251136},
			},

			expected: []*SuperTrendDelta{
				nil,
				nil,
				nil,
				nil,
				nil,
				nil,
				nil,
				nil,
				nil,
				&SuperTrendDelta{Shortband: 0.2725},
				&SuperTrendDelta{Shortband: 0.2725},
				&SuperTrendDelta{Shortband: 0.2725},
				&SuperTrendDelta{Shortband: 0.2725},
				&SuperTrendDelta{Shortband: 0.2724},
				&SuperTrendDelta{Shortband: 0.2724},
				&SuperTrendDelta{Shortband: 0.2724},
				&SuperTrendDelta{Shortband: 0.2724},
				&SuperTrendDelta{Shortband: 0.2714},
				&SuperTrendDelta{Shortband: 0.27},
				&SuperTrendDelta{Shortband: 0.27},
				&SuperTrendDelta{Shortband: 0.27},
				&SuperTrendDelta{Longband: 0.2674, Trend: Bull},
				&SuperTrendDelta{Longband: 0.2676, Trend: Bull},
				&SuperTrendDelta{Longband: 0.2678, Trend: Bull},
				&SuperTrendDelta{Longband: 0.2679, Trend: Bull},
				&SuperTrendDelta{Longband: 0.268, Trend: Bull},
				&SuperTrendDelta{Longband: 0.2683, Trend: Bull},
				&SuperTrendDelta{Longband: 0.2685, Trend: Bull},
				&SuperTrendDelta{Longband: 0.2687, Trend: Bull},
				&SuperTrendDelta{Longband: 0.2691, Trend: Bull},
				&SuperTrendDelta{Longband: 0.2698, Trend: Bull},
				&SuperTrendDelta{Longband: 0.27, Trend: Bull},
				&SuperTrendDelta{Longband: 0.2701, Trend: Bull},
				&SuperTrendDelta{Longband: 0.2704, Trend: Bull},
				&SuperTrendDelta{Longband: 0.2711, Trend: Bull},
				&SuperTrendDelta{Longband: 0.2724, Trend: Bull},
				&SuperTrendDelta{Longband: 0.2725, Trend: Bull},
				&SuperTrendDelta{Longband: 0.2726, Trend: Bull},
				&SuperTrendDelta{Longband: 0.2728, Trend: Bull},
				&SuperTrendDelta{Longband: 0.2728, Trend: Bull},
				&SuperTrendDelta{Longband: 0.2728, Trend: Bull},
				&SuperTrendDelta{Longband: 0.2728, Trend: Bull},
				&SuperTrendDelta{Longband: 0.2728, Trend: Bull},
				&SuperTrendDelta{Shortband: 0.2752, Trend: Bear},
				&SuperTrendDelta{Shortband: 0.2752, Trend: Bear},
			},
		},
	}

	for _, v := range superTrendTests {
		cs, _ := NewCandlesticks(IntervalOneDay, 100)
		for _, c := range v.candles {
			cs.AppendCandlestick(c)
		}

		if err := cs.GenerateIndicator(IndicatorTypeSuperTrend, v.arg); err != nil {
			t.Errorf("Expected ok but got error %+v appending Super Trend to period: 10, args: %+v for %s", err, v.arg, v.title)
		}

		for i := range v.candles {
			c := cs.ItemAtIndex(i)
			if v.expected[i] == nil {
				if c.Indicators != nil && c.Indicators.SuperTrends[10][3] != nil {
					t.Errorf("Expected nil but got %+v for test: %s, index: %d", c.Indicators, v.title, i)
				}
				continue
			} else if v.expected[i] != nil && (c.Indicators == nil || c.Indicators.SuperTrends == nil || c.Indicators.SuperTrends[10][3] == nil) {
				t.Errorf("Expected non nil but got nil %+v for test: %s, index: %d", c.Indicators, v.title, i)
				continue
			}
			st := c.Indicators.SuperTrends[10][3]
			e := v.expected[i]
			// i > 20 because trend is eventual consistency
			// and it will be inaccurate towards the beginning
			if i > 20 && e.Trend != st.Trend {
				t.Errorf("Expected super trend to be %+v but got %+v for test: %s, index: %d", e.Trend, st.Trend, v.title, i)
			}
			if e.Shortband > 0 {
				if !almostEqual(st.Shortband, e.Shortband, 0.001) {
					t.Errorf("Expected Shortband to be %+v but got %+v for test: %s, index: %d",
						e.Shortband,
						st.Shortband,
						v.title,
						i)
				}
			} else {
				if !almostEqual(st.Longband, e.Longband, 0.001) {
					t.Errorf("Expected Longband to be %+v but got %+v for test: %s, index: %d",
						e.Longband,
						st.Longband,
						v.title,
						i)
				}
			}
		}
	}
}
