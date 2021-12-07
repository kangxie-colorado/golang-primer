module github.com/kangxie-colorado/golang-primer/messaging/test

go 1.16

// terrible naming here
replace github.com/kangxie-colorado/golang-primer/messaging/lib => ../lib

replace github.com/kangxie-colorado/golang-primer/messaging/test/lib => ./lib

require (
	github.com/kangxie-colorado/golang-primer/messaging/lib v0.0.0-00010101000000-000000000000
	github.com/kangxie-colorado/golang-primer/messaging/test/lib v0.0.0-00010101000000-000000000000 // indirect
	github.com/sirupsen/logrus v1.8.1
)
