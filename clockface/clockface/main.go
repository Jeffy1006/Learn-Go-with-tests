package main

import (
	"os"
	"time"

	"example.com/myapp/clockface"
)

func main() {
	t := time.Now()
	clockface.SVGWriter(os.Stdout, t)
}
