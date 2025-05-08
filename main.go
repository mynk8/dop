package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: fob [options] [name]\n")
	flag.PrintDefaults()
	os.Exit(2)
}

var (
	watchWithTime = flag.Bool("wt", false, "Watch the file output with timestamp of each line")
	watchOutput = flag.String("g", "Hello", "Greet with `greeting`")
	tailOutput = flag.Bool("r", false, "Greet in Reverse")
	headOutput = flag.Bool("h", false, "Print the first 5 lines")
)

func main() {
	log.SetFlags(0)
	log.SetPrefix("hello: ")

	flag.Usage = usage
	flag.Parse()
	
	var filename string
	args := flag.Args()

	if len(args) >= 2 {
		usage()
	}

	if len(args) >= 1 {
		filename = args[0]
	}

	if filename == "" {
		log.Fatalf("invalid name %q", filename)
	}

}
