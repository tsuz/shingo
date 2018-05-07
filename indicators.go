package shingo

// Indicators provides indicator type for each candlestick
type Indicators struct {
	SMAs           map[int]*SMADelta
	EMAs           map[int]*EMADelta
	MACDs          map[int]map[int]map[int]*MACDDelta
	IchimokuClouds map[string]*IchimokuCloudDelta
	ATRs           map[int]*ATRDelta
	SuperTrends    map[int]map[float64]*SuperTrendDelta
	HeikinAshi     *HeikinAshiDelta
	StdDevs        map[int]*StdDevDelta
}

// Get provides indicator query interface
func (in *Indicators) Get(arg IndicatorInputArg) interface{} {
	p := arg.Period
	t := arg.Type
	switch t {
	case IndicatorTypeStdDev:
		if in.StdDevs == nil {
			return nil
		}
		return in.StdDevs[p]
	}
	return nil
}

// SMADelta is the value for this period and change since last period
type SMADelta struct {
	Value  float64
	Change float64
}

// EMADelta is the value for this period and change since last period
type EMADelta struct {
	Value  float64
	Change float64
}

// MACDDelta provides macd, signal and histogram for this candlestick
type MACDDelta struct {
	MACDValue     float64
	SignalValue   float64
	MACDHistogram float64
}

// IchimokuCloudDelta provides ichimoku cloud indicator for this candlestick
type IchimokuCloudDelta struct {
	Tenkan  float64
	Kijun   float64
	SenkouA float64
	SenkouB float64
	Chikou  float64
}

// ATRDelta provides average true range for this candlestick
type ATRDelta struct {
	Value  float64
	Change float64
}

type SuperTrendDelta struct {
	Longband  float64
	Shortband float64
	Trend     Trend
}

type HeikinAshiDelta struct {
	Open  float64
	High  float64
	Low   float64
	Close float64
}

type StdDevDelta struct {
	Value float64
}

type Trend int

const (
	// Undeterminable trend
	Undeterminable Trend = 0
	// Bear market
	Bear Trend = 1
	// Bull market
	Bull Trend = 2
)
