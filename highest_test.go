package shingo

import (
	"testing"
)

func TestHighest(t *testing.T) {
	hTests := []struct {
		title    string
		arg      IndicatorInputArg
		candles  []*Candlestick
		expected []float64
	}{
		{
			title: "Should get highest value for period up to that candlestick",
			arg: IndicatorInputArg{
				Type:   IndicatorTypeHighest,
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
				52.78,
				53.02,
				53.67,
				53.67,
				53.74,
				53.74,
				53.74,
				53.74,
				53.74,
				53.74,
				53.74,
				53.74,
				53.74,
				53.74,
				53.72,
				53.72,
				53.39,
				53.12,
				53.12,
				53.12,
				53.19,
				53.73,
				53.87,
				53.87,
				53.88,
				54.08,
				54.14,
				54.50,
				54.50,
				54.50,
				54.50,
			},
		},
	}

	for _, st := range hTests {
		cs, _ := NewCandlesticks(IntervalOneDay, 100)
		for _, c := range st.candles {
			cs.AppendCandlestick(c)
		}
		if err := cs.GenerateIndicator(IndicatorTypeHighest, st.arg); err != nil {
			t.Fatalf("Error appending stddev: %+v", err)
		}
		for i, e := range st.expected {
			v := cs.ItemAtIndex(i)
			high := v.GetHighest(st.arg.Period)
			if high == nil {
				continue
			}

			if !equalWithinPct(e, *high, 0.005) {
				t.Errorf("Expected value to be: %+v but got %+v", e, high)
			}
		}
	}
}
