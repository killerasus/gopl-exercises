package minmax

import "math"

func Max(list ...int) (int, bool) {
	ret := math.MinInt
	if len(list) == 0 {
		return ret, false
	}
	for _, v := range list {
		if v > ret {
			ret = v
		}
	}
	return ret, true
}

func Min(list ...int) (int, bool) {
	ret := math.MaxInt
	if len(list) == 0 {
		return ret, false
	}
	for _, v := range list {
		if v < ret {
			ret = v
		}
	}
	return ret, true
}
