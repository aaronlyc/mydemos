/*
Copyright © 2022 Aaron <blockchain_lyc@163.com>
This file is part of CLI application foo.
*/
package main

import "mydemos/mycmd/cmd"

//go:generate go mod tidy
//go:generate make build
func main() {
	cmd.Execute()
}
