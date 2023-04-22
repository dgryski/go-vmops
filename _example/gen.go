package main

import (
	"log"
	"os"
)

func main() {
	b, err := abcvm.MarshalBinary()
	if err != nil {
		log.Fatalf("error marshaling abcvm: %v\n", err)
	}

	const file = "abcbytes.ops"

	if err := os.WriteFile(file, b, 0644); err != nil {
		log.Fatalf("error writeing %v: %v", file, err)
	}
}
