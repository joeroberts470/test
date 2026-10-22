# Go Hiring Test

```
go-hiring-test/
├── README.md          # This file
├── go.mod             # Module file (no external dependencies)
├── solution.go        # Your implementation goes here
└── test.go            # Test suite (do not modify)
```

## Task
Implement `TransformArray(arr []int) int` that:
1. Filters out negative numbers (< 0)
2. Multiplies remaining numbers by 2
3. Returns the sum

**Examples:**
```go
TransformArray([]int{1, 2, 3, 4})     // → 20
TransformArray([]int{-5, 10, -3, 7})  // → 34
TransformArray([]int{})               // → 0
```

## Running the tests

```
go run .
```

No `go get`, no external packages — standard library only. If you have Go
installed, this just works.

## Time

This should take about **10 minutes**.
