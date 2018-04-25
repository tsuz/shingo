package shingo

import (
	"fmt"
)

type CandlestickMeta interface {
	Total() int
	Interval() Interval
	ItemAtIndex(idx int) *Candlestick
}
type CandlestickIndicators interface {
	AppendEMA(arg IndicatorInputArg) error
	AppendSMA(arg IndicatorInputArg) error
	AppendMACD(arg IndicatorInputArg) error
	AppendIchimokuCloud(arg IndicatorInputArg) error
}

type Candlesticks struct {
	CandlestickMeta
	CandlestickIndicators

	interval Interval
	items    []*Candlestick
}

func (cs *Candlesticks) AppendCandlestick(c *Candlestick) error {
	cs.items = append(cs.items, c)
	return nil
}

func NewCandlesticks(i Interval) (*Candlesticks, error) {
	return &Candlesticks{
		interval: i,
		items:    make([]*Candlestick, 0),
	}, nil
}

// Total returns total candlesticks
func (cs *Candlesticks) Total() int {
	return len(cs.items)
}

// ItemAtIndex returns the item at specific index
func (cs *Candlesticks) ItemAtIndex(idx int) *Candlestick {
	if len(cs.items) > idx {
		return cs.items[idx]
	}
	return nil
}

// Interval returns currently set interval for the series of candlesticks
func (cs *Candlesticks) Interval() Interval {
	return cs.interval
}

// GenerateIndicator generates requested signals on that series of candlesticks
func (cs *Candlesticks) GenerateIndicator(i IndicatorType, arg IndicatorInputArg) error {
	switch i {
	case IndicatorTypeSMA:
		return cs.AppendSMA(arg)
	case IndicatorTypeEMA:
		return cs.AppendEMA(arg)
	case IndicatorTypeMACD:
		return cs.AppendMACD(arg)
	case IndicatorTypeIchimokuCloud:
		return cs.AppendIchimokuCloud(arg)
	}
	return fmt.Errorf("Error unsupported indicator type %+v", i)
}
