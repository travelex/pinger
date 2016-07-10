default: build

build:
	docker run --rm -v "`pwd`:/src" -v /var/run/docker.sock:/var/run/docker.sock centurylink/golang-builder

run:
	docker run pinger

.PHONY: build run
