package utils

import (
	"errors"
	"regexp"
	"strconv"
)

func tokenize(s string,) (tokens []string) {

	r := regexp.MustCompile(`\s+`)
	s =  r.ReplaceAllString(s,"")

	separators := []string{"+","-","*","/","(",")"}
	tokens = make([]string, 0)
	start_idx := 0
	for crr_idx := 0 ; crr_idx < len(s) ; crr_idx++ {
		for _,sep := range separators {
			if s[crr_idx:crr_idx+1] == sep  {
				if crr_idx != start_idx {
					tokens = append(tokens, s[start_idx:crr_idx])
				}
				tokens = append(tokens , sep )
				start_idx = crr_idx + 1
			}

		} 
	}

	if start_idx != len(s)  {
		tokens = append(tokens, s[start_idx:])
	}

	return
}

func isOneOf( s string, candidates []string ) bool {
	for _,candidate := range candidates {
		if s == candidate {
			return true
		}
	}

	return false
}

func replaceParserList( elements []ParserElement,start,end int ,element ParserElement) []ParserElement {
	if start >= end {
		return elements
	}

	result := elements[0:len(elements) - (end - 1 - start)]
	copy(result[start+1:],elements[end:])
	result[start] = element
	return result
}

func indexOfClosingParenthesis(elements []ParserElement,start int) int  {
	for i := start + 1 ; i < len(elements)  ; i++ {
		el := elements[i]
		if el.Token != nil {
			token := *el.Token
			if token == ")" {
				return i
			} else if token == "(" {
				if r := indexOfClosingParenthesis(elements,i) ; r == -1 {
					return r
				} else {
					i =  r
				}
			}
		}
	}
	return -1
}

func internalParse(elements []ParserElement) ([]ParserElement,error)  {
	// parse sub expressions in parenthesis
	for i := 0 ; i < len(elements) ; i++{
		element := elements[i]
		if element.Token != nil {
			token := *element.Token
			if isOneOf(token, []string{"("} ) {
				closingIndex := indexOfClosingParenthesis(elements,i)
				if closingIndex == -1 {
					return nil,errors.New("invalid expression")
				}
				subElements,err := internalParse(elements[i+1:closingIndex])
				if err != nil {
					return nil,err
				}
				elements = replaceParserList(elements,i,closingIndex + 1 ,ParserElement{Exp: subElements[0].Exp})
				i--
			}
		}
	}

	if len(elements) > 0 && elements[0].Token != nil && isOneOf(*elements[0].Token, []string {"+","-"}) {
		if len(elements) == 1 || elements[1].Exp == nil  {
			// implies there is a unary operator without operand
			return nil,errors.New("invalid expression")
		} else{
			var ex SyntaxTree = UnaryOperation{ operator: UnaryOperator(*elements[0].Token) , Child: *elements[1].Exp  } 
			elements = replaceParserList(elements,0,2,ParserElement{Exp: &ex})
		}
	}

	// resolve binary operators , first it solves "*","/" and then  "+","-"
	// if the index is 0 it tries to 

	for _,operatorsGroup := range [][]string{ {"*","/"}, {"+","-"} } {
		for i := 0 ; i < len(elements) ; i++ {
			element := elements[i]
			if element.Token != nil {
				token := *element.Token
				if isOneOf(token, operatorsGroup ) {
					if i == 0 || i == len(elements) -1 || elements[i-1].Exp == nil || elements[i+1].Exp == nil {
						// implies there is a binary operation without operands
						return nil,errors.New("invalid expression")
					} else{
						ex := NewBinOp(*elements[i-1].Exp,*elements[i+1].Exp, BinaryOperator(token) )
						elements = replaceParserList(elements,i-1,i+2 ,ParserElement{Exp: &ex})
						i--
					}
				}
			}
		}
	}

	if len(elements) != 1 || elements[0].Exp == nil {
		return nil,errors.New("invalid expression")
	} 
	return elements,nil
}

func parse( tokens []string ) (syntaxTree SyntaxTree,returnErr error) {

	elements := make([]ParserElement,len(tokens))

	//parses all numbers

	for idx,token := range tokens {
		if num,err := strconv.ParseFloat(token,64); err == nil {
			val := NewVal(num)
			elements[idx] = ParserElement{Exp: &val  }
		} else {
			elements[idx] = ParserElement{Token: &tokens[idx]}
		}
	} 

	newElements,err := internalParse( elements )

	if  err != nil {
		return nil,err
	} 
	return *newElements[0].Exp,nil
}

func Evaluate(input string) (float64,error) {
	tokens := tokenize(input)
	syntaxTree,err := parse(tokens)

	if err != nil {
		return 0,err
	} 
	return syntaxTree.Evaluate(),nil
}


