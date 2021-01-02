package main

import (
	"os"

	"github.com/golang/example/stringutil"
)

func main() {
	os.Stdout.WriteString(stringutil.Reverse("Hello, OTUS!"))
}
