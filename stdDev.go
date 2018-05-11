package shingo

import (
	"math"

	"github.com/pkg/errors"
)

// AppendStdDev appends standard deviation data to each candlestick
func (cs *Candlesticks) AppendStdDev(arg IndicatorInputArg) error {

	p := arg.Period
	l := arg.Limit
	t := cs.Total()
	if t < 1 {
		return nil
	}
	if l < 1 {
		l = t
	}
	if p < 1 {
		return errors.New("Period must be larger than zero")
	}

	if err := cs.AppendSMA(IndicatorInputArg{Period: p, Limit: l}); err != nil {
		return errors.Wrap(err, "Error appending sma")
	}

	startIdx := (t - 1) - p - l
	if startIdx < 0 {
		startIdx = 0
	}
	for i := startIdx; i < t; i++ {
		v := cs.ItemAtIndex(i)

		if i < (p - 1) {
			// not enough data to generate average
			continue
		}
		avg := v.GetSMA(p).Value
		var total float64
		for j := 0; j < p; j++ {
			pv := cs.ItemAtIndex(i - j)
			diff := avg - pv.Close
			sqrd := math.Pow(diff, 2)
			total += sqrd
		}
		devSqAvg := total / float64(p)
		sqrt := math.Sqrt(devSqAvg)
		v.setStdDev(p, sqrt)
	}
	return nil
}

// GetStdDev gets standard deviation value for this candlestick for given lookback period
func (c *Candlestick) GetStdDev(period int) *StdDevDelta {
	if c.Indicators == nil || c.Indicators.StdDevs == nil {
		return nil
	}
	return c.Indicators.StdDevs[period]
}

func (c *Candlestick) setStdDev(p int, sqrt float64) {
	if c.Indicators == nil {
		c.Indicators = &Indicators{}
	}
	if c.Indicators.StdDevs == nil {
		c.Indicators.StdDevs = make(map[int]*StdDevDelta)
	}
	c.Indicators.StdDevs[p] = &StdDevDelta{Value: sqrt}
}
