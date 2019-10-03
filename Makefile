.PHONY: test
test:
	@$(MAKE) -C testdata
	go test

.PHONY: clean
clean:
	@$(MAKE) -C testdata clean


