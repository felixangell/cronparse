package internal

import (
	"fmt"
	parse "github.com/felixangell/cronparse/pkg"
	"log"
	"sort"
	"strconv"
	"strings"
)

type NodePrettyPrinter struct {
}

func PrettyPrintCron(result parse.ParseResult, out func(a string)) {
	dateStrings := []string{
		"minute", "hour", "day of month", "month", "day of week",
	}

	prettyPrinter := NodePrettyPrinter{}

	for idx := 0; idx < int(parse.ExpressionIndexCount); idx++ {
		exprType := parse.ExpressionIndex(idx)
		unit, ok := result.ExpressionNode.GetUnit(exprType)
		if !ok {
			log.Fatal("Bad unit at index", idx)
		}
		out(fmt.Sprintf("%v %v", dateStrings[idx], prettyPrinter.Print(*unit, exprType)))
	}
	out(fmt.Sprintf("command %v", result.Command))

}

func (p NodePrettyPrinter) Print(unit parse.Unit, typ parse.ExpressionIndex) string {
	switch unit.Kind {
	case parse.List:
		return p.printList(unit.Operands)
	case parse.Range:
		return p.printRange(unit.Operands)
	case parse.Interval:
		return p.printInterval(unit.Operands, typ)
	case parse.Wildcard:
		return p.printWildcard(unit.Operands, typ)
	case parse.Integer:
		return p.printInteger(unit.Operands)
	}

	log.Fatal("Unhandled unit type", unit.Kind)
	return ""
}

// ./cronparse '*/15 0 1,15 * 1-5 /usr/bin/find'
//minute 0 15 30 45
//hour 0
//day of month 1 15
//month 1 2 3 4 5 6 7 8 9 10 11 12
//day of week 1 2 3 4 5
//command /usr/bin/find

type ByIntValueAscending []string

func (a ByIntValueAscending) Len() int { return len(a) }
func (a ByIntValueAscending) Less(i, j int) bool {
	fst, _ := strconv.Atoi(a[i])
	snd, _ := strconv.Atoi(a[j])
	return fst < snd
}
func (a ByIntValueAscending) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

func (p NodePrettyPrinter) printList(vals []string) string {
	sort.Sort(ByIntValueAscending(vals))
	return strings.Join(vals, " ")
}

func (p NodePrettyPrinter) printRange(vals []string) string {
	// NOTE: we dont error check here as the printer
	// assumes that the input is valid and that these
	// checks are done at parse time.
	start, _ := strconv.ParseInt(vals[0], 10, 64)
	end, _ := strconv.ParseInt(vals[1], 10, 64)

	results := make([]string, end)
	idx := 0
	for i := start; i <= end; i++ {
		results[idx] = fmt.Sprintf("%v", i)
		idx++
	}

	return strings.Trim(strings.Join(results, " "), " ")
}

func (p NodePrettyPrinter) printInterval(vals []string, typ parse.ExpressionIndex) string {
	// todo base on length
	interval, _ := strconv.ParseInt(vals[0], 10, 64)
	start, end := typ.RangeForType()

	result := []string{}
	for i := int64(start); i < int64(end); i += interval {
		result = append(result, fmt.Sprintf("%v", i))
	}
	return strings.Join(result, " ")
}

func (p NodePrettyPrinter) printWildcard(vals []string, typ parse.ExpressionIndex) string {
	start, end := typ.RangeForType()
	result := []string{}
	for i := int64(start); i <= int64(end); i++ {
		result = append(result, fmt.Sprintf("%v", i))
	}
	return strings.Join(result, " ")
}

func (p NodePrettyPrinter) printInteger(vals []string) string {
	return vals[0]
}
