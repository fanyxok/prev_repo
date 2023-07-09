#!/bin/bash

s="Get /s/b HTTP/1.1"
c=$(echo ${s}|grep -i "Get"|awk '{print $2}')
echo $c