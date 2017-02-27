package main

import (
	"flag"
	"github.com/aschwinwester/fsplit/split"
	"log"
)

func main() {

	options := split.GetOptions()

	if len(flag.Args()) == 0 {
		log.Println("provide folder as argument and prefix argument with option flags.")
		return
	}

	var folderLocation = flag.Args()[0]
	split.SplitFolder(options, folderLocation)

}
