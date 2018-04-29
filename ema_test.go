package shingo

import (
	"testing"
)

func TestAppendEMA(t *testing.T) {

	emaTests := []struct {
		title    string
		args     IndicatorInputArg
		candles  []*Candlestick
		expected []*EMADelta
	}{
		{
			title: "It should not have any computed values",
			args: IndicatorInputArg{
				Limit:  1,
				Period: 4,
			},
			candles: []*Candlestick{
				&Candlestick{Close: 3.4},
				&Candlestick{Close: 2.1},
				&Candlestick{Close: 9},
			},
			expected: []*EMADelta{
				nil,
				nil,
				nil,
			},
		},
		{
			title: "It should have as many computed values (one) as possible",
			args: IndicatorInputArg{
				Limit:  0,
				Period: 4,
			},
			candles: []*Candlestick{
				&Candlestick{Close: 3.4},
				&Candlestick{Close: 2.1},
				&Candlestick{Close: 9},
				&Candlestick{Close: 9},
			},
			expected: []*EMADelta{
				nil,
				nil,
				nil,
				&EMADelta{Value: 5.875, Change: 0},
			},
		},
		{
			title: "It should have only one as limit is 1",
			args: IndicatorInputArg{
				Limit:  1,
				Period: 4,
			},
			candles: []*Candlestick{
				&Candlestick{Close: 3.1},
				&Candlestick{Close: 2.8},
				&Candlestick{Close: 3.4},
				&Candlestick{Close: 2.1},
				&Candlestick{Close: 9},
			},
			expected: []*EMADelta{
				nil,
				nil,
				nil,
				nil,
				&EMADelta{Value: 5.31, Change: 0.86315789},
			},
		},
		{
			title: "It should have two as limit is 2",
			args: IndicatorInputArg{
				Limit:  2,
				Period: 4,
			},
			candles: []*Candlestick{
				&Candlestick{Close: 3.1},
				&Candlestick{Close: 2.8},
				&Candlestick{Close: 3.4},
				&Candlestick{Close: 2.1},
				&Candlestick{Close: 9},
			},
			expected: []*EMADelta{
				nil,
				nil,
				nil,
				&EMADelta{Value: 2.85, Change: 0},
				&EMADelta{Value: 5.31, Change: 0.86315789},
			},
		},
		{
			title: "It should have only two as limit is 2",
			args: IndicatorInputArg{
				Limit:  2,
				Period: 4,
			},
			candles: []*Candlestick{
				&Candlestick{Close: 1.1},
				&Candlestick{Close: 3.1},
				&Candlestick{Close: 2.8},
				&Candlestick{Close: 3.4},
				&Candlestick{Close: 2.1},
				&Candlestick{Close: 9},
			},
			expected: []*EMADelta{
				nil,
				nil,
				nil,
				nil,
				&EMADelta{Value: 2.4, Change: -0.07692308},
				&EMADelta{Value: 5.04, Change: 1.1},
			},
		},
		{
			title: "It should have only two out of 4 possible as limit is 2",
			args: IndicatorInputArg{
				Limit:  2,
				Period: 4,
			},
			candles: []*Candlestick{
				&Candlestick{Close: 1.1},
				&Candlestick{Close: 1.1},
				&Candlestick{Close: 3.1},
				&Candlestick{Close: 2.8},
				&Candlestick{Close: 3.4},
				&Candlestick{Close: 2.1},
				&Candlestick{Close: 9},
			},
			expected: []*EMADelta{
				nil,
				nil,
				nil,
				nil,
				nil,
				&EMADelta{Value: 2.385, Change: -0.07378641},
				&EMADelta{Value: 5.031, Change: 1.10943396},
			},
		},
		{
			title: "It should calculate 5",
			args: IndicatorInputArg{
				Limit:  5,
				Period: 4,
			},
			candles: []*Candlestick{
				&Candlestick{Close: 3.1},
				&Candlestick{Close: 2.8},
				&Candlestick{Close: 3.4},
				&Candlestick{Close: 2.1},
				&Candlestick{Close: 9},
			},
			expected: []*EMADelta{
				nil,
				nil,
				nil,
				&EMADelta{Value: 2.85, Change: 0},
				&EMADelta{Value: 5.31, Change: 0.86315789},
			},
		},
	}

	for ti, v := range emaTests {
		cs, _ := NewCandlesticks(IntervalOneDay, 100)
		for _, c := range v.candles {
			cs.AppendCandlestick(c)
		}
		if err := cs.GenerateIndicator(IndicatorTypeEMA, v.args); err != nil {
			t.Fatalf("Expected ok but got error %+v for %s", err, v.title)
		}
		for i := range v.candles {
			c := cs.ItemAtIndex(i)
			if v.expected[i] == nil {
				if c.Indicators != nil {
					t.Errorf("Expected nil but got %+v in Test Idx: %+v idx: %+v for %s", c.Indicators, ti, i, v.title)
				}
				continue
			} else if v.expected[i] != nil && c.Indicators == nil {
				t.Errorf("Expected non nil but got nil for %s", v.title)
				continue
			}
			period := v.args.Period
			if !almostEqual(c.Indicators.EMAs[period].Value, v.expected[i].Value, 0.0001) {
				t.Errorf("Expected value %+v but got %+v for test  %+v index %+v for %s",
					v.expected[i].Value,
					c.Indicators.EMAs[period].Value,
					ti,
					i,
					v.title)
			}
			if !almostEqual(c.Indicators.EMAs[period].Change, v.expected[i].Change, 0.0001) {
				t.Errorf("Expected change %+v but got %+v for test  %+v index %+v for %s",
					v.expected[i].Change,
					c.Indicators.EMAs[period].Change,
					ti,
					i,
					v.title)
			}
		}
	}
}
