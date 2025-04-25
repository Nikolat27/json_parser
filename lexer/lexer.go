package lexer

type Token struct {
	TokenType string
	Value     string
}

const (
	LEFT_BRACE    = "LEFT_BRACE"
	RIGHT_BRACE   = "RIGHT_BRACE"
	LEFT_BRACKET  = "LEFT_BRACKET"
	RIGHT_BRACKET = "RIGHT_BRACKET"

	COLON = "COLON"
	COMMA = "COMMA"

	STRING = "STRING"
	NUMBER = "NUMBER"

	TRUE  = "TRUE"
	FALSE = "FALSE"
	NULL  = "NULL"
)

func Lexer(s string) []Token {
	counter := 0

	var tokens []Token

	for counter < len(s) {
		skipWhiteSpace(&s, &counter)
		
		char := string(s[counter])
		switch char {
		case "{":
			tokens = append(tokens, Token{TokenType: LEFT_BRACE, Value: char})
			counter++
		case "}":
			tokens = append(tokens, Token{TokenType: RIGHT_BRACE, Value: char})
			counter++
		case "[":
			tokens = append(tokens, Token{TokenType: LEFT_BRACKET, Value: char})
			counter++
		case "]":
			tokens = append(tokens, Token{TokenType: RIGHT_BRACKET, Value: char})
			counter++
		case ":":
			tokens = append(tokens, Token{TokenType: COLON, Value: char})
			counter++
		case ",":
			tokens = append(tokens, Token{TokenType: COMMA, Value: char})
			counter++
		case `"`:
			stringSequence := ""
			counter++ // skip the first `"`
			for counter < len(s) && string(s[counter]) != `"` {
				stringSequence += string(s[counter])
				counter++
			}
			counter++ // skip the second `"`
			tokens = append(tokens, Token{TokenType: STRING, Value: stringSequence})
		default:
			char := s[counter]
			if char == ' ' {
				counter++
			} else if char == '-' {
				start := counter
				counter++
				for counter < len(s) && s[counter] >= '0' && s[counter] <= '9'{
					counter++
				}
				tokens = append(tokens, Token{TokenType: NUMBER, Value: s[start:counter]})
			} else if '0' <= char && char <= '9' {
				start := counter
				for counter < len(s) && s[counter] >= '0' && s[counter] <= '9' {
					counter++
				}
				tokens = append(tokens, Token{TokenType: NUMBER, Value: s[start:counter]})
			} else if s[counter:counter+4] == "true" {
				tokens = append(tokens, Token{TokenType: TRUE, Value: "true"})
				counter += 4
			} else if s[counter:counter+5] == "false" {
				tokens = append(tokens, Token{TokenType: FALSE, Value: "false"})
				counter += 5
			} else if s[counter:counter+4] == "null" {
				tokens = append(tokens, Token{TokenType: NULL, Value: "null"})
				counter += 4
			} else {
				panic("Unknown character")
			}
		}
	}

	return tokens
}

func skipWhiteSpace(s *string, counter *int) {
	for *counter < len(*s) && ((*s)[*counter] == ' ' || (*s)[*counter] == '\n' || (*s)[*counter] == '\t') {
		*counter++
	}
}