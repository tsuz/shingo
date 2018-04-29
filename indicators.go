package shingo

// Indicators provides indicator type for each candlestick
type Indicators struct {
	SMAs           map[int]*SMADelta
	EMAs           map[int]*EMADelta
	MACDs          map[int]map[int]map[int]*MACDDelta
	IchimokuClouds map[string]*IchimokuCloudDelta
	ATRs           map[int]*ATRDelta
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
