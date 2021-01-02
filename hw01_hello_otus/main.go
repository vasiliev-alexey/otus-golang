package main

import (
	"github.com/golang/example/stringutil"
	"os"
)

func main() {
	os.Stdout.WriteString(stringutil.Reverse("Hello, OTUS!"))
}
