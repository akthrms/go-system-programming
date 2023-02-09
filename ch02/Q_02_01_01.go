package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Fprintf(os.Stdout, "Write with os.Stdout at %d", 123)
	fmt.Fprintf(os.Stdout, "Write with os.Stdout at %s", "golang")
	fmt.Fprintf(os.Stdout, "Write with os.Stdout at %f", 1.23)
}
