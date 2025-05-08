package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"dop.org/dop/head"
	"dop.org/dop/tail"
	"dop.org/dop/watch"
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: fob [options] [name]\n")
	flag.PrintDefaults()
	os.Exit(2)
}

var (
	watchWithTime = flag.Bool("wt", false, "Watch the file output with timestamp of each line")
	tailOutput    = flag.Bool("r", false, "Print the last 5 lines")
	headOutput    = flag.Bool("h", false, "Print the first 5 lines")
)

func main() {
	log.SetFlags(0)
	log.SetPrefix("hello: ")

	flag.Usage = usage
	flag.Parse()

	args := flag.Args()
	if len(args) != 1 {
		fmt.Fprintf(os.Stderr, "ERROR: must supply atleast one filename\n")
		usage()
	}
	filename := args[0]

	if *watchWithTime {
		watch.WatchWithTimestamps(filename)
	}

	if *tailOutput {
		tail.Tail(filename, 5)
	}

	if *headOutput {
		head.Head(filename, 5)
	}

	fmt.Printf("%s <-- arguments supplied\n", filename)
}



