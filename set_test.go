package main

import (
	"fmt"
	"testing"
)

func TestGetBucket(t *testing.T) {
	s := NewSet()
	b := s.getBucket(63)
	if b != 0 {
		t.Error("expected to get bucket 0 for 63")
	}
	b = s.getBucket(64)
	if b != 0 {
		t.Error("expected to get bucket 0 for 64")
	}
	b = s.getBucket(100)
	if b != 1 {
		t.Error("expected to get bucket 1 for 100")
	}
	b = s.getBucket(640)
	if b != 10 {
		t.Error("expected to get bucket 10 for 640")
	}
}

func TestSet(t *testing.T) {
	s := NewSet()
	s.Add(10)
	if !s.Contains(10) {
		t.Error("expected set to contain 10")
	}
	fmt.Printf("%b\n", s.buckets[0].set)
}
