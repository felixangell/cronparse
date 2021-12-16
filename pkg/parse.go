package parse

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
}

type UnitKind int

const (
	List     = iota // a,b,c
	Range           // a - b
	Wildcard        // *
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

func ParseCronString(input string) (*CronExpressionNode, error) {
	return nil, nil
}
