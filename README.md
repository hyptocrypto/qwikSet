# Abstract

Go does not really have a set. Maps work. But if we just need to track a set of integers, maps are overkill.
<br/>
The idea here is to use a 64uint to represent a set of unsigned ints. When adding some unsigned int 'n' to the set, we will set the n'th bit of the 64uint set. This means we can represent 1-64 in a single 64uint. Obviously we need to support more than 64 as an upper limit, so we will use a bucketing techniques where each bucket is a 64uint and represent n through n+64.
