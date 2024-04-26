package main

import (
	"fmt"
	"mydemos/mycmd/cmd/transform/app"
	"os"
)

func main() {
	command := app.NewTransformCommand()
	err := command.Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
	}
}
