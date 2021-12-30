package utils

import "fmt"

type SyntaxTree interface {
	Evaluate() float64
	String() string
}

// Binary Operation Expression

type BinaryOperator string 

const (
	Sum BinaryOperator = "+"
	Div BinaryOperator = "/"
	Mul BinaryOperator = "*"
	Sub BinaryOperator = "-"
)

type BinaryOperation struct {
	operator BinaryOperator
	Left SyntaxTree
	Right SyntaxTree
}

func ( t BinaryOperation ) Evaluate() float64 {
	var result float64
	
	switch t.operator {
	case Sum:
		result = t.Left.Evaluate() + t.Right.Evaluate()
	case Sub:
		result = t.Left.Evaluate() - t.Right.Evaluate()
	case Div:
		result = t.Left.Evaluate() / t.Right.Evaluate()
	case Mul:
		result = t.Left.Evaluate() * t.Right.Evaluate()
	}
	return result
}

func ( t BinaryOperation ) String() string {
	var name string

	switch t.operator {
	case "*":
		name = "MUL"
	case "+":
		name = "SUM"
	case "-":
		name = "SUB"
	case "/":
		name = "DIV"
		
	}

	return fmt.Sprintf("%s(%s,%s)",name,t.Left.String(),t.Right.String())
}



func NewBinOp(a,b SyntaxTree, op BinaryOperator) SyntaxTree {
	return BinaryOperation{
		operator: op,
		Left: a,
		Right: b,
	}
}

// Unary Operation Expression

type UnaryOperator string

const (
	Minus UnaryOperator = "-"
	Plus UnaryOperator = "+"
)

type UnaryOperation struct {
	operator UnaryOperator
	Child SyntaxTree
}

func ( t UnaryOperation ) Evaluate() float64 {
	var result float64
	switch t.operator {
	case Minus:
		result = -t.Child.Evaluate()
	case Plus:
		result = +t.Child.Evaluate()
	}
	return result
}

func ( t UnaryOperation ) String() string {
	var name string
	switch t.operator {
	case Minus:
		name = "Neg"
	case Plus:
		name = "Plus"
	}
	return fmt.Sprintf("%s(%s)",name,t.Child.String())
}




// Value Expression

type Value struct {
	value float64
}


func ( t Value ) Evaluate() float64 {
	return t.value
}



func PrintTokens(tokens []string) {
	for _, token := range tokens {
		fmt.Printf(`"%s" `,token)
	}
	fmt.Println()
}


func ( t Value ) String() string {
	return fmt.Sprintf(" N(%f) ",t.value)
}


func NewVal(x float64) SyntaxTree { return Value{value: x}; }

// a parser element is either a token string or a expression is useful when parsing 

type ParserElement struct {
	Token *string
	Exp   *SyntaxTree
}

func (e ParserElement) String() string {
	if e.Token != nil {
		return *e.Token
	}	else if e.Exp !=nil  {
		return (*e.Exp).String()
	}
	return "Error"
}

func NewParserElement(token *string, exp *SyntaxTree) ParserElement {
	if (token == nil && exp == nil) || (token != nil && exp != nil) {
		panic("NewParserElement function expects one nil pointer and one normal pointer")
	}
	return ParserElement{Exp: exp, Token: token}
}
