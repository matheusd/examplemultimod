module github.com/matheusd/examplemultimod

go 1.19

require (
	github.com/matheusd/examplemultimod/libone v0.0.0-00000000000000-000000000000
	github.com/matheusd/examplemultimod/libtwo v0.0.0-00000000000000-000000000000
)

replace (
	github.com/matheusd/examplemultimod/libone => ./libone
	github.com/matheusd/examplemultimod/libtwo => ./libtwo
)
