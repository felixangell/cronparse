package parse

import (
	"errors"
	"fmt"
	"strings"
)

// "*/15 0 1,15 * 1-5 /usr/bin/find"

// without looking at the spec for cron inputs,
// we have a few combinations here:
//
// 	*/N 	=> not sure what the term is for this
// 	1,15 	=> apperas to be set of inputs
// 	1-15 	=> appears to be a range
// 	* 		=> everything
// 	N 		=> exact number

// from the spec there are a few more things missing:
// ?, L, W 	=> only applicable to: day of month, day of week,
// * , - 	=> everything

// expression built up of
// minutes, hours, day of month, month, day of week, year (optional)

// non standard scheduling defs:
// @yearly, @annually, @monthly, weekly, daily, midnight, hourly, reboot (run at startup)
// these have equivalents

type ExpressionIndex int

const (
	Minute ExpressionIndex = iota
	Hour
	DayOfMonth
	Month
	DayOfWeek
	Year
	EXPRESSION_INDEX_COUNT
)

// optionalIndices a set of all the optional indices
// in the expression... we just have one.
var optionalIndices = []ExpressionIndex{Year}

/*
minute 0 15 30 45
hour 0
day of month 1 15
month 1 2 3 4 5 6 7 8 9 10 11 12
day of week 1 2 3 4 5
command /usr/bin/find
*/
type CronExpressionNode struct {
	indices []Unit
}

func (c CronExpressionNode) GetUnit(idx ExpressionIndex) Unit {
	return c.indices[idx]
}
func (c *CronExpressionNode) SetIndex(idx ExpressionIndex, u Unit) {
	c.indices[idx] = u
}

func NewExpressionNode() *CronExpressionNode {
	return &CronExpressionNode{make([]Unit, EXPRESSION_INDEX_COUNT)}
}

type UnitKind int

const (
	List     UnitKind = iota // a,b,c
	Range                    // a - b
	Wildcard                 // *
)

type Unit struct {
	// Operands on this particular unit
	// note that there is no constraint on length
	// this is a validation step
	// e.g. a, b, c => []{a, b, c}
	// a - b 		=> []{a, b}
	Operands []string
	Kind     UnitKind
}

func parseRange(value string) *Unit {
	parts := strings.Split(value, "-")
	return &Unit{
		Operands: parts,
		Kind:     Range,
	}
}

func parseUnit(value string) *Unit {
	if strings.Compare(value, "*") == 0 {
		return &Unit{
			Operands: []string{"*"},
			Kind:     Wildcard,
		}
	}

	if strings.ContainsAny(value, "-") {
		return parseRange(value)
	} else if strings.ContainsAny(value, "*") {

	}

	return nil
}

func convIndex(idx int) (ExpressionIndex, error) {
	if idx < 0 || idx >= int(EXPRESSION_INDEX_COUNT) {
		return -1, errors.New("Out of bounds")
	}
	return ExpressionIndex(idx), nil
}

func ParseCronString(input string) (*CronExpressionNode, error) {
	parts := strings.Split(input, " ")

	node := NewExpressionNode()
	for idx, value := range parts {
		exprIdx, err := convIndex(idx)
		if err != nil {
			return nil, errors.New("Failed to convert index")
		}
		unit := parseUnit(value)
		if unit == nil {
			return nil, fmt.Errorf("Failed to parse unit %s", value)
		}
		node.SetIndex(exprIdx, *unit)
	}

	return node, nil
}
