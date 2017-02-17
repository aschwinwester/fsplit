package split

import (
	"flag"
	"log"
)

type Options struct {
	Verbose bool
	DateTimeFormat string
	Folder string
}

func GetOptions() Options {


	var options Options
	var verbose= flag.Bool("v", false, "Verbose logging")
	var folderName = flag.String("f", "", "Folder name")
	var dateTimeFormat = flag.String("t", "", "datetime format")
	flag.Parse()

	options.Folder = *folderName
	options.Verbose = *verbose
	options.DateTimeFormat= *dateTimeFormat

	if (options.Verbose) {
		log.Printf("Found folder %s", options.Folder)
		log.Printf("Datetimeformat %s", options.DateTimeFormat)
	}
	for _, arg := range flag.Args() {
		if (options.Verbose) {
			log.Printf("found arg %s", arg)
		}

	}




	return options
}

