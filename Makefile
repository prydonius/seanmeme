build:
	docker run --rm -v "${PWD}":/go/src/seanmeme -w /go/src/seanmeme -e CGO_ENABLED=0 golang:1.8 go build

docker-image: build
	docker build -t prydonius/seanmeme .
