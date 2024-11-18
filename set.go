package main

import "math"

type bucket struct {
	idx int
	set int64
}

type Set struct {
	buckets []bucket
}

func (s *Set) addBucket() {
	s.buckets = append(s.buckets, bucket{len(s.buckets) - 1, 0})
}

func (s *Set) getBucket(i int64) int {
	if i == 64 {
		return 0 // This feels like a hack. But 64 should still be in the first bucket
	}
	return int(math.Floor(float64(i) / 64))
}

func (s *Set) removeBucket(idx int64) {}

func (s *Set) Contains(i int) bool {
	return true
}

func (s *Set) Add(i int64) {
	bIdx := s.getBucket(i)
	if len(s.buckets) < bIdx {
		s.addBucket()
	}
	set := s.buckets[bucketIdx].set
	s.buckets[bucketIdx].set = set | int64(1)<<i
}

func (s *Set) Remove(i int) {}

func NewSet() *Set {
	return &Set{buckets: []bucket{{0, 0}}}
}
