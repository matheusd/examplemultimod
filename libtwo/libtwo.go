package libtwo

import (
	"github.com/matheusd/examplemultimod/libone/v4"
)

func Name() string {
	return "libtwo"
}

func Version() string {
	return "0.0.1"
}

func Deps() string {
	return libone.Name() + " v" + libone.Version()
}
