package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	dirPtr := flag.String("dir", ".", "Directory to organize")
	flag.Parse()

	if _, err := os.Stat(*dirPtr); os.IsNotExist(err) {
		fmt.Println("Error: Directory does not exist")
		os.Exit(1)
	}

	fmt.Println("Organizing directory:", *dirPtr)
	// Call your organize function here
}
