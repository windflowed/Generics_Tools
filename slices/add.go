package slices

import "errors"

func Add[T any](sli []T, index int, value T) ([]T, error) {
	if index < 0 || index >= len(sli) {
		return nil, errors.New("index is out of range")
	} else {
		sli = append(sli, value)
		for i := len(sli) - 1; i > index; i-- {
			if i > 0 {
				sli[i] = sli[i-1]
			}
		}
		sli[index] = value
		return sli, nil
	}
}
