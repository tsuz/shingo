package model

type IndicatorInputArg struct {
	Limit  int
	Period int

	MacdLarge  int
	MacdSmall  int
	MacdSignal int

	IchimokuCloudTenkan  int
	IchimokuCloudKijun   int
	IchimokuCloudSenkouB int
	IchimokuCloudChikou  int
}
