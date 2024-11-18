package main

import (
	"math"
)

type Set struct {
	buckets []uint64
}

func (s *Set) ensureBucketCapacity(bucketIndex int) {
	// Add enough buckets to cover the requested bucket index
	for len(s.buckets) <= bucketIndex {
		s.buckets = append(s.buckets, 0)
	}
}

func (s *Set) getBucket(i int64) int {
	return int(math.Floor(float64(i) / 64))
}

func (s *Set) Contains(i int64) bool {
	bIdx := s.getBucket(i)
	if bIdx >= len(s.buckets) {
		return false // If the bucket doesn't exist, the value isn't in the set
	}
	bitIndex := i % 64
	return s.buckets[bIdx]&(uint64(1)<<bitIndex) != 0
}

func (s *Set) Add(i int64) {
	bIdx := s.getBucket(i)
	s.ensureBucketCapacity(bIdx)
	bitIndex := i % 64
	s.buckets[bIdx] |= uint64(1) << bitIndex
}

func (s *Set) Remove(i int64) {
	bIdx := s.getBucket(i)
	if bIdx < len(s.buckets) {
		bitIndex := i % 64
		s.buckets[bIdx] &^= uint64(1) << bitIndex // Clear the bit
	}
}

func NewSet() *Set {
	return &Set{buckets: []uint64{0}}
}
