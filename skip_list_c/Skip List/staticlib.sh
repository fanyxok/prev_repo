#!/bin/sh
gcc -c skiplist.c
ar -r liblist.a skiplist.o
gcc test.c -llist -L. -static -o staticlist 
