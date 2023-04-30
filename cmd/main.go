// Command hashify is entry point
package main

import (
	"log"
	"os"

	"github.com/avakarev/hashify/internal/cli"
	"github.com/avakarev/hashify/internal/cryptoutil"
)

func main() {
	cli.Validate()
	for i, f := range cli.Files() {
		if i > 0 {
			log.Println("")
		}
		log.Println(f)
		for _, h := range cli.HashFlags() {
			printChecksum(f, h)
		}
	}
}

func printChecksum(f string, algo string) {
	sum, err := cryptoutil.Checksum(f, algo)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	log.Printf("%6s | %s\n", algo, sum)
}

func init() {
	log.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))
}
