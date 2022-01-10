package internal

import (
    "fmt"
    parse "github.com/felixangell/cronparse/pkg"
    "github.com/stretchr/testify/assert"
    "testing"
)

// this is a hacky abstraction to
// keep track of what is written via.
// the pretty printer
type capturedLogger struct {
    history []string
}

func (l *capturedLogger) log(out string) {
    l.history = append(l.history, out)
    fmt.Println(out)
}

// this is the test case being asked for in the
// assignment so i will refer to this as the acceptance test.
func TestPrettyPrintAcceptanceTest(t *testing.T) {
    result, err := parse.ParseCronString("'*/15 0 1,15 * 1-5 /usr/bin/find")
    assert.NoError(t, err)

    logger := &capturedLogger{history: []string{}}
    PrettyPrintCron(*result, logger.log)

    expected := []string{
        "minute        0 15 30 45",
        "hour          0",
        "day of month  1 15",
        "month         1 2 3 4 5 6 7 8 9 10 11 12",
        "day of week   1 2 3 4 5",
        "command       /usr/bin/find",
    }

    assert.Equal(t, expected, logger.history)
}

func TestCanPrettyPrintInterval(t *testing.T) {
    result, err := parse.ParseCronString("'*/15 0 1,15 * 1-5 /usr/bin/find")
    assert.NoError(t, err)

    logger := &capturedLogger{history: []string{}}
    PrettyPrintCron(*result, logger.log)

    assert.Equal(t, "minute        0 15 30 45", logger.history[0])
}

func TestCanPrettyPrintWildcard(t *testing.T) {
    result, err := parse.ParseCronString("* * * * 1-5 /usr/bin/find")
    assert.NoError(t, err)

    logger := &capturedLogger{history: []string{}}
    PrettyPrintCron(*result, logger.log)

    expected := []string{
        "minute        0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29 30 31 32 33 34 35 36 37 38 39 40 41 42 43 44 45 46 47 48 49 50 51 52 53 54 55 56 57 58 59",
        "hour          0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23",
        "day of month  1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29 30 31",
        "month         1 2 3 4 5 6 7 8 9 10 11 12",
        "day of week   1 2 3 4 5",
        "command       /usr/bin/find",
    }

    assert.Equal(t, expected, logger.history)
}

func TestCanPrettyPrintInteger(t *testing.T) {
    result, err := parse.ParseCronString("1 2 3 4 5 /usr/bin/find")
    assert.NoError(t, err)

    logger := &capturedLogger{history: []string{}}
    PrettyPrintCron(*result, logger.log)

    expected := []string{
        "minute        1",
        "hour          2",
        "day of month  3",
        "month         4",
        "day of week   5",
        "command       /usr/bin/find",
    }
    assert.Equal(t, expected, logger.history)
}

func TestCanPrettyPrintListInOrder(t *testing.T) {
    result, err := parse.ParseCronString("1 1,2,3 4,5,6 1,5 3,2 /usr/bin/find")
    assert.NoError(t, err)

    logger := &capturedLogger{history: []string{}}
    PrettyPrintCron(*result, logger.log)

    expected := []string{
        "minute        1",
        "hour          1 2 3",
        "day of month  4 5 6",
        "month         1 5",
        "day of week   2 3",
        "command       /usr/bin/find",
    }
    assert.Equal(t, expected, logger.history)
}

func TestCanPrettyPrintRange(t *testing.T) {
    result, err := parse.ParseCronString("1-59 5-23 1-31 1-12 1-5 /usr/bin/find")
    assert.NoError(t, err)

    logger := &capturedLogger{history: []string{}}
    PrettyPrintCron(*result, logger.log)

    expected := []string{
        "minute        1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29 30 31 32 33 34 35 36 37 38 39 40 41 42 43 44 45 46 47 48 49 50 51 52 53 54 55 56 57 58 59",
        "hour          5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23",
        "day of month  1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29 30 31",
        "month         1 2 3 4 5 6 7 8 9 10 11 12",
        "day of week   1 2 3 4 5",
        "command       /usr/bin/find",
    }
    assert.Equal(t, expected, logger.history)
}
