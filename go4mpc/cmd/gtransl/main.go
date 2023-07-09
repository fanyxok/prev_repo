package main

import (
	"flag"
	"log"
	"s3l/mpcfgo/pkg/iast"
)

func main() {
	var dir string
	flag.StringVar(&dir, "L", "", "Path of .go file")
	flag.Parse()
	if dir == "" {
		log.Fatal("No -L flag")
	} else {
		log.Printf("Analyze Folder: %v\n", dir)
	}

	if err := iast.DoMain(dir); err != nil {
		log.Panicf("DoMain %v", err)
	}

}
