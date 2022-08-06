package main

import (
	"fmt"
	"ljw/cal/cmd"
	"os"
)

func main() {
	cmd := cmd.GenerateCalCmd()
	if err := cmd.Execute(); err != nil {
		fmt.Println("cal execute failed")
		os.Exit(1)
	}
	os.Exit(0)
}
