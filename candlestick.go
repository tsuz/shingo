package shingo

import (
	"time"
)

type Candlestick struct {
	OpenTime   time.Time
	Open       float64
	High       float64
	Low        float64
	Close      float64
	Volume     float64
	Indicators *Indicators
}

func NewCandlestick(open float64, close float64, high float64, low float64, ot time.Time, vol float64) (*Candlestick, error) {
	return &Candlestick{
		Open:     open,
		Close:    close,
		High:     high,
		Low:      low,
		OpenTime: ot,
		Volume:   vol,
	}, nil
}
