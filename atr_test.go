package shingo

import (
	"testing"
)

func TestAppendATR(t *testing.T) {

	atrTests := []struct {
		title    string
		arg      IndicatorInputArg
		candles  []*Candlestick
		expected []*ATRDelta
	}{
		{
			title: "Provide ATR values from candles",
			arg: IndicatorInputArg{
				Limit:  40,
				Period: 14,
			},
			candles: []*Candlestick{
				&Candlestick{High: 48.70, Low: 47.79, Close: 48.16},
				&Candlestick{High: 48.72, Low: 48.14, Close: 48.61},
				&Candlestick{High: 48.90, Low: 48.39, Close: 48.75},
				&Candlestick{High: 48.87, Low: 48.37, Close: 48.63},
				&Candlestick{High: 48.82, Low: 48.24, Close: 48.74},
				&Candlestick{High: 49.05, Low: 48.64, Close: 49.03},
				&Candlestick{High: 49.20, Low: 48.94, Close: 49.07},
				&Candlestick{High: 49.35, Low: 48.86, Close: 49.32},
				&Candlestick{High: 49.92, Low: 49.50, Close: 49.91},
				&Candlestick{High: 50.19, Low: 49.87, Close: 50.13},
				&Candlestick{High: 50.12, Low: 49.20, Close: 49.53},
				&Candlestick{High: 49.66, Low: 48.90, Close: 49.50},
				&Candlestick{High: 49.88, Low: 49.43, Close: 49.75},
				&Candlestick{High: 50.19, Low: 49.73, Close: 50.03},
				&Candlestick{High: 50.36, Low: 49.26, Close: 50.31},
				&Candlestick{High: 50.57, Low: 50.09, Close: 50.52},
				&Candlestick{High: 50.65, Low: 50.30, Close: 50.41},
				&Candlestick{High: 50.43, Low: 49.21, Close: 49.34},
				&Candlestick{High: 49.63, Low: 48.98, Close: 49.37},
				&Candlestick{High: 50.33, Low: 49.61, Close: 50.23},
				&Candlestick{High: 50.29, Low: 49.20, Close: 49.24},
				&Candlestick{High: 50.17, Low: 49.43, Close: 49.93},
				&Candlestick{High: 49.32, Low: 48.08, Close: 48.43},
				&Candlestick{High: 48.50, Low: 47.64, Close: 48.18},
				&Candlestick{High: 48.32, Low: 41.55, Close: 46.57},
				&Candlestick{High: 46.80, Low: 44.28, Close: 45.41},
				&Candlestick{High: 47.80, Low: 47.31, Close: 47.77},
				&Candlestick{High: 48.39, Low: 47.20, Close: 47.72},
				&Candlestick{High: 48.66, Low: 47.90, Close: 48.62},
				&Candlestick{High: 48.79, Low: 47.73, Close: 47.85},
			},

			expected: []*ATRDelta{
				nil, nil, nil, nil, nil,
				nil, nil, nil, nil, nil,
				nil, nil, nil,
				&ATRDelta{Value: 0.56},
				&ATRDelta{Value: 0.59, Change: 0.07014157014},
				&ATRDelta{Value: 0.59, Change: -0.01370157203},
				&ATRDelta{Value: 0.57, Change: -0.02875122091},
				&ATRDelta{Value: 0.62, Change: 0.08173615175},
				&ATRDelta{Value: 0.62, Change: 0.004009561668},
				&ATRDelta{Value: 0.64, Change: 0.03954280166},
				&ATRDelta{Value: 0.67, Change: 0.04977735326},
				&ATRDelta{Value: 0.69, Change: 0.02734687294},
				&ATRDelta{Value: 0.78, Change: 0.1193171375},
				&ATRDelta{Value: 0.78, Change: 0.007790244433},
				&ATRDelta{Value: 1.21, Change: 0.5473778742},
				&ATRDelta{Value: 1.30, Change: 0.07723166242},
				&ATRDelta{Value: 1.38, Change: 0.05962599367},
				&ATRDelta{Value: 1.37, Change: -0.009847305764},
				&ATRDelta{Value: 1.34, Change: -0.02230076997},
				&ATRDelta{Value: 1.32, Change: -0.01477085132},
			},
		},
	}

	for _, v := range atrTests {
		cs, _ := NewCandlesticks(IntervalOneDay, 100)
		for _, c := range v.candles {
			cs.AppendCandlestick(c)
		}

		if err := cs.GenerateIndicator(IndicatorTypeATR, v.arg); err != nil {
			t.Errorf("Expected ok but got error %+v appending ATR to period: 14, args: %+v for %s", err, v.arg, v.title)
		}

		for i := range v.candles {
			c := cs.ItemAtIndex(i)
			if v.expected[i] == nil {
				if c.Indicators != nil && c.Indicators.ATRs[14] != nil {
					t.Errorf("Expected nil but got %+v for test: %s, index: %d", c.Indicators, v.title, i)
				}
				continue
			} else if v.expected[i] != nil && (c.Indicators == nil || c.Indicators.ATRs == nil || c.Indicators.ATRs[14] == nil) {
				t.Errorf("Expected non nil but got nil %+v for test: %s, index: %d", c.Indicators, v.title, i)
				continue
			}
			atrs := c.Indicators.ATRs[14]
			if !almostEqual(atrs.Value, v.expected[i].Value, 0.01) {
				t.Errorf("Expected ATR value to be %+v but got %+v for test: %s, index: %d",
					v.expected[i].Value,
					atrs.Value,
					v.title,
					i)
			}
			if !almostEqual(atrs.Change, v.expected[i].Change, 0.01) {
				t.Errorf("Expected ATR change to be %+v but got %+v for test: %s, index: %d",
					v.expected[i].Change,
					atrs.Change,
					v.title,
					i)
			}
		}
	}
}
