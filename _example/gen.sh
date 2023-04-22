#!/bin/sh

# create the  Go version of the machine
re -r pcre -k pair -C -e abc -pl vmops_c 'ab*c' |awk -f ../scripts/vmops2go.awk -v package=main -v var=abcvm >abcvm.go

# Turn the Go version into bytes
go run gen.go abcvm.go

# load and the run the bytes via embed
go run match.go
