# pinger

[![Build Status Develop](https://travis-ci.org/johnpeterharvey/pinger.svg?branch=develop)](https://travis-ci.org/johnpeterharvey/pinger)
[![Build Status Master](https://travis-ci.org/johnpeterharvey/pinger.svg?branch=master)](https://travis-ci.org/johnpeterharvey/pinger)

Builds a small (~5MB) Docker container that calls out to an HTTP endpoint on an interval.


Used for when your scalable stateless microservice needs to use an async third-party API that is missing callbacks.

Edit the Dockerfile to set the environment variables:

  * TARGET_URL
  * METHOD (e.g. POST)
  * BODY (e.g. {})
  * INTERVAL (in seconds)
  * TIME (optional, e.g. 16:37:00)
  * TIMEZONE (optional, e.g. Europe/London)

Then run

    make build

The resulting container is

    pinger:latest

Plug this container into your composure, with a scale factor of 1.

The single instance of pinger will call out to your microservice endpoint, directed by your load-balancer to one instance of the service, which can then do whatever logic is needed to handle state-change on your 3rd party.

Currently if no body is supplied an empty JSON body ```{}``` is sent.
