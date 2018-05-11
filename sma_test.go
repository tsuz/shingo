package shingo

import (
	"math"
	"testing"
)

func TestAppendSMA(t *testing.T) {

	smaTests := []struct {
		title    string
		arg      IndicatorInputArg
		candles  []*Candlestick
		expected []*SMADelta
	}{
		{
			title: "When limit is zero, generate on as many candlestick as possible",
			arg: IndicatorInputArg{
				Limit:  0,
				Period: 4,
			},
			candles: []*Candlestick{
				&Candlestick{Close: 2.8},
				&Candlestick{Close: 3.4},
				&Candlestick{Close: 2.1},
				&Candlestick{Close: 9},
			},
			expected: []*SMADelta{
				nil,
				nil,
				nil,
				&SMADelta{Value: 4.325, Change: 0},
			},
		},
		{
			title: "When candlestick is less than period, return none as it can't be computed",
			arg: IndicatorInputArg{
				Limit:  1,
				Period: 4,
			},
			candles: []*Candlestick{
				&Candlestick{Close: 3.4},
				&Candlestick{Close: 2.1},
				&Candlestick{Close: 9},
			},
			expected: []*SMADelta{
				nil,
				nil,
				nil,
			},
		},
		// requesting the last one to have
		// SMADelta indicator of 4 periods
		// there are 5 candles so only have
		// the last one set
		{
			title: "When limit is set to 1, it should return that many count",
			arg: IndicatorInputArg{
				Limit:  1,
				Period: 4,
			},
			candles: []*Candlestick{
				&Candlestick{Close: 1.9},
				&Candlestick{Close: 2.8},
				&Candlestick{Close: 3.4},
				&Candlestick{Close: 2.1},
				&Candlestick{Close: 9},
			},
			expected: []*SMADelta{
				nil,
				nil,
				nil,
				nil,
				&SMADelta{Value: 4.325, Change: (4.325 - 2.55) / 2.55},
			},
		},

		// requesting the last two to have
		// SMADelta indicator of 4 periods
		{
			title: "When limit is set to > 1, it should return that many count",
			arg: IndicatorInputArg{
				Limit:  2,
				Period: 4,
			},
			candles: []*Candlestick{
				&Candlestick{Close: 8.9},  // out of limit range
				&Candlestick{Close: 10.2}, // out of limit range
				&Candlestick{Close: 8},    // does not have 4 periods total including previous
				&Candlestick{Close: 3.9},
				&Candlestick{Close: 4.2},
				&Candlestick{Close: 3.1},
				&Candlestick{Close: 2.2},
				&Candlestick{Close: 1.9},
				&Candlestick{Close: 2.8},
				&Candlestick{Close: 3.4},
				&Candlestick{Close: 2.1},
				&Candlestick{Close: 9},
			},
			expected: []*SMADelta{
				nil,
				nil,
				nil,
				nil,
				nil,
				nil,
				nil,
				nil,
				nil,
				nil,
				&SMADelta{Value: 2.55, Change: 2.55/2.575 - 1},
				&SMADelta{Value: 4.325, Change: (4.325 - 2.55) / 2.55},
			},
		},

		// requesting the last 10 to have
		// SMADelta indicator of 4 periods
		{
			title: "It should return less than number that can be generated",
			arg: IndicatorInputArg{
				Limit:  10,
				Period: 4,
			},
			candles: []*Candlestick{
				&Candlestick{Close: 8.9},  // out of limit range
				&Candlestick{Close: 10.2}, // out of limit range
				&Candlestick{Close: 8},    // does not have 4 periods total including previous
				&Candlestick{Close: 3.9},
				&Candlestick{Close: 4.2},
				&Candlestick{Close: 3.1},
				&Candlestick{Close: 2.2},
				&Candlestick{Close: 1.9},
				&Candlestick{Close: 2.8},
				&Candlestick{Close: 3.4},
				&Candlestick{Close: 2.1},
				&Candlestick{Close: 9},
			},
			expected: []*SMADelta{
				nil,
				nil,
				nil,
				&SMADelta{Value: 7.75, Change: 0},
				&SMADelta{Value: 6.575, Change: 6.575/7.75 - 1},
				&SMADelta{Value: 4.8, Change: 4.8/6.575 - 1},
				&SMADelta{Value: 3.35, Change: 3.35/4.8 - 1},
				&SMADelta{Value: 2.85, Change: 2.85/3.35 - 1},
				&SMADelta{Value: 2.5, Change: 2.5/2.85 - 1},
				&SMADelta{Value: 2.575, Change: (2.575 - 2.5) / 2.5},
				&SMADelta{Value: 2.55, Change: 2.55/2.575 - 1},
				&SMADelta{Value: 4.325, Change: (4.325 - 2.55) / 2.55},
			},
		},
	}

	for ti, v := range smaTests {
		cs, _ := NewCandlesticks(IntervalOneDay, 100)
		for _, c := range v.candles {
			cs.AppendCandlestick(c)
		}

		period := v.arg.Period

		if err := cs.GenerateIndicator(IndicatorTypeSMA, v.arg); err != nil {
			t.Fatalf("Expected ok but got error %+v for %s", err, v.title)
		}

		for i := range v.candles {
			c := cs.ItemAtIndex(i)
			sma := c.GetSMA(period)
			if v.expected[i] == nil {
				if sma != nil {
					t.Fatalf("Expected nil but got %+v for %s", c.Indicators, v.title)
				}
				continue
			} else if v.expected[i] != nil && sma == nil {
				t.Fatalf("Expected indicators to be non nil but got nil for %s", v.title)
			}
			if !almostEqual(sma.Value, v.expected[i].Value, 0.0001) {
				t.Errorf("Expected value %+v but got %+v for test  %+v index %+v for %s",
					sma.Value,
					v.expected[i].Value,
					ti,
					i,
					v.title)
			}
			if !almostEqual(sma.Change, v.expected[i].Change, 0.0001) {
				t.Errorf("Expected change %+v but got %+v for test  %+v index %+v for %s",
					sma.Change,
					v.expected[i].Change,
					ti,
					i,
					v.title)
			}
		}
	}
}

func almostEqual(x, y, epsilon float64) bool {
	return math.Nextafter(x, y) == y ||
		math.Abs(x-y) < epsilon // epislon hack
}
