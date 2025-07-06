package parser

import (
	"reflect"
	"regexp"
)

// AppendManyNew creates a new slice by appending multiple values or slices.
func AppendManyNew[T any](v []T, as ...any) []T {
	l := len(v)
	for _, a := range as {
		switch a := a.(type) {
		case T:
			l++
		case []T:
			l += len(a)
		default:
			panic("type error:" + reflect.TypeOf(a).String())
		}
	}
	n := make([]T, l)
	n = n[:0]
	n = append(n, v...)
	for _, a := range as {
		switch a := a.(type) {
		case T:
			n = append(n, a)
		case []T:
			n = append(n, a...)
		default:
			panic("type error:" + reflect.TypeOf(a).String())
		}
	}
	return n
}

func regexFindAll(s0, frx string) []string {
	return regexp.MustCompile(frx).FindAllString(s0, -1)
}
