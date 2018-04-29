package shingo

import (
	"fmt"
	"math"
)

// AppendIchimokuCloud appends ichimoku cloud indicator to candlestick
func (cs *Candlesticks) AppendIchimokuCloud(arg IndicatorInputArg) (err error) {
	limit := arg.Limit
	tenkan := arg.IchimokuCloudTenkan
	kijun := arg.IchimokuCloudKijun
	// sb := arg.IchimokuCloudSenkouB
	// chik := arg.IchimokuCloudChikou
	if limit < 1 {
		limit = cs.Total()
	}
	max := math.Max(float64(tenkan), float64(kijun))
	cl := cs.Total()
	startCalcIdx := int(math.Max(float64(cl)-max-float64(limit), 0.0))
	startIndicatorIdx := int(math.Max(float64(cl-limit), 0.0))
	var tenkanHIdx, kijunHIdx int = -1, -1
	var tenkanHVal, kijunHVal float64
	var tenkanLIdx, kijunLIdx int = -1, -1
	var tenkanLVal, kijunLVal float64 = math.Inf(+1), math.Inf(+1)

	for i := startCalcIdx; i < cl; i++ {
		v := cs.ItemAtIndex(i)
		if v.High > tenkanHVal {
			tenkanHVal = v.High
			tenkanHIdx = i
		}
		if v.Low < tenkanLVal {
			tenkanLVal = v.Low
			tenkanLIdx = i
		}
		if v.High > kijunHVal {
			kijunHVal = v.High
			kijunHIdx = i
		}
		if v.Low < kijunLVal {
			kijunLVal = v.Low
			kijunLIdx = i
		}

		// not reaching enough periods to determine tenkan
		if i-startCalcIdx < tenkan-1 {
			continue
		}

		// append tenkan
		if v.High > tenkanHVal {
			tenkanHVal = v.High
			tenkanHIdx = i
		}
		if v.High > kijunHVal {
			kijunHVal = v.High
			kijunHIdx = i
		}
		if v.Low < tenkanLVal {
			tenkanLVal = v.Low
			tenkanLIdx = i
		}
		if v.Low < kijunLVal {
			kijunLVal = v.Low
			kijunLIdx = i
		}

		// make sure the highest/lowest value
		// is within the range
		if tenkanHIdx < i-tenkan+1 || tenkanLIdx < i-tenkan+1 {
			// we have a problem. we'd need to re-calculate the highest value and index
			tenkanHVal = 0
			tenkanLVal = math.Inf(+1)
			for j := (i - tenkan + 1); j <= i; j++ {
				vj := cs.ItemAtIndex(j)
				if vj.High > tenkanHVal {
					tenkanHVal = vj.High
					tenkanHIdx = j
				}
				if vj.Low < tenkanLVal {
					tenkanLVal = vj.Low
					tenkanLIdx = j
				}
			}
		}

		tenkanVal := (tenkanHVal + tenkanLVal) / 2.0

		if i-startCalcIdx < kijun-1 {
			// can't append kijun value yet just append tenkan and move on
			if i < startIndicatorIdx {
				continue
			}
			appendIchimokuCloudIndicator(v, tenkan, kijun, tenkanVal, 0)
			continue
		}

		if kijunHIdx < i-kijun+1 || kijunLIdx < i-kijun+1 {
			kijunHVal = 0
			kijunLVal = math.Inf(+1)
			for j := (i - kijun + 1); j < i; j++ {
				vj := cs.ItemAtIndex(j)
				if vj.High > kijunHVal {
					kijunHVal = vj.High
					kijunHIdx = j
				}
				if vj.Low < kijunLVal {
					kijunLVal = vj.Low
					kijunLIdx = j
				}
			}
		}

		kijunVal := (kijunHVal + kijunLVal) / 2.0
		if i < startIndicatorIdx {
			continue
		}
		appendIchimokuCloudIndicator(v, tenkan, kijun, tenkanVal, kijunVal)
	}
	return
}

func appendIchimokuCloudIndicator(v *Candlestick, t int, k int, tv float64, kv float64) {
	if v.Indicators == nil {
		v.Indicators = &Indicators{}
	}
	key := fmt.Sprintf("%d,%d", t, k)
	if v.Indicators.IchimokuClouds == nil {
		v.Indicators.IchimokuClouds = make(map[string]*IchimokuCloudDelta)
	}
	if v.Indicators.IchimokuClouds[key] == nil {
		v.Indicators.IchimokuClouds[key] = &IchimokuCloudDelta{
			Tenkan:  tv,
			Kijun:   kv,
			SenkouA: (tv + kv) / 2,
		}
	}
}
