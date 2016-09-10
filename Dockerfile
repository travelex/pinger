FROM scratch
ADD pinger /
ENV TARGET_URL  http://localhost
ENV METHOD      POST
ENV INTERVAL    60
ADD zoneinfo.tar.gz /
ENTRYPOINT ["/pinger"]
