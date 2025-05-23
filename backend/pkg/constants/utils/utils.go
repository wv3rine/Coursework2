package utils

import (
	"fmt"
)

func MapArr[I, O any](arr []I, mapper func(in I) O) []O {
	res := make([]O, 0, len(arr))
	for _, v := range arr {
		res = append(res, mapper(v))
	}
	return res
}

func MapPointerVals[I, O any](num *I, mapper func(in I) O) *O {
	if num == nil {
		return nil
	}
	o := mapper(*num)
	return &o
}

func WithPrefix(slice []string, prefix string) []string {
	out := make([]string, 0, len(slice))
	for _, v := range slice {
		out = append(out, fmt.Sprintf("%s.%s", prefix, v))
	}
	return out
}
