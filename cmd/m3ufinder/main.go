package main

import (
	"fmt"
	"m3ufinder-cli/cmd"
	"os"
)

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}