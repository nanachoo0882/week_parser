package main

import (
	"encoding/json"
	"fmt"
	"os"

	"./parser"
)

func main() {
	arg := os.Args[1]
	scanner := new(parser.Scanner)
	scanner.Init(arg)
	ast := parser.Parse(scanner)

	resultMap := parser.Eval(ast)
	jsonObj := toJSON(resultMap)
	result, _ := json.Marshal(jsonObj)
	fmt.Print(string(result))
}

func toJSON(originalMap map[rune]bool) map[string]bool {
	m := map[string]bool{}
	for k, v := range originalMap {
		m[string(k)] = v
	}
	return m
}