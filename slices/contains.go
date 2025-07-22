package slices

// Contains 应该传入可比较类型的参数,判断 slice 中是否存在 value 这个元素
func Contains[T comparable](slice []T, value T) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

// ContainsAll 应该传入可比较类型的参数,判断 slice 中是否存在 value 中的所有元素
func ContainsAll[T comparable](slice []T, value []T) bool {
	for _, v := range value {
		singleBool := false
		for _, s := range slice {
			if v == s {
				singleBool = true
				break
			}
		}
		if !singleBool {
			return false
		}
	}
	return true
}

// ContainsAny 应该传入可比较类型的参数,判断 slice 中是否存在 value 中的任意一个元素
func ContainsAny[T comparable](slice []T, value []T) bool {
	for _, v := range value {
		singleBool := false
		for _, s := range slice {
			if v == s {
				singleBool = true
				break
			}
		}
		if singleBool {
			return true
		}
	}
	return false
}
