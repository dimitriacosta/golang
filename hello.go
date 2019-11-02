package main

import (
	"fmt"
	"github.com/dimitriacosta/stringutil"
)

func main() {
	fmt.Println("Normal: Hello world!")
	fmt.Println("Reversed:", stringutil.Reverse("Hello world!"))
}
