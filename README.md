# cronparse
## Architecture

    internal/ - stuff internal to the application
    pkg/ - reusable components
    main.go - the entry point

The implementation is split up into two parts:

1. parsing - parsing cron input strings into a well-formed structure
2. pretty printer - prints the given nodes into the desired output as specified
3. main.go - driver that joins together these two components

## Installing
You will need Go installed and Go modules enabled - I used 1.16.2.

## Build Instructions

```bash
$ go get ./...  # grab deps
$ go test ./... # run tests, these should all pass!
$ go build      # builds executable
```

## Running
```bash
$ ./cronparse "*/15 0 1,15,99 * 1-5 /usr/bin/find"
```

## Notes to reviewer on the assignment
This is my Deliveroo takehome test. My approach was to split the problem up into pieces:

1. first we take the input and parse it into a structure. i approached this
by producing a 'cron expression' and cutting each index into 'types'
these types are introduced as concrete types in Go which made the range checking
for pretty print/elaboration of the values quite nice to do

2. from this point onwards, it's a matter of taking the expression node and
pretty printing it. parsing the data and giving it context of the types (what
is represented by what input) made this a matter of iterating over the input
values and pretty printing them, i used a visitor esque pattern to approach this

I mostly followed a test driven approach with this, writing acceptance tests to
verify against the end goal and then testing individual components first, e.g.
the parser, and then writing tests/developing the pretty printer, and then eventually
testing them working together. Unit tests were used where necessary but I didn't feel
the need to write exhaustive unit tests (in the sense of achieving 90%+ code coverage
for example) as I feel like the main paths of behaviour are tested.

I did not entirely follow the CRON spec and some areas are not implemented, e.g.
@yearly, or other shortcuts that exist. Also there is not any constraint checking on ranges
as some of this is non standard, e.g. 0-6 week days vs 0-7 weekdays.

I tried to follow a reasonably idiomatic approach though I am not too familiar with
writing Go in terms of developing good software with idiomatic practices, but rather
for side projects and smaller tools/scripts used in a professional setting.

Additional improvements that could be made:

1. making it compliant to CRON spec
2. handling syntactic sugar e.g. @yearly, @weekly
3. constraint checking on ranges/inputs
4. proper go like command line parsing
5. handling range improvements from a starting range

## assignment notes:
given some input: `./cronparse "*/15 0 1,15 * 1-5 /usr/bin/find"`

should produce output:

```
minute 0 15 30 45
hour 0
day of month 1 15
month 1 2 3 4 5 6 7 8 9 10 11 12
day of week 1 2 3 4 5
command /usr/bin/find
```

#### Time tracking;
I did start this in December and continued it in January over two different sessions.
From the git timestamps i've collected to try keep within the allocated window of time.
Overall this took me around 1hr 30min to 2 hours.

Dec 16th start: 14:50:53 - 15:24:42 (34 minutes)
Jan 10th start: 21:47:43 - 22:42:00 (1 hour)

There is a Git log history which shows my workflow and time spans of development.
