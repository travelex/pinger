FROM scratch
ENV TARGET_URL  http://localhost
ENV METHOD      POST
ENV INTERVAL    60
ADD zoneinfo.zip /lib/time/zoneinfo.zip
ADD pinger       /
ENTRYPOINT ["/pinger"]
