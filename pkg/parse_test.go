package parse_test

import (
	"testing"

	parse "github.com/felixangell/cronparse/pkg"
	"github.com/stretchr/testify/assert"
)

func TestCanParseCronString(t *testing.T) {
	input := "*/15 0 1,15 * 1-5 /usr/bin/find"
	result, err := parse.ParseCronString(input)
	assert.NoError(t, err)
	assert.NotNil(t, result)
}

func TestCanSplitRange(t *testing.T) {
	result, err := parse.ParseCronString("1-15")
	assert.NoError(t, err)

	min, _ := result.GetUnit(parse.Minute)
	assert.Equal(t, parse.Range, min.Kind)
	assert.Equal(t, []string{"1", "15"}, min.Operands)
}

func TestCanParseList(t *testing.T) {
	result, err := parse.ParseCronString("1,5,13")
	assert.NoError(t, err)

	min, _ := result.GetUnit(parse.Minute)
	assert.Equal(t, parse.List, min.Kind)
	assert.Equal(t, []string{"1", "5", "13"}, min.Operands)
}

func TestCanParseWildcard(t *testing.T) {
	result, err := parse.ParseCronString("* * *")
	assert.NoError(t, err)
	for i := 0; i < int(parse.EXPRESSION_INDEX_COUNT); i++ {
		u, exists := result.GetUnit(parse.ExpressionIndex(i))
		if !exists {
			continue
		}
		assert.Equal(t, parse.Wildcard, u.Kind)
	}
}

func TestBadInputFails(t *testing.T) {
	result, err := parse.ParseCronString("this should fail")
	assert.Nil(t, result)
	assert.Error(t, err, "Failed to parse this should fail")
}

func TestCanParseInterval(t *testing.T) {
	result, err := parse.ParseCronString("15/*")
	assert.NoError(t, err)

	min, _ := result.GetUnit(parse.Minute)
	assert.Equal(t, parse.Interval, min.Kind)
	assert.Equal(t, []string{"15"}, min.Operands)
}

func TestCanBuildExpressionNode(t *testing.T) {
	expr := parse.NewExpressionNode()
	expr.SetIndex(parse.Minute, &parse.Unit{
		Operands: []string{"a", "b"},
		Kind:     parse.Range,
	})

	val, _ := expr.GetUnit(parse.Minute)
	assert.Equal(t, parse.Range, val.Kind)
	assert.Equal(t, []string{"a", "b"}, val.Operands)
}
