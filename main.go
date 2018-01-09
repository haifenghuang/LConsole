package main

import (
	"fmt"
	"lconsole/console"
	"os"
)

func main() {
	fmt.Println("Liner based console program for syntax highlight\n")
	console.Start(os.Stdout, true)
}
