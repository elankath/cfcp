package main

import (
	"fmt"
	"os"

	"github.com/jawher/mow.cli"
)

func main() {
	app := cli.App("cfcp", "Copy files from local to CF container")
	app.Spec = "[-a] SRC... DST"
	var (
		// declare the -r flag as a boolean flag
		appName = app.BoolOpt("r recursive", false, "Copy files recursively")
		// declare the SRC argument as a multi-string argument
		src = app.StringsArg("SRC", nil, "Source files to copy")
		// declare the DST argument as a single string (string slice) arguments
		dst = app.StringArg("DST", "", "Destination where to copy files to")
	)

	// Specify the action to execute when the app is invoked correctly
	app.Action = func() {
		fmt.Printf("Copying %v to %s [recursively: %v]\n", *src, *dst, *recursive)
	}

	// Invoke the app passing in os.Args
	app.Run(os.Args)
}
