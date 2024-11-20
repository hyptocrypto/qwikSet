package main

import (
	"unsafe"
)

const intBitSize = int(unsafe.Sizeof(int(0)) * 8)

type Set struct {
	buckets []uint
}

func (s *Set) ensureBucketCapacity(bucketIndex int) {
	// Add enough buckets to cover the requested bucket index
	for len(s.buckets) <= bucketIndex {
		s.buckets = append(s.buckets, 0)
	}
}

func (s *Set) getBucket(i int) int {
	return int(float64(i) / 64)
}

func (s *Set) Contains(i int) bool {
	bIdx := s.getBucket(i)
	if bIdx >= len(s.buckets) {
		return false // If the bucket doesn't exist, the value isn't in the set
	}
	bitIndex := i % intBitSize
	return s.buckets[bIdx]&(uint(1)<<bitIndex) != 0
}

func (s *Set) Add(i int) {
	bIdx := s.getBucket(i)
	s.ensureBucketCapacity(bIdx)
	bitIndex := i % intBitSize
	s.buckets[bIdx] |= uint(1) << bitIndex
}

func (s *Set) Remove(i int) {
	bIdx := s.getBucket(i)
	if bIdx < len(s.buckets) {
		bitIndex := i % intBitSize
		s.buckets[bIdx] &^= uint(1) << bitIndex // Clear the bit
	}
}

func (s *Set) Intersection(otherSet *Set) []int {
	ret := []int{}
	var inc int = 0
	for i := range s.buckets {
		if len(otherSet.buckets) < i {
			continue
		}
		r := s.buckets[i] & otherSet.buckets[i]
		if r > 0 {
			var p int = 0
			for r > 0 {
				if r&1 == 1 {
					ret = append(ret, p+inc)
				}
				r >>= 1
				p++
			}
			inc += intBitSize
		}
	}
	return ret
}

func NewSet() *Set {
	return &Set{buckets: []uint{0}}
}
