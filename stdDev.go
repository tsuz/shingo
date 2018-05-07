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
		avg := v.Indicators.SMAs[p].Value
		var total float64
		for j := 0; j < p; j++ {
			pv := cs.ItemAtIndex(i - j)
			diff := avg - pv.Close
			sqrd := math.Pow(diff, 2)
			total += sqrd
		}
		devSqAvg := total / float64(p)
		sqrt := math.Sqrt(devSqAvg)
		if v.Indicators == nil {
			v.Indicators = &Indicators{}
		}
		if v.Indicators.StdDevs == nil {
			v.Indicators.StdDevs = make(map[int]*StdDevDelta)
		}
		v.Indicators.StdDevs[p] = &StdDevDelta{Value: sqrt}
	}
	return nil
}
