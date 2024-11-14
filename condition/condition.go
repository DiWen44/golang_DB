package condition

import (
	"github.com/golang_db/utils"
	"regexp"
)

// TokenType Condition string token type enum
//
// OPERATOR - A condition operator (i.e. '=', '<', '>', '!=', '<=','>=', &, |)
// COLUMN_OPERAND - A db column name (note that this condition module doesn't check for the validity of database columns)
// LITERAL_OPERAND - A string literal
// BRACKET - An opening - '(' - or closing - ')' - bracket
// WHITESPACE - A string of whitespace chars (spaces, tabs etc.) of any length
type TokenType int

const (
	OPERATOR TokenType = iota
	COLUMN_OPERAND
	LITERAL_OPERAND
	OPENING_BRACKET
	CLOSING_BRACKET
	WHITESPACE
)

// Token A token from the user-inputted condition string
//
// ATTRIBUTES:
//
//	content - The actual underlying string of the token
//	kind - Type of token (operator, operand, bracket, whitespace) in form of the TokenType enum
type Token struct {
	content string
	kind    TokenType
}

// Uses RegEx to parse a user-inputted condition string into a stream of tokens
// Returns nil if encountered an error in the condition string
func conditionStringToTokenStream(conditionStr string) []Token {

	// Define a regex rule (in form of a string) for each token type
	// Note that all regexes here are anchored to beginning of string - so, when regexp.FindString()
	// is called, they will always find a match that starts at the cursor position
	regexRules := map[TokenType]string{

		OPERATOR: `^((<|>)=?|!?=|&|\|)`,

		OPENING_BRACKET: `^\(`,
		CLOSING_BRACKET: `^\)`,
		WHITESPACE:      `^\s+`, // Captures strings with all whitespace chars (spaces, tabs etc.) of any length

		LITERAL_OPERAND: `^'.*?'`, // Literal operands must be formatted like single-quotemark strings
		COLUMN_OPERAND:  `^\w+`,   // Column names can only have alphanumeric chars and underscores (i.e. only word characters)
	}

	remainingMatchStr := conditionStr
	tokenStream := make([]Token, 5)
	for len(remainingMatchStr) > 0 {

		// Determine the kind of token (operator, operand, bracket, whitespace) present at the cursor position
		// By running all regex rules against the string
		foundMatchingTokenType := false
		for tokenType, ruleStr := range regexRules {

			rule, _ := regexp.Compile(ruleStr)                  // Compile rule string to regexp object
			matchIdx := rule.FindStringIndex(remainingMatchStr) // Get position of match (if there is one)

			if matchIdx != nil { // If we found a match for the rule

				matchedStr := remainingMatchStr[matchIdx[0]:matchIdx[1]]

				if tokenType != WHITESPACE { // Don't add whitespace tokens to stream - these are superfluous
					tokenStream = append(tokenStream, Token{matchedStr, tokenType})
				}

				remainingMatchStr = remainingMatchStr[matchIdx[1]:] // Trim off just examined part of remaining match string
				foundMatchingTokenType = true
				break

			}
		}

		// If we didn't find a matching token type, must be something wrong with condition string
		if !foundMatchingTokenType {
			return nil
		}
	}

	return tokenStream
}

// ResolveCondition resolves a condition specified by a condition string on a database entry
//
// PARAMS:
//
//	conditionStr - The user-inputted condition string
//	entry - A hashmap with db column names as keys and the entry's row values for those columns as values
//
// RETURNS:
// - true if condition is true
// - false if not
func ResolveCondition(conditionStr string, entry map[string]string) bool {

	tokens := conditionStringToTokenStream(conditionStr) // Convert string to token stream first

	// Symbol and operand stack hold strings representing the actual content of tokens
	// operand stack contains operand strings
	// symbol stack contains operator strings and bracket strings
	symbolStack := utils.MakeStack[string]()
	operandStack := utils.MakeStack[string]()
	boolStack := utils.MakeStack[bool]()

	for _, token := range tokens {
		switch token.kind {

		case OPERATOR: // Push operators and opening brackets onto symbol stack
			symbolStack.Push(token.content)
		case OPENING_BRACKET:
			symbolStack.Push(token.content)

		case COLUMN_OPERAND: // Put entry's value at that column on operand stack
			operandStack.Push(entry[token.content])

		case LITERAL_OPERAND: // Put literals on operand stack
			operandStack.Push(token.content)

		case CLOSING_BRACKET: // If closing bracket, apply operation at top of stack

			operator := symbolStack.Pop() // Pop last-read operator
			_ = symbolStack.Pop()         // Symbol before an actual operator token (that was just popped) should always be an opening bracket token, so pop that opening bracket as well
			var res bool

			// If operator is AND/OR, pop top 2 elems from bool stack, apply operation, then push the result back to bool stack
			// If operator isn't AND/OR, do the same but pop the top 2 elems from operand stack instead of bool stack
			if operator == "&" {
				o1 := boolStack.Pop()
				o2 := boolStack.Pop()
				res = o1 && o2
			} else if operator == "|" {
				o1 := boolStack.Pop()
				o2 := boolStack.Pop()
				res = o1 || o2
			} else {

				o1 := operandStack.Pop()
				o2 := operandStack.Pop()

				switch operator {
				case "=":
					res = o1 == o2
				case "!=":
					res = o1 != o2
				case "<":
					res = o1 < o2
				case "<=":
					res = o1 <= o2
				case ">":
					res = o1 > o2
				case ">=":
					res = o1 >= o2
				}
			}
			boolStack.Push(res)

		default:
			continue
		}
	}

	return boolStack.Pop() // Final result is only element left on boolstack
}
