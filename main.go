package main

import (
	"flag"
	"fmt"

	"github.com/jjg/wsync/getfiles"
)

func main() {

	// Get command line arguments.
	var source = flag.String("s", ".", "Filespec of files to upload.")
	var destination = flag.String("d", "http://example.com", "URL to upload files to.")
	var recursive = flag.Bool("r", false, "Include subdirectories (recursive).")
	flag.Parse()

	// Get a list of files to send.
	var files []string

	// Use fs.WalkDir when doing recursive upload, otherwise use os.ReadDir.
	if *recursive {
		files = getfiles.Recursive(source)
	} else {
		files = getfiles.Flat(source)
	}

	// TODO: Actually upload the files.
	fmt.Printf("Total files selected for upload: %v\n", len(files))
	fmt.Println("Uploading:")
	for f := range files {
		fmt.Printf("%v%v\n", *destination, files[f])
	}
}
