package main





// Interface for a node in the condition tree
// Implemented by ConditionTreeOperatorNode and ConditionTreeOperandNode
//
// Attributes:
//	left - left child ConditionTreeNode
//	right - right child ConditionTreeNode
type ConditionTreeNode interface {
	left *ConditionTreeNode
	right *ConditionTreeNode
}


// Condition operator enum
type ConditionOperator
const (
	EQUAL ConditionOperator = iota
	NOT_EQUAL
	GREATER
	GREATER_EQUAL // Greater than or equal to
	LESS
	LESS_EQUAL // Less than or equal to
	AND
	OR
)


// IMPLEMENTS INTERFACE ConditionTreeNode
// A condition tree node containing a condition operator as its data 
// e.g. '='. '!='. '<'. '>' 
// This will always be an internal node (i.e. it has children / isn't a leaf node)
//
// Except for AND/OR (whose kids are operator nodes themselves), all operator nodes have 
// children of type string literal node or column node 
//
// Attributes:
//	data - Underlying condition operator
//	left - left child ConditionTreeNode
//	right - right child ConditionTreeNode
type conditionTreeOperatorNode struct {
	Data ConditionOperator;
	Left *ConditionTreeNode
	Right *ConditionTreeNode
}



// IMPLEMENTS INTERFACE ConditionTreeNode
// A condition tree node containing an operand (either a string literal or a db column name) as it's data
// This will always be a leaf node (i.e. w/ no children)
//
// Attributes:
//	data - Underlying string literal
//	left - left child ConditionTreeNode
//	right - right child ConditionTreeNode
type conditionTreeOperandNode struct {
	Data string;
	Left *ConditionTreeNode
	Right *ConditionTreeNode
}





// Parse a condition string into a condition tree
// Returns a pointer to the root node of the condition tree (which will always be an operator node)
func ParseConditionString(conditionStr string) *conditionTreeOperatorNode {

	symbolStack := make([]string, 5) // Holds non-operand symbols (i.e. condition operators and parentheses)
	operandStack := make([]string, 5) // Holds encountered operands
	
	
	
}
