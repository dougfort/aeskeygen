package main

import (
	"crypto/rand"
	"encoding/base64"
	"flag"
	"log"
	"os"
)

func main() {
	os.Exit(run())
}

func run() int {
	var size int

	flag.IntVar(&size, "size", 0, "key size. Must be 16, 24, or 32")
	flag.Parse()

	if size != 16 && size != 24 && size != 32 {
		log.Printf("Invalid size: %d; must be 16, 24, or 32", size)
		return 2
	}

	key := make([]byte, size)
	if _, err := rand.Read(key); err != nil {
		log.Printf("rand.Read failed: %v", err)
		return 2
	}

	println(base64.StdEncoding.EncodeToString(key))

	return 0
}
