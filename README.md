# pinger

Builds a small (~5MB) Docker container that calls out to an HTTP endpoint on an interval.


Used for when your scalable stateless microservice needs to use an async third-party API that is missing call-backs.

Edit the Dockerfile to set your endpoint and interval, and run

    make build

The resulting container is

    pinger:latest

Plug this container into your composure, with a scale factor of 1.

The single instance of pinger will call out to your microservice endpoint, directed by your load-balancer to one instance of the service, which can then do whatever logic is needed to handle state-change on your 3rd party.


Currently only supports empty POST requests. Other methods will be added soon.
