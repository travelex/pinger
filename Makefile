default: build

build: clean
	docker build -f build_Dockerfile -t build:latest . \
	&& docker run --rm -v "$(CURDIR):/gopath/src/github.com/johnpeterharvey/pinger:rw" build:latest \
	&& tar cfz zoneinfo.tar.gz /usr/share/zoneinfo \
	&& docker build -t pinger:latest . \
	&& rm -f zoneinfo.tar.gz

run:
	docker run pinger

clean:
	rm -f pinger || echo "No executable found" \
	&& docker rmi -f pinger:latest || echo "Pinger already removed" \
	&& docker rmi -f build:latest || echo "Build image already removed"

.PHONY: build run
