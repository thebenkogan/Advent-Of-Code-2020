package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/zyedidia/generic/stack"
)

func main() {
	day := os.Args[1]
	var inputFileName string
	if len(os.Args) > 2 && os.Args[2] == "test" {
		inputFileName = "test"
	} else {
		inputFileName = "in"
	}
	pwd, _ := os.Getwd()
	path := fmt.Sprintf("%s/cmd/%s/%s.txt", pwd, day, inputFileName)
	b, _ := os.ReadFile(path)
	input := string(b)

	fmt.Printf("Part 1: %v\n", part1(input))
	fmt.Printf("Part 2: %v", part2(input))
}

type Operator string

const (
	Add Operator = "+"
	Mul Operator = "*"
)

type Expr struct {
	leftExpr  *Expr
	rightExpr *Expr
	leftVal   int
	rightVal  int
	op        Operator
}

// adds the sub expression to either the left or right child.
// returns true if this expression is full (i.e. no more children to add)
func (e *Expr) addSubExpression(sub *Expr) bool {
	if e.leftExpr == nil && e.leftVal == 0 {
		e.leftExpr = sub
		return false
	} else {
		e.rightExpr = sub
		return true
	}
}

// adds the sub value to either the left or right child.
// returns true if this expression is full (i.e. no more children to add)
func (e *Expr) addSubValue(val int) bool {
	if e.leftExpr == nil && e.leftVal == 0 {
		e.leftVal = val
		return false
	} else {
		e.rightVal = val
		return true
	}
}

func indexOfClosingParen(expr string) int {
	stack := stack.New[string]()
	for i, c := range expr {
		if string(c) == "(" {
			stack.Push("(")
		} else if string(c) == ")" {
			stack.Pop()
			if stack.Size() == 0 {
				return i
			}
		}
	}
	panic("invalid expression: " + expr)
}

// exprStr must have no whitespace
func parseExpr(exprStr string) Expr {
	expr := &Expr{}

	i := 0
	for i < len(exprStr) {
		switch string(exprStr[i]) {
		case "(":
			closeParenIdx := indexOfClosingParen(exprStr[i:])
			subExpr := parseExpr(exprStr[i+1 : i+closeParenIdx])
			isFull := expr.addSubExpression(&subExpr)
			if isFull {
				expr = &Expr{
					leftExpr: expr,
				}
			}
			i += closeParenIdx + 1
			continue
		case "+":
			expr.op = Add
		case "*":
			expr.op = Mul
		default:
			num, _ := strconv.Atoi(string(exprStr[i]))
			isFull := expr.addSubValue(num)
			if isFull {
				expr = &Expr{
					leftExpr: expr,
				}
			}
		}
		i++
	}

	return *expr.leftExpr
}

func (e *Expr) evaluate() int {
	var leftVal int
	if e.leftExpr != nil {
		leftVal = e.leftExpr.evaluate()
	} else {
		leftVal = e.leftVal
	}

	var rightVal int
	if e.rightExpr != nil {
		rightVal = e.rightExpr.evaluate()
	} else {
		rightVal = e.rightVal
	}

	switch e.op {
	case Add:
		return leftVal + rightVal
	case Mul:
		return leftVal * rightVal
	}

	panic("invalid expression operator")
}

func part1(input string) int {
	sum := 0
	for _, line := range strings.Split(input, "\n") {
		e := parseExpr(strings.ReplaceAll(line, " ", ""))
		sum += e.evaluate()
	}
	return sum
}

// exprStr must have no whitespace
func parseExprV2(exprStr string) Expr {
	expr := &Expr{}

	i := 0
	for i < len(exprStr) {
		switch string(exprStr[i]) {
		case "+":
			expr.op = Add
		case "*":
			expr.op = Mul
		case "(":
			var subExpr Expr
			closeParenIdx := indexOfClosingParen(exprStr[i:])
			if expr.op == Mul && i+closeParenIdx != len(exprStr)-1 {
				subExpr = parseExprV2(exprStr[i:])
				i = len(exprStr)
			} else {
				subExpr = parseExprV2(exprStr[i+1 : i+closeParenIdx])
				i += closeParenIdx + 1
			}
			isFull := expr.addSubExpression(&subExpr)
			if isFull {
				expr = &Expr{
					leftExpr: expr,
				}
			}
			continue
		default:
			if i != len(exprStr)-1 && expr.op == Mul {
				subExpr := parseExprV2(exprStr[i:])
				isFull := expr.addSubExpression(&subExpr)
				if isFull {
					expr = &Expr{
						leftExpr: expr,
					}
				}
				i = len(exprStr)
				continue
			} else {
				num, _ := strconv.Atoi(string(exprStr[i]))
				isFull := expr.addSubValue(num)
				if isFull {
					expr = &Expr{
						leftExpr: expr,
					}
				}
			}
		}
		i++
	}

	return *expr.leftExpr
}

func part2(input string) int {
	sum := 0
	for _, line := range strings.Split(input, "\n") {
		e := parseExprV2(strings.ReplaceAll(line, " ", ""))
		sum += e.evaluate()
	}
	return sum
}
