package shingo

import "errors"

// GenerateEMA generates EMA on each candlestick for given period and limit
func (cs *Candlesticks) AppendEMA(arg IndicatorInputArg) (err error) {
	limit := arg.Limit
	period := arg.Period

	if period < 1 {
		return errors.New("Period must be positive")
	}

	cl := cs.Total()
	if cl-period < 0 {
		return
	}

	if limit < 1 {
		limit = cl
	}
	// start calculating from twice the periods to get
	// accurate previous results
	startCalcIdx := 0
	startIdx := cl - limit
	var emaValue float64
	var prevEmaValue float64
	for i := startCalcIdx; i < cl; i++ {
		v := cs.ItemAtIndex(i)
		// if period is more than number of candles left,
		// then don't allocate anything
		if i-(period-1) < 0 {
			continue
		}
		if i-(period-1) == 0 {
			// if this is the first value, it needs to be the SMA of
			// last `period` periods
			for j := 0; j <= i; j++ {
				vj := cs.ItemAtIndex(j)
				emaValue += vj.Close
			}
			emaValue /= float64(period)
		} else {
			k := 2.0 / float64(period+1)
			prevEmaValue = emaValue
			emaValue = (v.Close-prevEmaValue)*k + prevEmaValue
		}
		// calculated which is required for next one
		// but don't allocate as it's not in the specified range
		if i < startIdx {
			continue
		}

		v.setEMA(period, emaValue, prevEmaValue)

	}
	return nil
}

// GetEMA returns EMA value for this candlestick for given period
func (c *Candlestick) GetEMA(period int) *EMADelta {
	if c.Indicators == nil || c.Indicators.EMAs == nil {
		return nil
	}
	return c.Indicators.EMAs[period]
}

func (c *Candlestick) setEMA(period int, emaValue float64, prevEma float64) {
	if c.Indicators == nil {
		c.Indicators = &Indicators{}
	}
	if c.Indicators.EMAs == nil {
		c.Indicators.EMAs = make(map[int]*EMADelta)
	}
	chg := emaValue/prevEma - 1
	if prevEma == 0 {
		chg = 0
	}
	c.Indicators.EMAs[period] = &EMADelta{
		Value:  emaValue,
		Change: chg,
	}
}
