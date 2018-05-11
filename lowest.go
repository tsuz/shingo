package shingo

import (
	"github.com/pkg/errors"
)

// AppendLowest generates lowest value within that period of time up to that candlestick
func (cs *Candlesticks) AppendLowest(arg IndicatorInputArg) error {
	p := arg.Period
	if p < 1 {
		return errors.New("Period must be larger than 0")
	}
	t := cs.Total()
	l := arg.Limit
	if l < 1 {
		l = t
	}

	startIdx := t - 1 - l
	if startIdx < 0 {
		startIdx = 0
	}

	lastLowIdx := -1
	for i := startIdx; i < t; i++ {
		v := cs.ItemAtIndex(i)
		var lowest float64

		if lastLowIdx >= 0 && lastLowIdx > i-p+1 {
			// get from previous low
			pv := cs.ItemAtIndex(i - 1)
			if pv != nil {
				ll := pv.GetLowest(p)
				if ll == nil {
					continue
				}
				lastLow := *ll
				if lastLow < v.Close {
					lowest = lastLow
				} else {
					lowest = v.Close
					lastLowIdx = i
				}
			}
		}
		if lowest == 0.0 {
			lowest = v.Close
			lastLowIdx = i
			for j := i - p + 1; j < i; j++ {
				g := cs.ItemAtIndex(j)
				if g == nil {
					continue
				}
				if g.Close < lowest {
					lowest = g.Close
					lastLowIdx = j
				}
			}
		}
		if v.Indicators == nil {
			v.Indicators = &Indicators{}
		}
		if v.Indicators.Lowest == nil {
			v.Indicators.Lowest = make(map[int]*float64)
		}
		v.Indicators.Lowest[p] = &lowest
	}

	return nil
}

// GetLowest gets highest value for given past periods
func (c *Candlestick) GetLowest(period int) *float64 {
	if c.Indicators == nil || c.Indicators.Lowest == nil {
		return nil
	}
	return c.Indicators.Lowest[period]
}
