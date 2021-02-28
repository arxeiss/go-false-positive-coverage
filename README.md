# Go code coverage and False-Positive percentage result

The basic `go test` with cooperation with `go tool cover` can measure and show overall Code coverage of your code.
However the result can be false positive and show **much higher percentage** of covered code that actually is.

## Try it yourself

Clone this repo and try all 3 cases:

### Step 1 - tests only in `binary` package give 100%

```bash
go test -covermode=count -coverprofile=coverage.out ./...
# ?   	github.com/arxeiss/go-false-positive-coverage	[no test files]
# ok  	github.com/arxeiss/go-false-positive-coverage/binary	0.002s	coverage: 100.0% of statements
# ?   	github.com/arxeiss/go-false-positive-coverage/unary	[no test files]

go tool cover
# github.com/arxeiss/go-false-positive-coverage/binary/operations.go:3:	Add		100.0%
# github.com/arxeiss/go-false-positive-coverage/binary/operations.go:7:	Sub		100.0%
# github.com/arxeiss/go-false-positive-coverage/binary/operations.go:11:	Mul		100.0%
# github.com/arxeiss/go-false-positive-coverage/binary/operations.go:15:	Div		100.0%
# total:									(statements)	100.0%
```

### Step 2 - With more tests the overall coverage goes down

```bash
# First "activate" tests in unary package by renaming .bak file
mv unary/operations_test.go.bak unary/operations_test.go
# Now run the tests
go test -covermode=count -coverprofile=coverage.out ./...
# ?   	github.com/arxeiss/go-false-positive-coverage	[no test files]
# ok  	github.com/arxeiss/go-false-positive-coverage/binary	0.002s	coverage: 100.0% of statements
# ok  	github.com/arxeiss/go-false-positive-coverage/unary	0.002s	coverage: 33.3% of statements


go tool cover
# github.com/arxeiss/go-false-positive-coverage/binary/operations.go:3:	Add		100.0%
# github.com/arxeiss/go-false-positive-coverage/binary/operations.go:7:	Sub		100.0%
# github.com/arxeiss/go-false-positive-coverage/binary/operations.go:11:	Mul		100.0%
# github.com/arxeiss/go-false-positive-coverage/binary/operations.go:15:	Div		100.0%
# github.com/arxeiss/go-false-positive-coverage/unary/operations.go:5:	Abs		100.0%
# github.com/arxeiss/go-false-positive-coverage/unary/operations.go:9:	Sin		0.0%
# github.com/arxeiss/go-false-positive-coverage/unary/operations.go:13:	Cos		0.0%
# total:									(statements)	71.4%
```

### Step 3 - Measure "real" coverage with -coverpkg flag

```bash
# Now run the tests
go test -covermode=count -coverprofile=coverage.out -coverpkg=github.com/arxeiss/go-false-positive-coverage/... ./...
# ?   	github.com/arxeiss/go-false-positive-coverage	[no test files]
# ok  	github.com/arxeiss/go-false-positive-coverage/binary	0.002s	coverage: 44.4% of statements in github.com/arxeiss/go-false-positive-coverage/...
# ok  	github.com/arxeiss/go-false-positive-coverage/unary	0.002s	coverage: 11.1% of statements in github.com/arxeiss/go-false-positive-coverage/...


go tool cover
# github.com/arxeiss/go-false-positive-coverage/binary/operations.go:3:	Add		100.0%
# github.com/arxeiss/go-false-positive-coverage/binary/operations.go:7:	Sub		100.0%
# github.com/arxeiss/go-false-positive-coverage/binary/operations.go:11:	Mul		100.0%
# github.com/arxeiss/go-false-positive-coverage/binary/operations.go:15:	Div		100.0%
# github.com/arxeiss/go-false-positive-coverage/main.go:10:		main		0.0%
# github.com/arxeiss/go-false-positive-coverage/unary/operations.go:5:	Abs		100.0%
# github.com/arxeiss/go-false-positive-coverage/unary/operations.go:9:	Sin		0.0%
# github.com/arxeiss/go-false-positive-coverage/unary/operations.go:13:	Cos		0.0%
# total:									(statements)	55.6%
```
