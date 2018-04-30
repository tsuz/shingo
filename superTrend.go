package shingo

import (
	"math"

	"github.com/pkg/errors"
)

// AppendSuperTrend generates super trend delta for candlestick
func (cs *Candlesticks) AppendSuperTrend(arg IndicatorInputArg) error {
	limit := arg.Limit
	period := arg.Period
	multi := arg.Multiplier
	total := cs.Total()
	if period < 1 {
		return errors.New("expected period to be positive")
	}
	if multi <= 0.0 {
		return errors.New("multiplier must be greater than zero")
	}
	if total < 1 {
		return nil
	}
	if limit < 1 {
		limit = total
	}

	if err := cs.AppendATR(arg); err != nil {
		return errors.New("error appending ATR for super trend")
	}

	startIdx := total - 1 - limit - period
	if startIdx < 0 {
		startIdx = 0
	}

	var count int
	for i := startIdx; i < total; i++ {
		count++
		if count < period {
			continue
		}

		p := cs.ItemAtIndex(i - 1)
		pst := getSuperTrend(p, period, multi)
		v := cs.ItemAtIndex(i)
		atr := getATR(v, arg.Period)
		if atr == nil {
			continue
		}

		nshortb := (v.High+v.Low)/2.0 + atr.Value*multi
		nlongb := (v.High+v.Low)/2.0 - atr.Value*multi

		var long float64
		var short float64
		var trend Trend

		if pst != nil {
			if p.Close > pst.Longband {
				long = math.Max(nlongb, pst.Longband)
			} else {
				long = nlongb
			}

			if p.Close < pst.Shortband {
				short = math.Min(nshortb, pst.Shortband)
			} else {
				short = nshortb
			}

			if v.Close > pst.Shortband {
				trend = Bull
			} else if v.Close < pst.Longband {
				trend = Bear
			} else {
				trend = pst.Trend
			}
		} else {
			long = nlongb
			short = nshortb
			trend = Undeterminable
		}

		setSuperTrend(v, period, multi, short, long, trend)
	}

	return nil
}

func getSuperTrend(c *Candlestick, period int, multi float64) *SuperTrendDelta {
	if c == nil {
		return nil
	}
	if c.Indicators == nil {
		return nil
	}
	if c.Indicators.SuperTrends == nil {
		return nil
	}
	if c.Indicators.SuperTrends[period] == nil {
		return nil
	}
	return c.Indicators.SuperTrends[period][multi]
}

func setSuperTrend(c *Candlestick, period int, multi float64, short float64, long float64, trend Trend) {
	if c == nil {
		return
	}
	if c.Indicators == nil {
		c.Indicators = &Indicators{}
	}
	if c.Indicators.SuperTrends == nil {
		c.Indicators.SuperTrends = make(map[int]map[float64]*SuperTrendDelta)
	}
	if c.Indicators.SuperTrends[period] == nil {
		c.Indicators.SuperTrends[period] = make(map[float64]*SuperTrendDelta)
	}
	c.Indicators.SuperTrends[period][multi] = &SuperTrendDelta{
		Shortband: short,
		Longband:  long,
		Trend:     trend,
	}
}
