package main

import (
	"os"
	"time"

	"learn-go-with-tests/clockface/svg"
)

func main() {
	t := time.Now()
	svg.SVGWriter(os.Stdout, t)
}
