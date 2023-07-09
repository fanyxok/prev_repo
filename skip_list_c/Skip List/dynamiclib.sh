#!/bin/sh
gcc -fPIC -shared -o liblist.so skiplist.c 
gcc test.c -o dynamiclist ./liblist.so
