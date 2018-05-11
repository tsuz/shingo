package shingo

import (
	"testing"
)

func TestIchimokuCloud(t *testing.T) {
	testCases := []struct {
		arg      IndicatorInputArg
		candles  []*Candlestick
		expected []*IchimokuCloudDelta
	}{
		{
			arg: IndicatorInputArg{
				IchimokuCloudKijun:  26,
				IchimokuCloudTenkan: 9,
				Limit:               30,
			},
			// doing this in hypothesis will look like this
			// it will just allocate all kijun
			// and see if it has any effect
			candles: []*Candlestick{
				&Candlestick{High: 100, Low: 96.19478155},
				&Candlestick{High: 98.81323044, Low: 97.19536099},
				&Candlestick{High: 97.72607321, Low: 92.43009678},
				&Candlestick{High: 98.74199781, Low: 92.46008389},
				&Candlestick{High: 99.64794697, Low: 97.5653654},
				&Candlestick{High: 100.8729436, Low: 93.41954471},
				&Candlestick{High: 101.369714, Low: 96.17759904},
				&Candlestick{High: 102.6772269, Low: 101.3640108},
				&Candlestick{High: 103.4477903, Low: 95.09547817},
				&Candlestick{High: 104.4410349, Low: 98.38312574},
				&Candlestick{High: 105.9360566, Low: 104.8939137},
				&Candlestick{High: 103.9737331, Low: 98.80595814},
				&Candlestick{High: 105.3275966, Low: 96.19665714},
				&Candlestick{High: 106.4343921, Low: 104.7522734},
				&Candlestick{High: 107.5452098, Low: 104.5212317},
				&Candlestick{High: 108.4881512, Low: 100.3822134},
				&Candlestick{High: 108.9629734, Low: 106.9588466},
				&Candlestick{High: 107.0446541, Low: 102.2707067},
				&Candlestick{High: 107.4052403, Low: 106.582743},
				&Candlestick{High: 106.1942337, Low: 103.4501847},
				&Candlestick{High: 107.4426701, Low: 101.7474732},
				&Candlestick{High: 108.7271967, Low: 98.83746035},
				&Candlestick{High: 107.5575668, Low: 98.65200935},
				&Candlestick{High: 108.4980317, Low: 106.9756229},
				&Candlestick{High: 107.0199466, Low: 102.9748333},
				&Candlestick{High: 105.11815, Low: 95.89888709},
				&Candlestick{High: 104.1014068, Low: 101.5395117},
				&Candlestick{High: 103.0379469, Low: 98.45218655},
				&Candlestick{High: 104.5844922, Low: 95.19718538},
				&Candlestick{High: 104.1217667, Low: 100.7465228},
			},
			expected: []*IchimokuCloudDelta{
				nil, nil, nil, nil, nil,
				nil, nil, nil,
				&IchimokuCloudDelta{Tenkan: 97.93894356},
				&IchimokuCloudDelta{Tenkan: 98.43556583},
				&IchimokuCloudDelta{Tenkan: 99.18307667},
				&IchimokuCloudDelta{Tenkan: 99.19807022},
				&IchimokuCloudDelta{Tenkan: 99.67780063},
				&IchimokuCloudDelta{Tenkan: 99.92696841},
				&IchimokuCloudDelta{Tenkan: 101.320344},
				&IchimokuCloudDelta{Tenkan: 101.7918147},
				&IchimokuCloudDelta{Tenkan: 102.0292258},
				&IchimokuCloudDelta{Tenkan: 102.5798153},
				&IchimokuCloudDelta{Tenkan: 102.5798153},
				&IchimokuCloudDelta{Tenkan: 102.5798153},
				&IchimokuCloudDelta{Tenkan: 102.5798153},
				&IchimokuCloudDelta{Tenkan: 103.9002169},
				&IchimokuCloudDelta{Tenkan: 103.8074914},
				&IchimokuCloudDelta{Tenkan: 103.8074914},
				&IchimokuCloudDelta{Tenkan: 103.8074914},
				&IchimokuCloudDelta{Tenkan: 102.3130419, Kijun: 100.6965351, SenkouA: (102.3130419 + 100.6965351) / 2},
				&IchimokuCloudDelta{Tenkan: 102.3130419, Kijun: 100.6965351, SenkouA: (102.3130419 + 100.6965351) / 2},
				&IchimokuCloudDelta{Tenkan: 102.3130419, Kijun: 100.6965351, SenkouA: (102.3130419 + 100.6965351) / 2},
				&IchimokuCloudDelta{Tenkan: 101.9621911, Kijun: 100.7115286, SenkouA: (101.9621911 + 100.7115286) / 2},
				&IchimokuCloudDelta{Tenkan: 101.9621911, Kijun: 101.191259, SenkouA: (101.9621911 + 101.191259) / 2},
			},
		},
		{
			// case of getting just one value in the candlestick
			// to detect live trends
			arg: IndicatorInputArg{
				IchimokuCloudKijun:  26,
				IchimokuCloudTenkan: 9,
				Limit:               1,
			},
			candles: []*Candlestick{
				&Candlestick{High: 100, Low: 96.19478155},
				&Candlestick{High: 98.81323044, Low: 97.19536099},
				&Candlestick{High: 97.72607321, Low: 92.43009678},
				&Candlestick{High: 98.74199781, Low: 92.46008389},
				&Candlestick{High: 99.64794697, Low: 97.5653654},
				&Candlestick{High: 100.8729436, Low: 93.41954471},
				&Candlestick{High: 101.369714, Low: 96.17759904},
				&Candlestick{High: 102.6772269, Low: 101.3640108},
				&Candlestick{High: 103.4477903, Low: 95.09547817},
				&Candlestick{High: 104.4410349, Low: 98.38312574},
				&Candlestick{High: 105.9360566, Low: 104.8939137},
				&Candlestick{High: 103.9737331, Low: 98.80595814},
				&Candlestick{High: 105.3275966, Low: 96.19665714},
				&Candlestick{High: 106.4343921, Low: 104.7522734},
				&Candlestick{High: 107.5452098, Low: 104.5212317},
				&Candlestick{High: 108.4881512, Low: 100.3822134},
				&Candlestick{High: 108.9629734, Low: 106.9588466},
				&Candlestick{High: 107.0446541, Low: 102.2707067},
				&Candlestick{High: 107.4052403, Low: 106.582743},
				&Candlestick{High: 106.1942337, Low: 103.4501847},
				&Candlestick{High: 107.4426701, Low: 101.7474732},
				&Candlestick{High: 108.7271967, Low: 98.83746035},
				&Candlestick{High: 107.5575668, Low: 98.65200935},
				&Candlestick{High: 108.4980317, Low: 106.9756229},
				&Candlestick{High: 107.0199466, Low: 102.9748333},
				&Candlestick{High: 105.11815, Low: 95.89888709},
				&Candlestick{High: 104.1014068, Low: 101.5395117},
				&Candlestick{High: 103.0379469, Low: 98.45218655},
				&Candlestick{High: 104.5844922, Low: 95.19718538},
				&Candlestick{High: 104.1217667, Low: 100.7465228},
			},
			expected: []*IchimokuCloudDelta{
				nil, nil, nil, nil, nil,
				nil, nil, nil, nil,
				nil, nil, nil, nil, nil,
				nil, nil, nil, nil, nil,
				nil, nil, nil, nil, nil,
				nil, nil, nil, nil, nil,
				&IchimokuCloudDelta{Tenkan: 101.9621911, Kijun: 101.191259, SenkouA: (101.9621911 + 101.191259) / 2},
			},
		},
		{
			// adding extra values 26+9 in front of the last candlestick
			// should not matter. Should get same result as last time
			arg: IndicatorInputArg{
				IchimokuCloudKijun:  26,
				IchimokuCloudTenkan: 9,
				Limit:               1,
			},
			candles: []*Candlestick{
				&Candlestick{High: 100, Low: 96.19478155},
				&Candlestick{High: 200, Low: 9.19478155},
				&Candlestick{High: 200, Low: 9.19478155},
				&Candlestick{High: 200, Low: 9.19478155},
				&Candlestick{High: 98.81323044, Low: 97.19536099},
				&Candlestick{High: 97.72607321, Low: 92.43009678},
				&Candlestick{High: 98.74199781, Low: 92.46008389},
				&Candlestick{High: 99.64794697, Low: 97.5653654},
				&Candlestick{High: 100.8729436, Low: 93.41954471},
				&Candlestick{High: 101.369714, Low: 96.17759904},
				&Candlestick{High: 102.6772269, Low: 101.3640108},
				&Candlestick{High: 103.4477903, Low: 95.09547817},
				&Candlestick{High: 104.4410349, Low: 98.38312574},
				&Candlestick{High: 105.9360566, Low: 104.8939137},
				&Candlestick{High: 103.9737331, Low: 98.80595814},
				&Candlestick{High: 105.3275966, Low: 96.19665714},
				&Candlestick{High: 106.4343921, Low: 104.7522734},
				&Candlestick{High: 107.5452098, Low: 104.5212317},
				&Candlestick{High: 108.4881512, Low: 100.3822134},
				&Candlestick{High: 108.9629734, Low: 106.9588466},
				&Candlestick{High: 107.0446541, Low: 102.2707067},
				&Candlestick{High: 107.4052403, Low: 106.582743},
				&Candlestick{High: 106.1942337, Low: 103.4501847},
				&Candlestick{High: 107.4426701, Low: 101.7474732},
				&Candlestick{High: 108.7271967, Low: 98.83746035},
				&Candlestick{High: 107.5575668, Low: 98.65200935},
				&Candlestick{High: 108.4980317, Low: 106.9756229},
				&Candlestick{High: 107.0199466, Low: 102.9748333},
				&Candlestick{High: 105.11815, Low: 95.89888709},
				&Candlestick{High: 104.1014068, Low: 101.5395117},
				&Candlestick{High: 103.0379469, Low: 98.45218655},
				&Candlestick{High: 104.5844922, Low: 95.19718538},
				&Candlestick{High: 104.1217667, Low: 100.7465228},
			},
			expected: []*IchimokuCloudDelta{
				nil, nil, nil, nil,
				nil, nil, nil, nil, nil,
				nil, nil, nil,
				nil, nil, nil, nil, nil,
				nil, nil, nil, nil, nil,
				nil, nil, nil, nil, nil,
				nil, nil, nil, nil, nil,
				&IchimokuCloudDelta{Tenkan: 101.9621911, Kijun: 101.191259, SenkouA: (101.9621911 + 101.191259) / 2},
			},
		},
	}

	for ti, v := range testCases {

		cs, _ := NewCandlesticks(IntervalOneDay, 100)
		for _, c := range v.candles {
			cs.AppendCandlestick(c)
		}

		if err := cs.GenerateIndicator(IndicatorTypeIchimokuCloud, v.arg); err != nil {
			t.Errorf("Expected ok but got error appending EMA to period: 12, args: %+v for %s", v.arg, "no title")
		}

		for i, c := range v.candles {
			arg := v.arg
			ic := c.GetIchimokuCloud(arg.IchimokuCloudKijun, arg.IchimokuCloudTenkan)
			if v.expected[i] == nil {
				if ic != nil {
					t.Errorf("Expected nil but got %+v at test %d index %d", ic, ti, i)
				}
				continue
			}
			if !almostEqual(ic.Tenkan, v.expected[i].Tenkan, 0.0001) {
				t.Errorf("Expected tenkan %+v but got %+v for test  %+v index %+v",
					v.expected[i].Tenkan,
					ic.Tenkan,
					ti,
					i)
			}
			if v.expected[i].Kijun == 0 {
				continue
			}
			if !almostEqual(ic.Kijun, v.expected[i].Kijun, 0.0001) {
				t.Errorf("Expected kijun %+v but got %+v for test  %+v index %+v",
					v.expected[i].Kijun,
					ic.Kijun,
					ti,
					i)
			}

			if !almostEqual(ic.SenkouA, v.expected[i].SenkouA, 0.0001) {
				t.Errorf("Expected senkou span A of %+v but got %+v for test  %+v index %+v",
					v.expected[i].SenkouA,
					ic.SenkouA,
					ti,
					i)
			}

			if !almostEqual(ic.SenkouB, v.expected[i].SenkouB, 0.0001) {
				t.Errorf("Expected senkou span B of %+v but got %+v for test  %+v index %+v",
					v.expected[i].SenkouB,
					ic.SenkouB,
					ti,
					i)
			}
		}
	}
}
