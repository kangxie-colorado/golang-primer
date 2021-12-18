module github.com/kangxie-colorado/golang-primer/messaging/test

go 1.16

// terrible naming here
replace github.com/kangxie-colorado/golang-primer/messaging/lib => ../lib

replace github.com/kangxie-colorado/golang-primer/messaging/test/lib => ./lib

replace github.com/kangxie-colorado/golang-primer/messaging/test/traffic_light => ./traffic_light

require (
	github.com/kangxie-colorado/golang-primer/messaging/lib v0.0.0-20211217002829-8e7d9237e95a
	github.com/kangxie-colorado/golang-primer/messaging/test/lib v0.0.0-00010101000000-000000000000
	github.com/kangxie-colorado/golang-primer/messaging/test/traffic_light v0.0.0-00010101000000-000000000000
	github.com/sirupsen/logrus v1.8.1
)
