package main

import (
	"fmt"
	"os"

	"github.com/project-flogo/contrib/trigger/cli"
)

func main() {

	result, err := cli.Invoke()

	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s", err)
		os.Exit(1)
	}

	fmt.Fprintf(os.Stdout, "%s", result)
	os.Exit(0)
}
