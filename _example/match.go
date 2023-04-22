package main

import (
	_ "embed"
	"fmt"

	"github.com/dgryski/go-vmops"
)

//go:embed abcbytes.ops
var abcbytes []byte

func main() {
	var abc vmops.VM
	abc.UnmarshalBinaryUnsafe(abcbytes)
	fmt.Println(abc.MatchString("abbbc"))
}
