package main

import (
	"flag"
	"fmt"
)

func main() {
	mode := flag.String("o", "w", "output format (o = oscillogram, s = spectogram, w = wav file)")
	time := flag.Int("t", 10, "length of output (in seconds)")
	name := flag.String("f", "", "the output filename")

	flag.Parse()

	if *name == "" {
		flagError("Invalid filename")
		return
	}

	f, err := parseFunction(flag.Args())
	if err != nil {
		flagError("Invalid function: " + err.Error())
		return
	}

	switch *mode {
	case "o":
		outputOscillogram(*name, *time, f)
	case "w":
		outputWav(*name, *time, f)
	case "s":
		outputSpectogram(*name, *time, f)
	default:
		flagError("Invalid mode")
		return
	}
}

func flagError(msg string) {
	fmt.Println(msg)
	flag.Usage()
}
