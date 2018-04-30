package shingo

import (
	"fmt"
	"math"
)

// AppendATR appends Average True Range to candlesticks
func (cs *Candlesticks) AppendATR(arg IndicatorInputArg) error {
	limit := arg.Limit
	period := arg.Period
	len := cs.Total()
	if len < 1 {
		return nil
	}
	if limit < 1 {
		limit = len
	}
	if period < 1 {
		return fmt.Errorf("period must be larger than 0")
	}
	var count int = 1
	startIdx := (len - 1) - limit - period
	if startIdx < 0 {
		startIdx = 0
	}
	var firstTRTotal float64
	for i := startIdx; i < len; i++ {
		p := cs.ItemAtIndex(i - 1)
		v := cs.ItemAtIndex(i)
		var tr float64
		if p != nil {
			tr = findHighestValue(v.High-v.Low, math.Abs(v.High-p.Close), math.Abs(v.Low-p.Close))
			prev := getATR(p, period)
			if count < period {
				firstTRTotal += tr
			} else if count == period {
				setATRIndicator(v, period, (firstTRTotal+tr)/float64(period), 0)
			} else {
				tr = (prev.Value*float64(period-1) + tr) / float64(period)
				var chg float64
				if prev.Value > 0.0 {
					chg = tr/prev.Value - 1
				} else {
					chg = 0
				}
				setATRIndicator(v, period, tr, chg)
			}
		} else {
			firstTRTotal = v.High - v.Low
		}
		count++
	}
	return nil
}

func findHighestValue(vals ...float64) float64 {
	if len(vals) < 1 {
		return 0
	}
	f := vals[0]
	for _, v := range vals {
		if v > f {
			f = v
		}
	}
	return f
}

func findLowestValue(vals ...float64) float64 {
	if len(vals) < 1 {
		return 0
	}
	f := vals[0]
	for _, v := range vals {
		if v < f {
			f = v
		}
	}
	return f
}

func setATRIndicator(v *Candlestick, period int, val float64, chg float64) {
	if v.Indicators == nil {
		v.Indicators = &Indicators{}
	}
	if v.Indicators.ATRs == nil {
		v.Indicators.ATRs = make(map[int]*ATRDelta)
	}
	if v.Indicators.ATRs[period] == nil {
		v.Indicators.ATRs[period] = &ATRDelta{Value: val, Change: chg}
	}
}

func getATR(p *Candlestick, period int) *ATRDelta {
	if p == nil {
		return nil
	}
	if p.Indicators == nil {
		return nil
	}
	if p.Indicators.ATRs == nil {
		return nil
	}
	if p.Indicators.ATRs[period] == nil {
		return nil
	}
	return p.Indicators.ATRs[period]
}
