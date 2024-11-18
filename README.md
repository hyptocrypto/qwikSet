# Abstract

Go does not have a built-in set data structure. While maps can be used to simulate sets, they are often overkill if we just need to track a set of integers efficiently.
<br/>
<br/>
The idea here is to use a uint64 to represent a set of unsigned integers. Each bit in the uint64 corresponds to an integer: setting the n-th bit marks the integer n as being part of the set. This allows us to compactly represent integers in the range [0, 63].
