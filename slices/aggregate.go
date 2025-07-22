package slices

import (
	"errors"
	"github.com/windflowed/Generics_Tools"
	"slices"
)

// Max 返回最大值
func Max[T Generics_Tools.RealNumber](sli []T) (T, error) {
	var res T
	if len(sli) == 0 {
		return res, errors.New("sli is empty")
	}
	res = sli[0]
	for i := 1; i < len(sli); i++ {
		if sli[i] > res {
			res = sli[i]
		}
	}
	return res, nil
}

// Min 返回最小值
func Min[T Generics_Tools.RealNumber](sli []T) (T, error) {
	var res T
	if len(sli) == 0 {
		return res, errors.New("sli is empty")
	}
	res = sli[0]
	for i := 1; i < len(sli); i++ {
		if sli[i] < res {
			res = sli[i]
		}
	}
	return res, nil
}

// Sum 返回总和
func Sum[T Generics_Tools.Number](sli []T) (T, error) {
	var res T
	if len(sli) == 0 {
		return res, errors.New("sli is empty")
	}
	for _, v := range sli {
		res += v
	}
	return res, nil
}

// Avg 返回平均值
func Avg[T Generics_Tools.RealNumber](sli []T) (T, error) {
	var res T
	if len(sli) == 0 {
		return res, errors.New("sli is empty")
	}
	sum, err := Sum(sli)
	if err != nil {
		return res, err
	}
	lens := len(sli)
	return sum / T(lens), nil
}

// Count_Distinct 返回不重复元素的个数
func Count_Distinct[T Generics_Tools.RealNumber](sli []T) (int, error) {
	var res int
	if len(sli) == 0 {
		return res, errors.New("sli is empty")
	} else if len(sli) == 1 {
		return 1, nil
	}
	slices.Sort(sli)
	for i := 1; i < len(sli); i++ {
		if sli[i] != sli[i-1] {
			res++
		}
	}
	return res, nil
}
