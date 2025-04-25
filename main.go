package main

import (
	"fmt"
	"json_parser/lexer"
	"json_parser/parser"
)

func main() {
	jsonString := `{"arr1": [1,3,4]}`

	tokens := lexer.Lexer(jsonString)
	parsedJson := parser.Parser(tokens)
	
	value, err := parsedJson.GetArray("arr1")
	if err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Println(value)
}

