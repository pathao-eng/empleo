#!/bin/bash

#go build -buildmode=plugin -o bin/tiger.so plugins/tiger.go

src=./sources/
bin=./bin/

for s in $(ls $src)
do
  go build -v -buildmode=plugin -o $bin$s.so $src$s
done
