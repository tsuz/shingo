package model

// Indicators provides indicator type for each candlestick
type Indicators struct {
	SMAs map[int]*SMADelta
}

// SMADelta is value for this period and change since last period
type SMADelta struct {
	Value  float64
	Change float64
}
