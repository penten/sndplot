package main

import (
	"flag"
	"fmt"
)

func main() {
	mode := flag.String("o", "w", "output format (o = oscillogram, s = spectogram, w = wav file)")
	name := flag.String("f", "", "the output filename")

	flag.Parse()

	if *name == "" {
		fmt.Println("Invalid filename")
		flag.Usage()
		return
	}

	switch *mode {
	case "o":
	case "w":
	case "s":
		return
	default:
		fmt.Println("Invalid mode")
		flag.Usage()
		return
	}
}
