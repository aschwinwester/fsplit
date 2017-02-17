package main

import (
	"github.com/aschwinwester/fsplit/split"
	"fmt"
	"os"
)

func main() {


	options := split.GetOptions()

	if os.Args[0] == "" {
		fmt.Println("provide folder as argument and prefix argument with option flags.")
		return
	}

	var folderLocation string = os.Args[0]
	split.SplitFolder(options, folderLocation)

}

