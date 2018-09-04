package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	for {
		fmt.Fprintf(os.Stdout, "printing to stdout\n")
		time.Sleep(100 * time.Millisecond)
	}
}
