BINARY=eksdoctor

build:
	go build -o $(BINARY)

install: build
	mv $(BINARY) /usr/local/bin/$(BINARY)

clean:
	rm -f $(BINARY)
