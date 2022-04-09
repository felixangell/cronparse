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

## Notes
Does not entirely follow the CRON spec and some areas are not implemented, e.g.
@yearly, or other shortcuts that exist. Also there is not any constraint checking on ranges
as some of this is non standard, e.g. 0-6 week days vs 0-7 weekdays.
