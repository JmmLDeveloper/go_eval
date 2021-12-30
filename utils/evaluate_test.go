package utils

import (
	"fmt"
	"math"
	"testing"
)


func TestCorrectExpressions(t *testing.T) {
	tests := []struct {
        input string
        want float64
    }{
		{"2+2",4.0},
		{"	2 +2 -\n6     ",-2.0},
		{" 2 - 5 * 7 + 9 / 8 ",-31.875},
		{" (2 + 6) * ( (6+8) / 8 ) ",14},
		{"( (3 + 5) * ( (6-8) / 7 ) )",-2.28571428571},
        {"( 12 / 8 -2 )", -0.5},
        {"-3", -3},
        {"+6", 6},
        {"2 * 3 * (-2/3)", -4},


    }

    for _, tt := range tests {
        testname := fmt.Sprintf(` "%s" should evaluate to %.2f `, tt.input, tt.want)
        t.Run(testname, func(t *testing.T) {
            ans,err := Evaluate(tt.input)
            if  err != nil {
                t.Errorf("got invalid expression, want %f", tt.want)
            } else if  math.Abs(ans - tt.want) > 0.01 {
                t.Errorf("got %f, want %f", ans, tt.want)
            }
        })
	}
}

func TestInvalidExpressions(t *testing.T) {
	tests := []string {
		"+",
		"++++",
		"+q 23++*",
		"(((())))",
        " 2 + a3 ",
 		"(2+3 -(( 3 * 4 ) )",
        "(2+3 -( 3 * 4 )) )",
        "",
    }

    for _, tt := range tests {
        testname := fmt.Sprintf(` "%s" should evaluate to invalid expression`, tt)
        t.Run(testname, func(t *testing.T) {
            _,err := Evaluate(tt)
            if  err == nil {
                t.Errorf("this expression should return an error")
            } 
        })
	}
}