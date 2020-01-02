.DEFAULT all: build test

build: .glibLastBuild

.glibLastBuild: ./glib/*.go
	go build github.com/rrbrussell/glib-go/glib
	touch .glibLastBuild

test: glibCoverage.html

glibCoverage.html: ./glib/*.go
	go test -cover -coverprofile=glibCoverage.out github.com/rrbrussell/glib-go/glib
	go tool cover -html=glibCoverage.out -o glibCoverage.html
