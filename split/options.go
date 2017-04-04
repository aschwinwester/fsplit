package split

import (
	"flag"
	"log"
)

type Options struct {
	Verbose bool
	DateTimeFormat string
	Folder string
	Hours int
}

func GetOptions() Options {


	var options Options
	var verbose= flag.Bool("v", false, "Verbose logging")
	var folderName = flag.String("f", "", "Folder name")
	var dateTimeFormat = flag.String("t", "", "datetime format")
	var hours = flag.Int("h", 6, "hours difference")
	flag.Parse()

	options.Folder = *folderName
	options.Verbose = *verbose
	options.DateTimeFormat= *dateTimeFormat
	options.Hours = *hours

	if (options.Verbose) {
		log.Printf("Found folder %s", options.Folder)
		log.Printf("Datetimeformat %s", options.DateTimeFormat)
		log.Printf("Hours difference %d", options.Hours)
	}
	for _, arg := range flag.Args() {
		if (options.Verbose) {
			log.Printf("found arg %s", arg)
		}

	}

	return options
}

