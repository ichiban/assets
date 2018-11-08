.PHONY: test
test:
	go get -t ./...
	@$(MAKE) -C testdata
	go test

.PHONY: clean
clean:
	@$(MAKE) -C testdata clean


