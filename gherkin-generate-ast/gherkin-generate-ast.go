package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	gherkin "../"
)

func main() {

	var readers []io.Reader
	if len(os.Args) <= 1 {
		readers = append(readers, os.Stdin)
	} else {
		for i := range os.Args[1:] {
			file, err := os.Open(os.Args[i+1])
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error: %s\n", err)
				os.Exit(1)
				return
			}
			readers = append(readers, file)
		}
	}

	for i := range readers {

		feature, err := gherkin.ParseFeature(readers[i])
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			os.Exit(1)
			return
		}

		b, err := json.Marshal(feature)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			os.Exit(1)
			return
		}
		os.Stdout.Write(b)
		fmt.Fprint(os.Stdout, "\n")
	}
}
