BASE_PACKAGE = github.com/soramitsu/iroha-go
PACKAGE_LIST := . 

flatbuf:
	flatc --gen-mutable --go --grpc -o . schema/*.fbs
install-go:
	sh util/go.sh 1.8.2
install-glide:
	sh util/glide.sh v0.12.3
deps: install-glide
	glide install
test: 
	@for p in $(PACKAGE_LIST); do \
		echo "==> Unit Testing $$p"; \
		go test -v $(BASE_PACKAGE)/$$p || exit 1; \
	done
