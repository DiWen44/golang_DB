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
//	kind - Type of token (operator, operand, bracket, whitespace) in form of the TokenType enum 
type Token struct {
	content string;
	kind TokenType
}


// Uses RegEx to parse a user-inputted condition string into a stream of tokens
// Returns nil if encountered an error in the condition string
func conditionStringToTokenStream(conditionStr string) []Token {

	// Define a regex rule for each token type
	regexRules := map[TokenType]Regexp {

		OPERATOR : regexp.Compile("^((<|>)=?|!?=|AND|OR)"),

		LITERAL_OPERAND  : regexp.Compile("^(?<=').*(?=')"), // Literal operands must be formatted like single-quotemark strings
		COLUMN_OPERAND : regexp.Compile("^\w*"), // Column names can only have alphanumeric chars and underscores (i.e. only word characters)

		BRACKET  : regexp.Compile("^(\(|\))"),
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


// resolves a condition specified by a condition string on a database entry
//
// PARAMS:
//	conditionStr - The user-inputted condition string
//	entry - A hashmap with db column names as keys and the entry's row values for those columns as values 
func resolveCondition(conditionStr string, entry map[string]string){

	tokens := conditionStringToTokenStream(conditionStr) // Convert string to token stream first

	symbolStack := makeStack()
	operandStack := makeStack()
	boolStack := makeStack()

	// First, resolve basic conditions
	// A basic condition is any condition that does not contain AND/OR
	for _, token : range tokens {


		switch token.kind {

			case OPERATORS: // Push operators onto symbol stack
				symbolStack.push(token.content)

			case COLUMN_OPERAND: // Put entry's value at that column on operand stack
				operandStack.push(entry[token.content])

			case LITERAL_OPERAND: // Put literals on operand stack
				operandStack.push(token.content)

			case BRACKET:

				if token.content == ")" { // If closing bracket, apply operation at top of stack

					operator := symbolStack.pop()
					var res bool

					// If operator is AND/OR, pop top 2 elems from bool stack, apply operation, then push the result back to bool stack
					// If operator isn't AND/OR, do the same but pop the top 2 elems from operand stack instead of bool stack
					if operator == "AND"{
						o1 := boolStack.pop()
						o2 := boolStack.pop()
						res = o1 && o2
					} 
					else if operator == "OR" {
						o1 := boolStack.pop()
						o2 := boolStack.pop()
						res = o1 || o2
					}
					else {

						o1 := operandStack.pop()
						o2 := operandStack.pop()

						
						switch operator {
							case "=":
								res = (o1 == o2)
							case "!=":
								res = (o1 != o2)
							case "<":
								res = (o1 < o2)
							case "<=":
								res = (o1 <= o2)
							case ">":
								res = (o1 > o2)
							case ">=":
								res = (o1 >= o2)
						}
					}
					boolStack.push(res)


				}
				else { // Push opening brackets to symbol stack
					symbolStack.push(token.content)

				}

			default:
				continue
		}
	}

	return boolStack.pop() // Final result is only element left on boolstack
}
