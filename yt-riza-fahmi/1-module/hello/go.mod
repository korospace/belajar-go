module example.com/hello

go 1.20

replace example.com/greetings => ../greetings

require (
	example.com/greetings v0.0.0-00010101000000-000000000000
	github.com/dustinkirkland/golang-petname v0.0.0-20191129215211-8e5a1ed0cff0
)
