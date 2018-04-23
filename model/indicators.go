package model

// Indicators provides indicator type for each candlestick
type Indicators struct {
	SMAs map[int]*SMADelta
	EMAs map[int]*EMADelta
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
