package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	name := os.Args[1]

	nf, err := os.Create("file.txt")

	if err != nil {
		log.Fatal("error creating file", err)
	}
	defer nf.Close()
	defer fmt.Println(1)
	defer fmt.Println(2)

	io.Copy(nf, strings.NewReader(name))
}
