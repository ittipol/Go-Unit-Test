# Go programming language - Unit test

## Packages

- testify [https://pkg.go.dev/github.com/stretchr/testify#section-readme](https://pkg.go.dev/github.com/stretchr/testify#section-readme)

``` bash
#install testify package
go get github.com/stretchr/testify
```

## Run test
``` bash
# go test {modulte_name}/{package_name}
go test gotest/services

go test gotest/services -v

go test gotest/services -v --cover

# Test all function
go test ./...

go test ./... -v

# Specific test function
go test gotest/services -v --cover -run=TestLogin
```

## Run benchmark
``` bash
# Check all function
go test gotest/services -bench=.

# Specific benchmark function
go test gotest/services -bench=BenchmarkLogin

# Check usage memory
go test gotest/services -bench=BenchmarkLogin -benchmem
```

## More info
[https://blog.logrocket.com/benchmarking-golang-improve-function-performance/](https://blog.logrocket.com/benchmarking-golang-improve-function-performance/)