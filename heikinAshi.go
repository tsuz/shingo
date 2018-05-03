package shingo

// AppendHeikinAshi appends heikin ashi values for each candlestick
func (cs *Candlesticks) AppendHeikinAshi(arg IndicatorInputArg) error {
	limit := arg.Limit
	cs.mux.Lock()
	defer cs.mux.Unlock()
	total := cs.Total()
	if total < 1 {
		return nil
	}
	if limit < 1 {
		limit = total
	}
	startIdx := total - 1 - limit
	if startIdx < 0 {
		startIdx = 0
	}
	for i := startIdx; i < total; i++ {
		p := cs.ItemAtIndex(i - 1)
		v := cs.ItemAtIndex(i)

		close := (v.Open + v.High + v.Low + v.Close) / 4
		prev := getHeikinAshiDelta(p)
		var open, high, low float64
		if prev == nil {
			open = v.Open
			low = v.Low
			high = v.High
		} else {
			open = (prev.Open + prev.Close) / 2
			high = findHighestValue(v.High, open, close)
			low = findLowestValue(v.Low, open, close)
		}

		setHeikinAshiDelta(v, open, close, high, low)
	}

	return nil
}

func getHeikinAshiDelta(c *Candlestick) *HeikinAshiDelta {
	if c == nil || c.Indicators == nil {
		return nil
	}
	return c.Indicators.HeikinAshi
}

func setHeikinAshiDelta(c *Candlestick, open, close, high, low float64) {
	if c.Indicators == nil {
		c.Indicators = &Indicators{}
	}
	if c.Indicators.HeikinAshi == nil {
		c.Indicators.HeikinAshi = &HeikinAshiDelta{
			Open:  open,
			Close: close,
			High:  high,
			Low:   low,
		}
	} else {
		c.Indicators.HeikinAshi.Open = open
		c.Indicators.HeikinAshi.Close = close
		c.Indicators.HeikinAshi.High = high
		c.Indicators.HeikinAshi.Low = low
	}
}
