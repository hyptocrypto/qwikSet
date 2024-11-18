# Abstract

Go does not have a built-in set data structure. While maps can be used to simulate sets, they are often overkill if we just need to track a set of integers efficiently.
<br/>
<br/>
The idea here is to use a uint64 to represent a set of unsigned integers. Each bit in the uint64 corresponds to an integer: setting the n-th bit marks the integer n as being part of the set. This allows us to compactly represent integers in the range [0, 63].

## Results

Simple benchmarks show that this greatly out preforms maps for basic operations.
`go test -bench=.`

BenchmarkSet_Add-8 138949357 8.624 ns/op
BenchmarkMap_Add-8 19391040 54.42 ns/op
BenchmarkSet_Contains-8 155588612 7.730 ns/op
BenchmarkMap_Contains-8 24343856 48.92 ns/op
BenchmarkSet_Remove-8 145143080 8.274 ns/op
BenchmarkMap_Remove-8 42810177 23.52 ns/op
