package helper

func If[T any](cond bool, trueValue, falseValue T) T {
	if cond {
		return trueValue
	} else {
		return falseValue
	}
}
