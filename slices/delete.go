package slices

import "errors"

// Delete 传入切片以及要删除的元素的索引，返回删除后的切片
func Delete[T any](sli []T, index int) ([]T, error) {
	if len(sli) == 0 {
		return nil, errors.New("slice is empty")
	} else if index < 0 || index >= len(sli) {
		return sli, errors.New("index out of range")
	} else {
		return append(sli[:index], sli[index+1:]...), nil
	}
}
