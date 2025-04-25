package parser

import (
	"fmt"
	"json_parser/lexer"
	"strings"
)

type ObjectNode struct {
	nodeType string
	value    lexer.Token
	children map[string]*ObjectNode
	array    []*ObjectNode
}

// It receives the token type and the value, and returns a reference to the node.
func buildNode(tokenType string, value string) *ObjectNode {
	return &ObjectNode{
		nodeType: strings.ToLower(tokenType),
		value: lexer.Token{
			TokenType: tokenType,
			Value:     value,
		},
		children: make(map[string]*ObjectNode),
		array:    make([]*ObjectNode, 0),
	}
}

func parseValue(tokens []lexer.Token, currentPosition *int) *ObjectNode {
	currentTokenType := tokens[*currentPosition].TokenType
	switch currentTokenType {
	case lexer.STRING, lexer.NUMBER, lexer.TRUE, lexer.FALSE, lexer.NULL:
		node := buildNode(currentTokenType, tokens[*currentPosition].Value)
		*currentPosition++
		return node
	case lexer.LEFT_BRACE:
		return parseObject(tokens, currentPosition)
	case lexer.LEFT_BRACKET:
		return parseArray(tokens, currentPosition)
	default:
		panic("Unexpected token: " + currentTokenType)
	}
}

func parseObject(tokens []lexer.Token, currentPosition *int) *ObjectNode {
	// Skip the opening left brace '{'
	*currentPosition++

	objectNode := buildNode("object", "")

	// Loop until the closing right brace '}' is encountered
	for tokens[*currentPosition].TokenType != lexer.RIGHT_BRACE {
		// Check for end of tokens
		if *currentPosition >= len(tokens) {
			panic("Unexpected end of tokens, expected right brace '}'")
		}

		// Skip commas ','
		if tokens[*currentPosition].TokenType == lexer.COMMA {
			*currentPosition++
			continue
		}

		if tokens[*currentPosition].TokenType == lexer.STRING {
			// Extract the key
			key := tokens[*currentPosition].Value
			*currentPosition++

			if tokens[*currentPosition].TokenType != lexer.COLON {
				panic("Invalid JSON, expected a colon after key")
			}
			*currentPosition++

			value := parseValue(tokens, currentPosition)
			objectNode.children[key] = value
		} else {
			panic("Invalid JSON, expected a string key")
		}
	}

	// Skip the closing right brace '}'
	*currentPosition++
	return objectNode
}

func parseArray(tokens []lexer.Token, currentPosition *int) *ObjectNode {
	// Skip the opening left bracket '['
	*currentPosition++

	arrayNode := buildNode("array", "")

	// Loop until the closing right bracket ']' is encountered
	for tokens[*currentPosition].TokenType != lexer.RIGHT_BRACKET {
		// Check for end of tokens
		if *currentPosition >= len(tokens) {
			panic("Unexpected end of tokens, expected right bracket ']'")
		}

		// Skip commas ','
		if tokens[*currentPosition].TokenType == lexer.COMMA {
			*currentPosition++
			continue
		}

		// Parse the value (this handles nested arrays and objects)
		value := parseValue(tokens, currentPosition)
		arrayNode.array = append(arrayNode.array, value)
	}

	// Skip the closing right bracket ']'
	*currentPosition++
	return arrayNode
}

func Parser(tokens []lexer.Token) *ObjectNode {
	if len(tokens) == 0 {
		panic("No tokens to parse")
	}

	i := 0
	return parseValue(tokens, &i)
}

func GetAllArrayElements(arr *ObjectNode) []string {
	var parsedArray []string
	for _, element := range arr.array {
		parsedArray = append(parsedArray, element.value.Value)
	}
	return parsedArray
}

func (n *ObjectNode) GetElement(key string) (string, error) {
	if key == "" {
		return "", fmt.Errorf("key is empty")
	}
	element, exists := n.children[key]
	if !exists {
		return "", fmt.Errorf("key '%s' not found", key)
	}
	return element.value.Value, nil
}

// GetArray Access an array by key
func (n *ObjectNode) GetArray(key string) ([]string, error) {
	if key == "" {
		return nil, fmt.Errorf("key is empty")
	}

	array := n.children[key]
	newArray := GetAllArrayElements(array)
	return newArray, nil
}

// GetArrayElement Access a specific element in an array by key and index
func (n *ObjectNode) GetArrayElement(key string, index int) (string, error) {
	if key == "" {
		return "", fmt.Errorf("len is 0")
	}

	array, exists := n.children[key]

	if !exists {
		return "", fmt.Errorf("key '%s' not found", key)
	}

	if index > len(array.array) {
		return "", fmt.Errorf("index out of range")
	}

	element := array.array[index].value.Value
	return element, nil
}

// GetObject Get the whole object
func (n *ObjectNode) GetObject(key string) (map[string]string, error) {
	newObjects := map[string]string{}
	objects := n.children[key]
	for key, value := range objects.children {
		newObjects[key] = value.value.Value
	}
	return newObjects, nil
}

// GetObjectValue Get the object`s value by its key
func (n *ObjectNode) GetObjectValue(key1, key2 string) (string, error) {
	if key1 == "" || key2 == "" {
		return "", fmt.Errorf("one key is empty")
	}

	object, exists := n.children[key1]
	if !exists {
		return "", fmt.Errorf("object '%s' not found", key1)
	}
	value, exists := object.children[key2]
	if !exists {
		return "", fmt.Errorf("value '%s' not found", key2)
	}
	return value.value.Value, nil
}
