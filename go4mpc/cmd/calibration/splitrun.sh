#!/bin/bash
 go build -ldflags "-s -w"
 
/bin/sh -ec './calibration -N 100 -Role 0 -O ./ -Addr "127.0.0.1:23344;"' &
sleep 0.1s 
/bin/sh -ec './calibration -N 100 -Role 1 -O ./ -Addr "127.0.0.1:23344;"'