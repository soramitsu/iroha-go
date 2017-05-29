flatbuf:
	flatc --gen-mutable --go --grpc -o . schema/*.fbs
install-go:
	sh util/go.sh 1.8.2
install-glide:
	sh util/glide.sh v0.12.3
deps: install-glide
	glide install
test:
	go test -v -race $(go list ./... | grep -v vendor)
