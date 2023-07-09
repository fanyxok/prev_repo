#!/bin/bash
go build -ldflags "-s -w" ./cryptonet/4_codegen/0

myArray=()
for item in {1..86410}; do 
    myArray+=('1')
done

echo "${myArray[@]}" | ./0 

