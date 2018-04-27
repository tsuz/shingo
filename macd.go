package shingo

import (
	"math"
	"time"

	"github.com/pkg/errors"
)

// AppendMACD appends Moving Average Convergence Divergence indicators to each candlestick
func (cs *Candlesticks) AppendMACD(args IndicatorInputArg) error {
	period1 := args.MacdLarge
	period2 := args.MacdSmall
	signalLine := args.MacdSignal
	limit := args.Limit

	if period1 == 0 {
		return errors.New("MacdLarge must be greater than zero")
	}
	if period2 == 0 {
		return errors.New("MacdSmall must be greater than zero")
	}
	if signalLine == 0 {
		return errors.New("signalLine must be greater than zero")
	}
	if period1 >= period2 {
		return errors.New("Period1 must be less than Period2 in MACD")
	}

	cl := cs.Total()
	if cl < 1 {
		return nil
	}
	if limit < 1 {
		limit = cs.Total()
	}

	if err := cs.GenerateIndicator(IndicatorTypeEMA, IndicatorInputArg{
		Period: period1,
		Limit:  cl,
	}); err != nil {
		return errors.Wrap(err, "Error generating period1 indicator")
	}

	if err := cs.GenerateIndicator(IndicatorTypeEMA, IndicatorInputArg{
		Period: period2,
		Limit:  cl,
	}); err != nil {
		return errors.Wrap(err, "Error generating period2 indicator")
	}

	cst, err := NewCandlesticks(IntervalOneDay, 100)
	if err != nil {
		return errors.Wrap(err, "Error creating candlesticks for macd signal line")
	}
	for i := 0; i < cl; i++ {
		v := cs.ItemAtIndex(i)
		if v.Indicators == nil {
			continue
		}
		if v.Indicators.EMAs[period1] == nil {
			continue
		}
		if v.Indicators.EMAs[period2] == nil {
			continue
		}
		val := v.Indicators.EMAs[period1].Value - v.Indicators.EMAs[period2].Value
		c, err := NewCandlestick(0, val, 0, 0, time.Time{}, 0)
		if err != nil {
			return errors.Wrap(err, "Error creating candlestick in macd signal line")
		}
		cst.AppendCandlestick(c)
	}

	cstl := cst.Total()
	err = cst.GenerateIndicator(IndicatorTypeEMA, IndicatorInputArg{
		Period: signalLine,
		Limit:  cl,
	})
	if err != nil {
		return errors.Wrap(err, "Error creating ema for macd signal line")
	}
	endIdx := cl - cstl
	var count int
	for i := cl - 1; i >= endIdx; i-- {
		if count == limit {
			return nil
		}

		v := cs.ItemAtIndex(i)

		ci := int(math.Abs(float64(cl - i - cstl)))

		vi := cst.ItemAtIndex(ci)
		if vi != nil && (vi.Indicators == nil || vi.Indicators.EMAs[signalLine] == nil) {
			continue
		}

		macdValue := v.Indicators.EMAs[period1].Value - v.Indicators.EMAs[period2].Value

		signalValue := vi.Indicators.EMAs[signalLine].Value

		setMACDIndicator(v, period1, period2, signalLine, macdValue, signalValue)

		count++
	}
	return nil
}

func setMACDIndicator(v *Candlestick, period1 int, period2 int, signal int, macdValue float64, signalValue float64) {
	if v.Indicators.MACDs == nil {
		v.Indicators.MACDs = make(map[int]map[int]map[int]*MACDDelta)
	}
	if v.Indicators.MACDs[period1] == nil {
		v.Indicators.MACDs[period1] = make(map[int]map[int]*MACDDelta)
	}
	if v.Indicators.MACDs[period1][period2] == nil {
		v.Indicators.MACDs[period1][period2] = make(map[int]*MACDDelta)
	}
	if v.Indicators.MACDs[period1][period2][signal] == nil {
		v.Indicators.MACDs[period1][period2][signal] = &MACDDelta{
			MACDValue:     macdValue,
			SignalValue:   signalValue,
			MACDHistogram: macdValue - signalValue,
		}
	}
}
