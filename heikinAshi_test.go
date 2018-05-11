package shingo

import (
	"fmt"
	"testing"
)

func TestHeikinAshi(t *testing.T) {
	// Raw data
	// 1109.75	1112.25	1103.50	1107.25	1109.75	1112.25	1103.50	1108.19
	// 1104.25	1107.50	1100.00	1103.50	1108.97	1108.97	1100.00	1103.81
	// 1103.25	1108.00	1101.50	1108.00	1106.39	1108.00	1101.50	1105.19
	// 1108.00	1109.00	1104.00	1108.75	1105.79	1109.00	1104.00	1107.44
	// 1107.75	1109.50	1100.75	1105.00	1106.61	1109.50	1100.75	1105.75
	// 1105.00	1107.00	1101.25	1104.00	1106.18	1107.00	1101.25	1104.31
	// 1104.00	1108.50	1103.00	1108.00	1105.25	1108.50	1103.00	1105.88
	// 1100.25	1100.75	1086.50	1089.00	1105.56	1105.56	1086.50	1094.13
	// 1088.75	1092.00	1088.00	1090.75	1099.84	1099.84	1088.00	1089.88
	// 1090.75	1095.25	1088.25	1093.25	1094.86	1095.25	1088.25	1091.88
	// 1087.50	1092.75	1085.25	1086.00	1093.37	1093.37	1085.25	1087.88
	// 1086.25	1089.50	1085.50	1089.00	1090.62	1090.62	1085.50	1087.56
	// 1089.00	1092.75	1088.25	1090.75	1089.09	1092.75	1088.25	1090.19
	// 1102.25	1111.50	1102.25	1107.75	1089.64	1111.50	1089.64	1105.94
	// 1107.50	1108.25	1103.75	1104.25	1097.79	1108.25	1097.79	1105.94
	// 1104.25	1105.50	1101.25	1104.00	1101.86	1105.50	1101.25	1103.75
	// 1105.50	1105.50	1095.00	1100.50	1102.81	1105.50	1095.00	1101.63
	// 1100.50	1102.75	1099.75	1102.25	1102.22	1102.75	1099.75	1101.31
	// 1102.00	1106.75	1100.75	1104.50	1101.76	1106.75	1100.75	1103.50
	// 1106.75	1109.75	1103.25	1109.75	1102.63	1109.75	1102.63	1107.38
	// 1109.75	1110.00	1106.50	1108.25	1105.00	1110.00	1105.00	1108.63
	// 1108.00	1110.25	1107.00	1109.00	1106.81	1110.25	1106.81	1108.56
	// 1094.50	1095.00	1083.00	1084.75	1107.69	1107.69	1083.00	1089.31
	// 1078.75	1098.25	1077.75	1094.00	1098.50	1098.50	1077.75	1087.19
	// 1094.00	1097.25	1088.25	1089.50	1092.84	1097.25	1088.25	1092.25
	// 1090.00	1096.50	1085.00	1087.00	1092.55	1096.50	1085.00	1089.63
	// 1086.75	1089.75	1085.00	1088.75	1091.09	1091.09	1085.00	1087.56
	// 1088.50	1096.50	1087.75	1094.75	1089.32	1096.50	1087.75	1091.88
	// 1104.25	1108.75	1102.25	1107.50	1090.60	1108.75	1090.60	1105.69
	// 1107.75	1111.75	1107.00	1111.25	1098.14	1111.75	1098.14	1109.44
	// 1111.25	1111.50	1106.50	1107.50	1103.79	1111.50	1103.79	1109.19
	// 1108.25	1115.50	1107.25	1109.50	1106.49	1115.50	1106.49	1110.13
	// 1109.25	1109.75	1104.25	1107.50	1108.31	1109.75	1104.25	1107.69
	// 1107.75	1109.75	1105.00	1108.00	1108.00	1109.75	1105.00	1107.63
	// 1110.50	1117.00	1105.25	1107.50	1107.81	1117.00	1105.25	1110.06
	// 1107.75	1110.75	1107.25	1110.25	1108.94	1110.75	1107.25	1109.00
	// 1110.25	1110.75	1098.00	1099.25	1108.97	1110.75	1098.00	1104.56
	// 1113.50	1119.00	1103.75	1104.75	1106.77	1119.00	1103.75	1110.25
	// 1104.75	1105.75	1095.25	1101.25	1108.51	1108.51	1095.25	1101.75
	// 1101.25	1106.75	1095.50	1105.00	1105.13	1106.75	1095.50	1102.13

	haTests := []struct {
		title    string
		arg      IndicatorInputArg
		candles  []*Candlestick
		expected []*HeikinAshiDelta
	}{
		{
			arg: IndicatorInputArg{
				Limit: 0,
			},
			candles: []*Candlestick{
				&Candlestick{Open: 1109.75, High: 1112.25, Low: 1103.50, Close: 1107.25},
				&Candlestick{Open: 1104.25, High: 1107.50, Low: 1100.00, Close: 1103.50},
				&Candlestick{Open: 1103.25, High: 1108.00, Low: 1101.50, Close: 1108.00},
				&Candlestick{Open: 1108.00, High: 1109.00, Low: 1104.00, Close: 1108.75},
				&Candlestick{Open: 1107.75, High: 1109.50, Low: 1100.75, Close: 1105.00},
				&Candlestick{Open: 1105.00, High: 1107.00, Low: 1101.25, Close: 1104.00},
				&Candlestick{Open: 1104.00, High: 1108.50, Low: 1103.00, Close: 1108.00},
				&Candlestick{Open: 1100.25, High: 1100.75, Low: 1086.50, Close: 1089.00},
				&Candlestick{Open: 1088.75, High: 1092.00, Low: 1088.00, Close: 1090.75},
				&Candlestick{Open: 1090.75, High: 1095.25, Low: 1088.25, Close: 1093.25},
				&Candlestick{Open: 1087.50, High: 1092.75, Low: 1085.25, Close: 1086.00},
				&Candlestick{Open: 1086.25, High: 1089.50, Low: 1085.50, Close: 1089.00},
				&Candlestick{Open: 1089.00, High: 1092.75, Low: 1088.25, Close: 1090.75},
				&Candlestick{Open: 1102.25, High: 1111.50, Low: 1102.25, Close: 1107.75},
				&Candlestick{Open: 1107.50, High: 1108.25, Low: 1103.75, Close: 1104.25},
				&Candlestick{Open: 1104.25, High: 1105.50, Low: 1101.25, Close: 1104.00},
				&Candlestick{Open: 1105.50, High: 1105.50, Low: 1095.00, Close: 1100.50},
				&Candlestick{Open: 1100.50, High: 1102.75, Low: 1099.75, Close: 1102.25},
				&Candlestick{Open: 1102.00, High: 1106.75, Low: 1100.75, Close: 1104.50},
				&Candlestick{Open: 1106.75, High: 1109.75, Low: 1103.25, Close: 1109.75},
				&Candlestick{Open: 1109.75, High: 1110.00, Low: 1106.50, Close: 1108.25},
				&Candlestick{Open: 1108.00, High: 1110.25, Low: 1107.00, Close: 1109.00},
				&Candlestick{Open: 1094.50, High: 1095.00, Low: 1083.00, Close: 1084.75},
				&Candlestick{Open: 1078.75, High: 1098.25, Low: 1077.75, Close: 1094.00},
				&Candlestick{Open: 1094.00, High: 1097.25, Low: 1088.25, Close: 1089.50},
				&Candlestick{Open: 1090.00, High: 1096.50, Low: 1085.00, Close: 1087.00},
				&Candlestick{Open: 1086.75, High: 1089.75, Low: 1085.00, Close: 1088.75},
				&Candlestick{Open: 1088.50, High: 1096.50, Low: 1087.75, Close: 1094.75},
				&Candlestick{Open: 1104.25, High: 1108.75, Low: 1102.25, Close: 1107.50},
				&Candlestick{Open: 1107.75, High: 1111.75, Low: 1107.00, Close: 1111.25},
				&Candlestick{Open: 1111.25, High: 1111.50, Low: 1106.50, Close: 1107.50},
				&Candlestick{Open: 1108.25, High: 1115.50, Low: 1107.25, Close: 1109.50},
				&Candlestick{Open: 1109.25, High: 1109.75, Low: 1104.25, Close: 1107.50},
				&Candlestick{Open: 1107.75, High: 1109.75, Low: 1105.00, Close: 1108.00},
				&Candlestick{Open: 1110.50, High: 1117.00, Low: 1105.25, Close: 1107.50},
				&Candlestick{Open: 1107.75, High: 1110.75, Low: 1107.25, Close: 1110.25},
				&Candlestick{Open: 1110.25, High: 1110.75, Low: 1098.00, Close: 1099.25},
				&Candlestick{Open: 1113.50, High: 1119.00, Low: 1103.75, Close: 1104.75},
				&Candlestick{Open: 1104.75, High: 1105.75, Low: 1095.25, Close: 1101.25},
				&Candlestick{Open: 1101.25, High: 1106.75, Low: 1095.50, Close: 1105.00},
			},
			expected: []*HeikinAshiDelta{
				&HeikinAshiDelta{Open: 1109.75, High: 1112.25, Low: 1103.50, Close: 1108.19},
				&HeikinAshiDelta{Open: 1108.97, High: 1108.97, Low: 1100.00, Close: 1103.81},
				&HeikinAshiDelta{Open: 1106.39, High: 1108.00, Low: 1101.50, Close: 1105.19},
				&HeikinAshiDelta{Open: 1105.79, High: 1109.00, Low: 1104.00, Close: 1107.44},
				&HeikinAshiDelta{Open: 1106.61, High: 1109.50, Low: 1100.75, Close: 1105.75},
				&HeikinAshiDelta{Open: 1106.18, High: 1107.00, Low: 1101.25, Close: 1104.31},
				&HeikinAshiDelta{Open: 1105.25, High: 1108.50, Low: 1103.00, Close: 1105.88},
				&HeikinAshiDelta{Open: 1105.56, High: 1105.56, Low: 1086.50, Close: 1094.13},
				&HeikinAshiDelta{Open: 1099.84, High: 1099.84, Low: 1088.00, Close: 1089.88},
				&HeikinAshiDelta{Open: 1094.86, High: 1095.25, Low: 1088.25, Close: 1091.88},
				&HeikinAshiDelta{Open: 1093.37, High: 1093.37, Low: 1085.25, Close: 1087.88},
				&HeikinAshiDelta{Open: 1090.62, High: 1090.62, Low: 1085.50, Close: 1087.56},
				&HeikinAshiDelta{Open: 1089.09, High: 1092.75, Low: 1088.25, Close: 1090.19},
				&HeikinAshiDelta{Open: 1089.64, High: 1111.50, Low: 1089.64, Close: 1105.94},
				&HeikinAshiDelta{Open: 1097.79, High: 1108.25, Low: 1097.79, Close: 1105.94},
				&HeikinAshiDelta{Open: 1101.86, High: 1105.50, Low: 1101.25, Close: 1103.75},
				&HeikinAshiDelta{Open: 1102.81, High: 1105.50, Low: 1095.00, Close: 1101.63},
				&HeikinAshiDelta{Open: 1102.22, High: 1102.75, Low: 1099.75, Close: 1101.31},
				&HeikinAshiDelta{Open: 1101.76, High: 1106.75, Low: 1100.75, Close: 1103.50},
				&HeikinAshiDelta{Open: 1102.63, High: 1109.75, Low: 1102.63, Close: 1107.38},
				&HeikinAshiDelta{Open: 1105.00, High: 1110.00, Low: 1105.00, Close: 1108.63},
				&HeikinAshiDelta{Open: 1106.81, High: 1110.25, Low: 1106.81, Close: 1108.56},
				&HeikinAshiDelta{Open: 1107.69, High: 1107.69, Low: 1083.00, Close: 1089.31},
				&HeikinAshiDelta{Open: 1098.50, High: 1098.50, Low: 1077.75, Close: 1087.19},
				&HeikinAshiDelta{Open: 1092.84, High: 1097.25, Low: 1088.25, Close: 1092.25},
				&HeikinAshiDelta{Open: 1092.55, High: 1096.50, Low: 1085.00, Close: 1089.63},
				&HeikinAshiDelta{Open: 1091.09, High: 1091.09, Low: 1085.00, Close: 1087.56},
				&HeikinAshiDelta{Open: 1089.32, High: 1096.50, Low: 1087.75, Close: 1091.88},
				&HeikinAshiDelta{Open: 1090.60, High: 1108.75, Low: 1090.60, Close: 1105.69},
				&HeikinAshiDelta{Open: 1098.14, High: 1111.75, Low: 1098.14, Close: 1109.44},
				&HeikinAshiDelta{Open: 1103.79, High: 1111.50, Low: 1103.79, Close: 1109.19},
				&HeikinAshiDelta{Open: 1106.49, High: 1115.50, Low: 1106.49, Close: 1110.13},
				&HeikinAshiDelta{Open: 1108.31, High: 1109.75, Low: 1104.25, Close: 1107.69},
				&HeikinAshiDelta{Open: 1108.00, High: 1109.75, Low: 1105.00, Close: 1107.63},
				&HeikinAshiDelta{Open: 1107.81, High: 1117.00, Low: 1105.25, Close: 1110.06},
				&HeikinAshiDelta{Open: 1108.94, High: 1110.75, Low: 1107.25, Close: 1109.00},
				&HeikinAshiDelta{Open: 1108.97, High: 1110.75, Low: 1098.00, Close: 1104.56},
				&HeikinAshiDelta{Open: 1106.77, High: 1119.00, Low: 1103.75, Close: 1110.25},
				&HeikinAshiDelta{Open: 1108.51, High: 1108.51, Low: 1095.25, Close: 1101.75},
				&HeikinAshiDelta{Open: 1105.13, High: 1106.75, Low: 1095.50, Close: 1102.13},
			},
		},
	}

	for _, v := range haTests {
		cs, _ := NewCandlesticks(IntervalOneDay, 100)
		for _, c := range v.candles {
			cs.AppendCandlestick(c)
		}

		if err := cs.GenerateIndicator(IndicatorTypeHeikinAshi, v.arg); err != nil {
			t.Errorf("Expected ok but got error: %+v, args: %+v, title: %s", err, v.arg, v.title)
		}

		var str string
		for i := range v.candles {
			c := cs.ItemAtIndex(i)
			ha := c.GetHeikinAshi()
			str += fmt.Sprintf("%v,", c.Close)
			if v.expected[i] == nil {
				if c.Indicators != nil && ha != nil {
					t.Errorf("Expected nil but got %+v for test %s index %d", c.Indicators, v.title, i)
				}
				continue
			} else if v.expected[i] != nil && ha == nil {
				t.Errorf("Expected non nil but got nil %+v for test %s index %d", ha, v.title, i)
				continue
			}
			if !almostEqual(ha.Open, v.expected[i].Open, 0.01) {
				t.Errorf("Expected Open to be: %+v, but got: %+v, test: %s, index %d",
					v.expected[i].Open,
					ha.Open,
					v.title,
					i)
			}
			if !almostEqual(ha.High, v.expected[i].High, 0.01) {
				t.Errorf("Expected High to be: %+v, but got: %+v, test: %s, index %d",
					v.expected[i].High,
					ha.High,
					v.title,
					i)
			}
			if !almostEqual(ha.Low, v.expected[i].Low, 0.01) {
				t.Errorf("Expected Low to be: %+v, but got: %+v, test: %s, index %d",
					v.expected[i].Low,
					ha.Low,
					v.title,
					i)
			}
			if !almostEqual(ha.Close, v.expected[i].Close, 0.01) {
				t.Errorf("Expected Close to be: %+v, but got: %+v, test: %s, index %d",
					v.expected[i].Close,
					ha.Close,
					v.title,
					i)
			}
		}
	}
}
