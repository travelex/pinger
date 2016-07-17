FROM scratch
COPY bin/pinger /
ENV TARGET_URL  http://localhost
ENV METHOD      POST
ENV INTERVAL    60
ENTRYPOINT ["/pinger"]
