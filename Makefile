default: build

build: clean
		 docker build -f build_Dockerfile -t build:latest . \
	&& docker run --rm -v "`pwd`:/gopath/src/github.com/johnpeterharvey/pinger" build:latest \
	&& docker build -t pinger:latest .

run:
	docker run pinger

clean:
	 rm pinger \
&& docker rmi -f pinger:latest || echo "Pinger already removed" \
&& docker rmi -f build:latest || echo "Build image already removed"

.PHONY: build run
