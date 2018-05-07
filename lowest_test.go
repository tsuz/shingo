package shingo

import "testing"

// 52.22			52.22
// 52.78			52.22
// 53.02			52.22
// 53.67			52.22
// 53.67			52.22
// 53.74			52.22
// 53.45			52.22
// 53.72			52.22
// 53.39			52.22
// 52.51			52.22
// 52.32			52.32
// 51.45			51.45
// 51.60			51.45
// 52.43			51.45
// 52.47			51.45
// 52.91			51.45
// 52.07			51.45
// 53.12			51.45
// 52.77			51.45
// 52.73			51.45
// 52.09			51.45
// 53.19			51.60
// 53.73			52.07
// 53.87			52.07
// 53.85			52.07
// 53.88			52.07
// 54.08			52.09
// 54.14			52.09
// 54.50			52.09
// 54.30			52.09
// 54.40			53.19
// 54.16			53.73

func TestLowest(t *testing.T) {
	lTests := []struct {
		title    string
		arg      IndicatorInputArg
		candles  []*Candlestick
		expected []float64
	}{
		{
			title: "Should get lowest value for period up to that candlestick",
			arg: IndicatorInputArg{
				Type:   IndicatorTypeLowest,
				Period: 10,
			},
			candles: []*Candlestick{
				&Candlestick{Close: 52.22},
				&Candlestick{Close: 52.78},
				&Candlestick{Close: 53.02},
				&Candlestick{Close: 53.67},
				&Candlestick{Close: 53.67},
				&Candlestick{Close: 53.74},
				&Candlestick{Close: 53.45},
				&Candlestick{Close: 53.72},
				&Candlestick{Close: 53.39},
				&Candlestick{Close: 52.51},
				&Candlestick{Close: 52.32},
				&Candlestick{Close: 51.45},
				&Candlestick{Close: 51.60},
				&Candlestick{Close: 52.43},
				&Candlestick{Close: 52.47},
				&Candlestick{Close: 52.91},
				&Candlestick{Close: 52.07},
				&Candlestick{Close: 53.12},
				&Candlestick{Close: 52.77},
				&Candlestick{Close: 52.73},
				&Candlestick{Close: 52.09},
				&Candlestick{Close: 53.19},
				&Candlestick{Close: 53.73},
				&Candlestick{Close: 53.87},
				&Candlestick{Close: 53.85},
				&Candlestick{Close: 53.88},
				&Candlestick{Close: 54.08},
				&Candlestick{Close: 54.14},
				&Candlestick{Close: 54.50},
				&Candlestick{Close: 54.30},
				&Candlestick{Close: 54.40},
				&Candlestick{Close: 54.16},
			},
			expected: []float64{
				52.22,
				52.22,
				52.22,
				52.22,
				52.22,
				52.22,
				52.22,
				52.22,
				52.22,
				52.22,
				52.32,
				51.45,
				51.45,
				51.45,
				51.45,
				51.45,
				51.45,
				51.45,
				51.45,
				51.45,
				51.45,
				51.60,
				52.07,
				52.07,
				52.07,
				52.07,
				52.09,
				52.09,
				52.09,
				52.09,
				53.19,
				53.73,
			},
		},
	}

	for _, st := range lTests {
		cs, _ := NewCandlesticks(IntervalOneDay, 100)
		for _, c := range st.candles {
			cs.AppendCandlestick(c)
		}
		if err := cs.GenerateIndicator(IndicatorTypeLowest, st.arg); err != nil {
			t.Fatalf("Error appending stddev: %+v", err)
		}
		for i, e := range st.expected {
			v := cs.ItemAtIndex(i)
			indicator := v.Indicators
			if indicator == nil {
				continue
			}
			low := v.Indicators.Get(st.arg)
			if low == nil {
				t.Fatalf("Expected lowest to be non nil")
			}
			if val, ok := low.(float64); ok {
				if !equalWithinPct(e, val, 0.005) {
					t.Errorf("Expected value to be: %+v but got %+v at idx: %d", e, val, i)
				}
			} else {
				t.Fatalf("Expected type float64 but got type: %T", low)
			}
		}
	}
}
