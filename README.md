# cronparse

# installing

# example usage

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

Time trackin

Dec 16th start: 14:50:53 - 15:24:42 (34 minutes)