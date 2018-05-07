package shingo

import (
	"github.com/pkg/errors"
)

// AppendHighest generates highest value within that period of time up to that candlestick
func (cs *Candlesticks) AppendHighest(arg IndicatorInputArg) error {
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

	lastHighIdx := -1
	for i := startIdx; i < t; i++ {
		v := cs.ItemAtIndex(i)
		var highest float64

		if lastHighIdx >= 0 && lastHighIdx > i-p+1 {
			// get from previous high
			pv := cs.ItemAtIndex(i - 1)
			if pv != nil {
				lastHigh := pv.Indicators.Highest[p]
				if lastHigh > v.Close {
					highest = lastHigh
				} else {
					highest = v.Close
					lastHighIdx = i
				}
			}
		}
		if highest == 0.0 {
			for j := i - p + 1; j < i; j++ {
				g := cs.ItemAtIndex(j)
				if g == nil {
					continue
				}
				if g.Close > highest {
					highest = g.Close
					lastHighIdx = j
				}
			}
		}
		if highest == 0.0 {
			highest = v.Close
			lastHighIdx = i
		}
		if v.Indicators == nil {
			v.Indicators = &Indicators{}
		}
		if v.Indicators.Highest == nil {
			v.Indicators.Highest = make(map[int]float64)
		}
		v.Indicators.Highest[p] = highest
	}

	return nil
}
