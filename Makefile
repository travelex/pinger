default: clean test build

build:
	docker build -f build_Dockerfile -t build:latest . \
	&& docker run --rm -v "$(CURDIR):/gopath/src/github.com/johnpeterharvey/pinger:rw" build:latest \
	&& docker build -t pinger:latest .

run:
	docker run pinger

test:
	go test

clean:
	rm -f pinger || echo "No executable found" \
	&& rm -f zoneinfo.zip || echo "No zoneinfo found" \
	&& docker rmi -f pinger:latest || echo "Pinger already removed" \
	&& docker rmi -f build:latest || echo "Build image already removed"

.PHONY: build run test
