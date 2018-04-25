package model

import (
	"fmt"
	"testing"
)

func TestAppendMACD(t *testing.T) {

	macdTests := []struct {
		title    string
		limit    int
		candles  []*Candlestick
		expected []*MACDDelta
	}{
		{
			title: "From 40 candles, it requires at least 34 candles so it returns none",
			limit: 40,
			candles: []*Candlestick{
				&Candlestick{Close: 2.8},
				&Candlestick{Close: 3.4},
				&Candlestick{Close: 2.1},
				&Candlestick{Close: 9},
				&Candlestick{Close: 4.9},
				&Candlestick{Close: 3.2},
				&Candlestick{Close: 1.8},
				&Candlestick{Close: 7.5},
				&Candlestick{Close: 7.8},
				&Candlestick{Close: 7.5},
				&Candlestick{Close: 4.9},
				&Candlestick{Close: 3.2},
				&Candlestick{Close: 4.9},
				&Candlestick{Close: 3.2},
				&Candlestick{Close: 1.8},
				&Candlestick{Close: 6.6},
				&Candlestick{Close: 7.8},
				&Candlestick{Close: 7.5},
				&Candlestick{Close: 7.8},
				&Candlestick{Close: 4.9},
				&Candlestick{Close: 6.6},
				&Candlestick{Close: 7.5},
				&Candlestick{Close: 7.5},
				&Candlestick{Close: 4.9},
				&Candlestick{Close: 3.2},
				&Candlestick{Close: 7.8},
				&Candlestick{Close: 4.9},
				&Candlestick{Close: 1.8},
				&Candlestick{Close: 7.8},
				&Candlestick{Close: 7.6},
				&Candlestick{Close: 7.3},
				&Candlestick{Close: 7.1},
				&Candlestick{Close: 6.8},
			},
			expected: []*MACDDelta{
				nil, nil, nil, nil, nil, nil,
				nil, nil, nil, nil, nil, nil,
				nil, nil, nil, nil, nil, nil,
				nil, nil, nil, nil, nil, nil,
				nil, nil, nil, nil, nil, nil,
				nil, nil, nil,
			},
		},
		{
			title: "From 40 candles, it returns 8",
			limit: 40,
			candles: []*Candlestick{
				&Candlestick{Close: 2.8},
				&Candlestick{Close: 3.4},
				&Candlestick{Close: 2.1},
				&Candlestick{Close: 9},
				&Candlestick{Close: 4.9},
				&Candlestick{Close: 3.2},
				&Candlestick{Close: 1.8},
				&Candlestick{Close: 7.5},
				&Candlestick{Close: 7.8},
				&Candlestick{Close: 7.5},
				&Candlestick{Close: 4.9},
				&Candlestick{Close: 3.2},
				&Candlestick{Close: 4.9},
				&Candlestick{Close: 3.2},
				&Candlestick{Close: 1.8},
				&Candlestick{Close: 6.6},
				&Candlestick{Close: 7.8},
				&Candlestick{Close: 7.5},
				&Candlestick{Close: 7.8},
				&Candlestick{Close: 4.9},
				&Candlestick{Close: 6.6},
				&Candlestick{Close: 7.5},
				&Candlestick{Close: 7.5},
				&Candlestick{Close: 4.9},
				&Candlestick{Close: 3.2},
				&Candlestick{Close: 7.8},
				&Candlestick{Close: 4.9},
				&Candlestick{Close: 1.8},
				&Candlestick{Close: 7.8},
				&Candlestick{Close: 7.6},
				&Candlestick{Close: 7.3},
				&Candlestick{Close: 7.1},
				&Candlestick{Close: 6.8},
				&Candlestick{Close: 6.3},
				&Candlestick{Close: 6},
				&Candlestick{Close: 7.1},
				&Candlestick{Close: 6.1},
				&Candlestick{Close: 5.9},
				&Candlestick{Close: 5.1},
				&Candlestick{Close: 4.8},
				&Candlestick{Close: 5.3},
			},
			expected: []*MACDDelta{
				nil, nil, nil, nil, nil, nil,
				nil, nil, nil, nil, nil, nil,
				nil, nil, nil, nil, nil, nil,
				nil, nil, nil, nil, nil, nil,
				nil, nil, nil, nil, nil, nil,
				nil, nil, nil,
				&MACDDelta{
					MACDValue:     6.338822758123295 - 5.827396793334697,
					SignalValue:   0.4420618107841297,
					MACDHistogram: 6.338822758123295 - 5.827396793334697 - 0.4420618107841297,
				},
				&MACDDelta{
					MACDValue:     6.2867118179239325 - 5.840186690948596,
					SignalValue:   0.44300102610456293,
					MACDHistogram: 6.2867118179239325 - 5.840186690948596 - 0.44300102610456293,
				},
				&MACDDelta{
					MACDValue:     6.4117955403272315 - 5.933538857149305,
					SignalValue:   0.45010768419554115,
					MACDHistogram: 6.4117955403272315 - 5.933538857149305 - 0.45010768419554115,
				},
				&MACDDelta{
					MACDValue:     6.3638413862249035 - 5.945873627834541,
					SignalValue:   0.4437271307574296,
					MACDHistogram: 6.3638413862249035 - 5.945873627834541 - 0.4437271307574296,
				},
				&MACDDelta{
					MACDValue:     6.292502581023514 - 5.942474392012002,
					SignalValue:   0.4250253064995333,
					MACDHistogram: 6.292502581023514 - 5.942474392012002 - 0.4250253064995333,
				},
				&MACDDelta{
					MACDValue:     6.109095684062098 - 5.880047039563912,
					SignalValue:   0.38584887751455155,
					MACDHistogram: 6.109095684062098 - 5.880047039563912 - 0.38584887751455155,
				},
				&MACDDelta{
					MACDValue:     5.907756767853347 - 5.8000155539322265,
					SignalValue:   0.33022730366163033,
					MACDHistogram: 5.907756767853347 - 5.8000155539322265 - 0.33022730366163033,
				},
				&MACDDelta{
					MACDValue:     5.814283776957502 - 5.762964401385848,
					SignalValue:   0.27443855956659446,
					MACDHistogram: 5.814283776957502 - 5.762964401385848 - 0.27443855956659446,
				},
			},
		},
		{
			title: "Requesting for 40 MACD values and returns 1 since that's the limit",
			limit: 1,
			candles: []*Candlestick{
				&Candlestick{Close: 2.8},
				&Candlestick{Close: 3.4},
				&Candlestick{Close: 2.1},
				&Candlestick{Close: 9},
				&Candlestick{Close: 4.9},
				&Candlestick{Close: 3.2},
				&Candlestick{Close: 1.8},
				&Candlestick{Close: 7.5},
				&Candlestick{Close: 7.8},
				&Candlestick{Close: 7.5},
				&Candlestick{Close: 4.9},
				&Candlestick{Close: 3.2},
				&Candlestick{Close: 4.9},
				&Candlestick{Close: 3.2},
				&Candlestick{Close: 1.8},
				&Candlestick{Close: 6.6},
				&Candlestick{Close: 7.8},
				&Candlestick{Close: 7.5},
				&Candlestick{Close: 7.8},
				&Candlestick{Close: 4.9},
				&Candlestick{Close: 6.6},
				&Candlestick{Close: 7.5},
				&Candlestick{Close: 7.5},
				&Candlestick{Close: 4.9},
				&Candlestick{Close: 3.2},
				&Candlestick{Close: 7.8},
				&Candlestick{Close: 4.9},
				&Candlestick{Close: 1.8},
				&Candlestick{Close: 7.8},
				&Candlestick{Close: 7.6},
				&Candlestick{Close: 7.3},
				&Candlestick{Close: 7.1},
				&Candlestick{Close: 6.8},
				&Candlestick{Close: 6.3},
				&Candlestick{Close: 6},
				&Candlestick{Close: 7.1},
				&Candlestick{Close: 6.1},
				&Candlestick{Close: 5.9},
				&Candlestick{Close: 5.1},
				&Candlestick{Close: 4.8},
				&Candlestick{Close: 5.3},
			},
			expected: []*MACDDelta{
				nil, nil, nil, nil, nil, nil,
				nil, nil, nil, nil, nil, nil,
				nil, nil, nil, nil, nil, nil,
				nil, nil, nil, nil, nil, nil,
				nil, nil, nil, nil, nil, nil,
				nil, nil, nil, nil, nil, nil,
				nil, nil, nil, nil,
				&MACDDelta{
					MACDValue:     5.814283776957502 - 5.762964401385848,
					SignalValue:   0.27443855956659446,
					MACDHistogram: 5.814283776957502 - 5.762964401385848 - 0.27443855956659446,
				},
			},
		},
	}

	for _, v := range macdTests {
		cs, _ := NewCandlesticks(IntervalOneDay)
		for _, c := range v.candles {
			cs.AppendCandlestick(c)
		}

		args := IndicatorInputArg{
			MacdLarge:  12,
			MacdSmall:  26,
			MacdSignal: 9,
			Limit:      v.limit,
		}

		if err := cs.GenerateIndicator(IndicatorTypeMACD, args); err != nil {
			t.Errorf("Expected ok but got error appending EMA to period: 12, args: %+v for %s", args, v.title)
		}

		var str string
		for i := range v.candles {
			c := cs.ItemAtIndex(i)
			str += fmt.Sprintf("%v,", c.Close)
			if v.expected[i] == nil {
				if c.Indicators != nil && c.Indicators.MACDs[12][26][9] != nil {
					t.Errorf("Expected nil but got %+v for test %s index %d", c.Indicators, v.title, i)
				}
				continue
			} else if v.expected[i] != nil && (c.Indicators.MACDs == nil || c.Indicators.MACDs[12] == nil || c.Indicators.MACDs[12][26] == nil || c.Indicators.MACDs[12][26][9] == nil) {
				t.Errorf("Expected non nil but got nil %+v for test %s index %d", c.Indicators.MACDs, v.title, i)
				continue
			}
			mcds := c.Indicators.MACDs[12][26][9]
			if !almostEqual(mcds.MACDValue, v.expected[i].MACDValue) {
				t.Errorf("Expected MACDValue %+v but got %+v for test %s index %d",
					v.expected[i].MACDValue,
					mcds.MACDValue,
					v.title,
					i)
			}
			if !almostEqual(mcds.MACDHistogram, v.expected[i].MACDHistogram) {
				t.Errorf("Expected MACDHistogram %+v but got %+v for test %s index %d",
					v.expected[i].MACDHistogram,
					mcds.MACDHistogram,
					v.title,
					i)
			}
			if !almostEqual(mcds.SignalValue, v.expected[i].SignalValue) {
				t.Errorf("Expected SignalValue %+v but got %+v for test %s index %d",
					v.expected[i].SignalValue,
					mcds.SignalValue,
					v.title,
					i)
			}
		}
	}
}
