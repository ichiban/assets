.PHONY: all
all: bin/bundled bin/zipslip

bin/bundled: bin/main zip/assets.zip
	cat bin/main zip/assets.zip > bin/bundled
	chmod +x bin/bundled

bin/zipslip: bin/main zip/zipslip.zip
	cat bin/main zip/zipslip.zip > bin/zipslip
	chmod +x bin/zipslip

zip/assets.zip:
	cd assets && zip -vr ../zip/assets.zip .

zip/zipslip.zip:
	cd zip && zip -vr zipslip.zip ../assets

bin/main:
	go build -o bin/main

.PHONY: clean
clean:
	rm -f bin/main
	rm -f zip/assets.zip
	rm -f zip/zipslip.zip
	rm -f bin/bundled
	rm -f bin/zipslip

