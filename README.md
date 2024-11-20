# Abstract

Go does not have a built-in set data structure. While maps can be used to simulate sets, they are often overkill if we just need to track a set of integers efficiently.
<br/>
<br/>
The idea here is to use a uint to represent a set of unsigned integers. Each bit in the uint64 corresponds to an integer: setting the n-th bit marks the integer n as being part of the set. This allows us to compactly represent integers in the range [0, 63].

## Results

Simple benchmarks show that this greatly out preforms maps for basic operations.
`go test -bench=.`

```plaintext
BenchmarkSet_Add-8            131569177          8.737 ns/op
BenchmarkMap_Add-8            20781769         55.25 ns/op
BenchmarkSet_Contains-8       147043798          8.173 ns/op
BenchmarkMap_Contains-8       24885315         50.67 ns/op
BenchmarkSet_Remove-8         129638083          9.048 ns/op
BenchmarkMap_Remove-8         20823296         56.79 ns/op
BenchmarkSetIntersection-8          38   29161691 ns/op
BenchmarkMapIntersection-8           2  569464792 ns/op
```
