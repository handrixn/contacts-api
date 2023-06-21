.PHONY: build run clean

build:
	make clean
	go build -o contact-api

run:
	./contact-api

clean:
	rm -f contact-api
