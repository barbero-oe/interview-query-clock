# Simple Alarm

Create a clock program that print the clock time each second and an alarm that
sounds at a predefined time.
The alarm should not block the main process while sounding. After the alarm
sounds once exit the program.

# Tests

The clock module have some tests, to run them execute

```shell
go test -v ./...
```

# Run

The program expects as only argument the desired moment to play the alarm.
The expected argument should be in 24-hour format, divided by colons.
For example:

```shell
./interview-query-clock 20:13:28
```
