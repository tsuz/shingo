package shingo

import (
	"math"
	"testing"
)

// 52.22
// 52.78
// 53.02
// 53.67
// 53.67
// 53.74
// 53.45
// 53.72
// 53.39
// 52.51	(STDEVP)
// 52.32	0.51
// 51.45	0.73
// 51.60	0.86
// 52.43	0.83
// 52.47	0.79
// 52.91	0.72
// 52.07	0.68
// 53.12	0.58
// 52.77	0.51
// 52.73	0.52
// 52.09	0.53
// 53.19	0.48
// 53.73	0.49
// 53.87	0.58
// 53.85	0.62
// 53.88	0.67
// 54.08	0.62
// 54.14	0.66
// 54.50	0.69
// 54.30	0.65
// 54.40	0.36
// 54.16	0.24

func TestStdDev(t *testing.T) {
	sdTests := []struct {
		title    string
		arg      IndicatorInputArg
		candles  []*Candlestick
		expected []*StdDevDelta
	}{
		{
			title: "Standard Deviation test",
			arg: IndicatorInputArg{
				Type:   IndicatorTypeStdDev,
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
			expected: []*StdDevDelta{
				nil,
				nil,
				nil,
				nil,
				nil,
				nil,
				nil,
				nil,
				nil,
				&StdDevDelta{Value: 0.523018},
				&StdDevDelta{Value: 0.505411},
				&StdDevDelta{Value: 0.730122},
				&StdDevDelta{Value: 0.857364},
				&StdDevDelta{Value: 0.833642},
				&StdDevDelta{Value: 0.788707},
				&StdDevDelta{Value: 0.716251},
				&StdDevDelta{Value: 0.675498},
				&StdDevDelta{Value: 0.584679},
				&StdDevDelta{Value: 0.507870},
				&StdDevDelta{Value: 0.518353},
				&StdDevDelta{Value: 0.526061},
				&StdDevDelta{Value: 0.480964},
				&StdDevDelta{Value: 0.490176},
				&StdDevDelta{Value: 0.578439},
				&StdDevDelta{Value: 0.622905},
				&StdDevDelta{Value: 0.670093},
				&StdDevDelta{Value: 0.622025},
				&StdDevDelta{Value: 0.661064},
				&StdDevDelta{Value: 0.690358},
				&StdDevDelta{Value: 0.651152},
				&StdDevDelta{Value: 0.360466},
				&StdDevDelta{Value: 0.242959},
			},
		},
	}

	for _, st := range sdTests {
		cs, _ := NewCandlesticks(IntervalOneDay, 100)
		for _, c := range st.candles {
			cs.AppendCandlestick(c)
		}
		if err := cs.GenerateIndicator(IndicatorTypeStdDev, st.arg); err != nil {
			t.Fatalf("Error appending stddev: %+v", err)
		}
		for i, e := range st.expected {
			v := cs.ItemAtIndex(i)
			delta := v.GetStdDev(st.arg.Period)
			if e == nil && delta == nil {
				continue
			}
			if e != nil && delta == nil {
				t.Fatalf("Expected non-nil: %+v but got nil", e)
			}

			if e == nil && delta != nil {
				t.Fatalf("Expected nil but got non nil %+v  %d", delta, i)
			}

			if !equalWithinPct(e.Value, delta.Value, 0.005) {
				t.Errorf("Expected value to be: %+v but got %+v", e.Value, delta.Value)
			}
		}
	}
}

func equalWithinPct(x, y, d float64) bool {
	if x == 0.0 && y == 0.0 {
		return true
	} else if x != 0.0 {
		return math.Abs(y/x-1) < d
	} else {
		return math.Abs(x/y-1) < d
	}
}
