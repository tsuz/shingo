package shingo

import (
	"github.com/pkg/errors"
)

func (cs *Candlesticks) AppendSMA(arg IndicatorInputArg) error {
	period := arg.Period
	limit := arg.Limit
	if period < 1 {
		return errors.Wrap(nil, "Period must be positive")
	}
	list := cs
	l := cs.Total()
	if l < 1 {
		return nil
	}
	var count int
	var firstVal float64
	var avg float64
	var prevAvg float64

	for i := l - 1; i >= 0; i-- {

		v := list.ItemAtIndex(i)

		// first one. set initial average
		if i == l-1 {
			count++
			firstVal = v.Close
			// not enough periods
			if i-(period-1) < 0 {
				continue
			}
			for j := i; j > i-(period); j-- {
				p := list.ItemAtIndex(j)
				avg += p.Close
			}
			avg = avg / float64(period)

			setSMAIndicator(v, period, avg)

			continue
		}

		// for all other ones
		// don't even append indicators if
		// it's not required to set. This saves
		// a lot of allocations when running
		// detection batches
		if i-(period-1) < 0 {
			break
		}
		n := list.ItemAtIndex(i - (period - 1))
		prevAvg = avg
		avg += (n.Close - firstVal) / float64(period)
		firstVal = v.Close
		// Next candle stick
		cn := list.ItemAtIndex(i + 1)
		cn.Indicators.SMAs[period].Change = prevAvg/avg - 1

		if limit > 0 && count >= limit {
			break
		}

		setSMAIndicator(v, period, avg)

		count++

	}
	return nil
}

func setSMAIndicator(v *Candlestick, period int, avg float64) {
	if v.Indicators == nil {
		v.Indicators = &Indicators{}
	}
	if v.Indicators.SMAs == nil {
		v.Indicators.SMAs = make(map[int]*SMADelta)
	}
	v.Indicators.SMAs[period] = &SMADelta{
		Value: avg,
	}
}
