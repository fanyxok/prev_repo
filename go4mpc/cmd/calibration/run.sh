#!/bin/bash

# N
# O - file path
# Role -- role
# Addr -- ip addresses
./calibration -N 100 -Role 0 -O ./cost0.json -Addr "127.0.0.1:23344;" &
./calibration -N 100 -Role 1 -O ./cost1.json -Addr "127.0.0.1:23344;" 