#!/bin/sh

cp /goroot/lib/time/zoneinfo.zip .
go build -ldflags "-extldflags '-static'" pinger.go
