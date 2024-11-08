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
)


// IMPLEMENTS INTERFACE ConditionTreeNode
// A condition tree node containing a condition operator as its data 
// e.g. '='. '!='. '<'. '>' 
//
// Attributes:
//	data - Underlying condition operator
//	left - left child ConditionTreeNode
//	right - right child ConditionTreeNode
type conditionTreeOperatorNode struct {
	data ConditionOperator;
	left *ConditionTreeNode
	right *ConditionTreeNode
}



// IMPLEMENTS INTERFACE ConditionTreeNode
// A condition tree node containing a string literal operand as its data 
//
// Attributes:
//	data - Underlying string literal
//	left - left child ConditionTreeNode
//	right - right child ConditionTreeNode
type conditionTreeStringLiteralNode struct {
	data string;
	left *ConditionTreeNode
	right *ConditionTreeNode
}


// IMPLEMENTS INTERFACE ConditionTreeNode
// A condition tree node containing a db column name as an operand 
//
// Attributes:
//	col - Underlying db column name
//	left - left child ConditionTreeNode
//	right - right child ConditionTreeNode
type conditionTreeColumnNode struct {
	col string;
	left *ConditionTreeNode
	right *ConditionTreeNode
}


