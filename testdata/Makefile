.PHONY: all
all: bin/hello-bundled bin/hello-zipslip

bin/hello-bundled: bin/hello zip/assets.zip
	cat bin/hello zip/assets.zip > bin/hello-bundled
	chmod +x bin/hello-bundled

bin/hello-zipslip: bin/hello zip/zipslip.zip
	cat bin/hello zip/zipslip.zip > bin/hello-zipslip
	chmod +x bin/hello-zipslip

zip/assets.zip: zip
	cd assets && zip -vr ../zip/assets.zip .

zip/zipslip.zip: zip
	cd zip && zip -vr zipslip.zip ../assets

bin/hello: bin
	go build -o bin/hello

zip:
	mkdir -p zip

bin:
	mkdir -p bin

.PHONY: clean
clean:
	rm -rf bin
	rm -rf zip

