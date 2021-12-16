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

func TestCanBuildExpressionNode(t *testing.T) {
	expr := parse.NewExpressionNode()
	expr.SetIndex(parse.Minute, parse.Unit{
		Operands: []string{"a", "b"},
		Kind:     parse.Range,
	})

	val := expr.GetUnit(parse.Minute)
	assert.Equal(t, parse.Range, val.Kind)
	assert.Equal(t, []string{"a", "b"}, val.Operands)
}
