package stdlib

import (
	"github.com/mew-sh/tengo/v2"
)

func wrapError(err error) tengo.Object {
	if err == nil {
		return tengo.TrueValue
	}
	return &tengo.Error{Value: &tengo.String{Value: err.Error()}}
}
