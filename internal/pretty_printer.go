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

		out(fmt.Sprintf("%-14v%v", dateStrings[idx], prettyPrinter.Print(*unit, exprType)))
	}
	out(fmt.Sprintf("%-14v%v", "command", result.Command))

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

type ByIntValueAscending []string

func (a ByIntValueAscending) Len() int { return len(a) }
func (a ByIntValueAscending) Less(i, j int) bool {
	fst, _ := strconv.Atoi(a[i])
	snd, _ := strconv.Atoi(a[j])
	return fst < snd
}
func (a ByIntValueAscending) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

func (p NodePrettyPrinter) printList(values []string) string {
	sort.Sort(ByIntValueAscending(values))
	return strings.Join(values, " ")
}

func (p NodePrettyPrinter) printRange(values []string) string {
	// NOTE: we don't error check here as the printer
	// assumes that the input is valid and that these
	// checks are done at parse time.
	start, _ := strconv.ParseInt(values[0], 10, 64)
	end, _ := strconv.ParseInt(values[1], 10, 64)

	results := make([]string, end)
	idx := 0
	for i := start; i <= end; i++ {
		results[idx] = fmt.Sprintf("%v", i)
		idx++
	}

	return strings.Trim(strings.Join(results, " "), " ")
}

func (p NodePrettyPrinter) printInterval(values []string, typ parse.ExpressionIndex) string {
	// todo base on length
	interval, _ := strconv.ParseInt(values[0], 10, 64)
	start, end := typ.RangeForType()

	var result []string
	for i := int64(start); i < int64(end); i += interval {
		result = append(result, fmt.Sprintf("%v", i))
	}
	return strings.Join(result, " ")
}

func (p NodePrettyPrinter) printWildcard(_ []string, typ parse.ExpressionIndex) string {
	start, end := typ.RangeForType()
	var result []string
	for i := int64(start); i <= int64(end); i++ {
		result = append(result, fmt.Sprintf("%v", i))
	}
	return strings.Join(result, " ")
}

func (p NodePrettyPrinter) printInteger(values []string) string {
	return values[0]
}
