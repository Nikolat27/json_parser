package main

import (
	"fmt"
	"json_parser/lexer"
	"json_parser/parser"
)

func main() {
	jsonString := `{"dic1": [1,[5, 6, 7],3,4]}`

	tokens := lexer.Lexer(jsonString)
	fmt.Println(tokens)
	parsedJson := parser.Parser(tokens)
	
	value, err := parsedJson.GetArray("dic1")
	if err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Println(value)
}

