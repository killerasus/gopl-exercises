# Exercise 2.3

Rewrit e PopCount to use a loop instead of a single expression. Compare the performance of the two versions. (Section 11.4 shows how to compare the performance of different implementations systematically.)

## Running the benchmark

1. `go test -bench=PopCountUnroll`
2. `go test -bench=PopCountLoop`
