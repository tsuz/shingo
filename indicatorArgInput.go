package shingo

type IndicatorInputArg struct {
	Type IndicatorType

	Limit      int
	Period     int
	Multiplier float64

	MacdLarge  int
	MacdSmall  int
	MacdSignal int

	IchimokuCloudTenkan  int
	IchimokuCloudKijun   int
	IchimokuCloudSenkouB int
	IchimokuCloudChikou  int
}
