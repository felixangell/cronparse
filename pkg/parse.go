package parse

import (
	"errors"
	"fmt"
	"log"
	"strconv"
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
	ExpressionIndexCount
)

func (e ExpressionIndex) RangeForType() (int, int) {
	switch e {
	case Minute:
		return 0, 59
	case Hour:
		return 0, 23
	case DayOfMonth:
		return 1, 31
	case Month:
		return 1, 12
	case DayOfWeek:
		// non standardised range (7 is sunday?)
		return 0, 7
	case ExpressionIndexCount:
		fallthrough
	default:
		panic("Invalid state")
	}
}

/*
minute 0 15 30 45
hour 0
day of month 1 15
month 1 2 3 4 5 6 7 8 9 10 11 12
day of week 1 2 3 4 5
command /usr/bin/find
*/
type CronExpressionNode struct {
	indices []*Unit
}

type ParseResult struct {
	ExpressionNode *CronExpressionNode
	Command        string
}

// GetUnit will return the unit at the given index and
// whether or not it exist or not.
func (c CronExpressionNode) GetUnit(idx ExpressionIndex) (*Unit, bool) {
	if idx < 0 || idx > ExpressionIndexCount {
		return nil, false
	}
	val := c.indices[idx]
	return val, val != nil
}
func (c *CronExpressionNode) SetIndex(idx ExpressionIndex, u *Unit) {
	c.indices[idx] = u
}

func NewExpressionNode() *CronExpressionNode {
	return &CronExpressionNode{make([]*Unit, ExpressionIndexCount)}
}

type UnitKind string

const (
	List     UnitKind = "list"     // a,b,c
	Range    UnitKind = "range"    // a - b
	Interval UnitKind = "interval" // */N
	Wildcard UnitKind = "wildcard" // *
	Integer  UnitKind = "integer"  // N
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

func (u Unit) String() string {
	return fmt.Sprintf("%v: %v", u.Kind, u.Operands)
}

func parseRange(value string) *Unit {
	parts := strings.Split(value, "-")
	return &Unit{
		Operands: parts,
		Kind:     Range,
	}
}

func parseList(value string) *Unit {
	parts := strings.Split(value, ",")
	return &Unit{
		Operands: parts,
		Kind:     List,
	}
}

func parseInterval(value string) *Unit {
	parts := strings.Split(value, "/")

	// purely to validate the integer/interval
	if _, err := strconv.ParseInt(parts[1], 10, 64); err != nil {
		log.Println(err)
		return nil
	}

	return &Unit{
		Operands: []string{parts[1]},
		Kind:     Interval,
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
	} else if strings.ContainsAny(value, ",") {
		return parseList(value)
	} else if strings.ContainsAny(value, "*") {
		return parseInterval(value)
	}

	if _, err := strconv.ParseInt(value, 10, 64); err == nil {
		return &Unit{Operands: []string{value}, Kind: Integer}
	}

	return nil
}

func convIndex(idx int) (ExpressionIndex, error) {
	if idx < 0 || idx >= int(ExpressionIndexCount) {
		return -1, errors.New("Out of bounds")
	}
	return ExpressionIndex(idx), nil
}

func ParseCronString(input string) (*ParseResult, error) {
	parts := strings.Split(input, " ")

	// take the last value in the array
	// as the command. i dont think this is to spec though.
	maxLen := minOf(int(ExpressionIndexCount), len(parts))
	command := parts[maxLen:]
	if len(command) == 0 {
		return nil, errors.New("No command specified")
	}

	node := NewExpressionNode()

	for idx, value := range parts[0:ExpressionIndexCount] {
		exprIdx, err := convIndex(idx)
		if err != nil {
			return nil, errors.New("Failed to convert index")
		}
		unit := parseUnit(value)
		if unit == nil {
			return nil, fmt.Errorf("Failed to parse unit %s", value)
		}
		node.SetIndex(exprIdx, unit)
	}

	return &ParseResult{
		ExpressionNode: node,
		Command:        command[0],
	}, nil
}

func minOf(values ...int) int {
	min := values[0]
	for _, i := range values {
		if min > i {
			min = i
		}
	}
	return min
}
