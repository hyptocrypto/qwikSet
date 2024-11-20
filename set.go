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
	return int(i / intBitSize)
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
	result := []int{}

	minBuckets := len(s.buckets)
	if len(otherSet.buckets) < minBuckets {
		minBuckets = len(otherSet.buckets)
	}

	for bucketIndex := 0; bucketIndex < minBuckets; bucketIndex++ {
		commonBits := s.buckets[bucketIndex] & otherSet.buckets[bucketIndex]
		if commonBits != 0 {
			bitOffset := bucketIndex * intBitSize
			for bit := 0; commonBits != 0; bit++ {
				if commonBits&1 != 0 {
					result = append(result, bitOffset+bit)
				}
				commonBits >>= 1
			}
		}
	}

	return result
}

func NewSet() *Set {
	return &Set{buckets: []uint{0}}
}
