package main

import (
    "fmt"
    "os"

    "./parser"
)

func main() {
    arg := os.Args[1]
    scanner := new(parser.Scanner)
    scanner.Init(arg)
    ast := parser.Parse(scanner)
    fmt.Printf("%#v", ast)
}