package main


import (
	"strings"
	"regexp"
)


// Condition string token type enum
//
// OPERATOR - A condition operator (i.e. '=', '<', '>', '!=', '<=','>=', AND, OR)
// COLUMN_OPERAND - A db column name (note that this condition module doesn't check for the validity of database columns)
// LITERAL_OPERAND - A string literal
// BRACKET - An opening - '(' - or closing - ')' - bracket
// WHITESPACE - A string of whitespace chars (spaces, tabs etc.) of any length
type TokenType
const (
	OPERATOR TokenType = iota
	COLUMN_OPERAND
	LITERAL_OPERAND 
	BRACKET 
	WHITESPACE 
)


// A token from the user-inputted condition string
// 
// ATTRIBUTES:
//	content - The actual underlying string of the token
//	Type - Kind of token (operator, operand, bracket, whitespace) in form of the TokenType enum 
type Token struct {
	Content string;
	Type TokenType
}


// Uses RegEx to parse a user-inputted condition string into a stream of tokens
// Returns nil if encountered an error in the condition string
func ConditionStringToTokenStream(conditionStr string) []Token {

	// Define a regex rule for each token type
	regexRules := map[TokenType]Regexp {

		OPERATOR : regexp.Compile("^((<|>)=?|!?=|AND|OR)")

		LITERAL_OPERAND  : regexp.Compile("^(?<=').*(?=')") // Literal operands must be formatted like single-quotemark strings
		COLUMN_OPERAND : regexp.Compile("^\w*") // Column names can only have alphanumeric chars and underscores (i.e. only word characters)

		BRACKET  : regexp.Compile("^(\(|\))")
		WHITESPACE : regexp.Compile("^(\s+)") // Captures strings with all whitespace chars (spaces, tabs etc) of any length

	}


	remainingMatchStr := ConditionStr
	tokenStream := make([]Token, 5)
	while(len(remainingMatchStr > 0)){

		// Determine the kind of token (operator, operand, bracket, whitespace) present at the cursor position
		// By running all regex rules against the string 
		foundMatchingToken := false
		for tokenType, rule : range regexRules {

			tokenStr := rule.Find(remainingMatchStr)

			if tokenStr != nil { // If we found a match

				tokenStream := append(tokenStream, Token{tokenStr, tokenType})
				remainingMatchStr := remainingMatchStr[len(tokenStrk):]// Trim off just examined part of remaining match string
				foundMatchingToken = true
				break;
			}	
		}	

		if (!foundMatchingToken){ // If we didn't find a matching token, must be something wrong with condition string
			return
		}
	}

	return tokenStream;
}
