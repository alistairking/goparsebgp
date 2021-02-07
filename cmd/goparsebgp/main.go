package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/alistairking/goparsebgp"
)

var (
	FLAG_version = flag.Bool("version", false, "Version")
)

func log(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, format+"\n", args...)
}

func printVersion() {
	v := goparsebgp.Version
	fmt.Printf("GoParseBGP v%s (commit: %s)\n", v.GoParseBGP.Version, v.GoParseBGP.Commit)
	fmt.Printf("libParseBGP v%s (commit: %s)\n", v.LibParseBGP.Version, v.LibParseBGP.Commit)
}

func main() {
	flag.Parse()

	if *FLAG_version {
		printVersion()
		os.Exit(0)
	}

	for _, fname := range flag.Args() {
		log("Processing %s", fname)
	}
}
