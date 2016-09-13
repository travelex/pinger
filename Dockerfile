FROM scratch
ADD zoneinfo.zip /usr/local/go/lib/time/zoneinfo.zip
ADD pinger       /
ENTRYPOINT ["/pinger"]
