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
