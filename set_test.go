package main

import (
	"testing"
)

func TestGetBucket(t *testing.T) {
	tests := []struct {
		name     string
		input    int64
		expected int
	}{
		{"Bucket for 63", 63, 0},
		{"Bucket for 64", 64, 0},
		{"Bucket for 100", 100, 1},
		{"Bucket for 640", 640, 10},
		{"Bucket for 0", 0, 0},
		{"Bucket for 128", 128, 2},
		{"Bucket for 512", 512, 8},
		{"Bucket for 1023", 1023, 15},
	}

	s := NewSet()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := s.getBucket(tt.input)
			if result != tt.expected {
				t.Errorf("getBucket(%d) = %d; want %d", tt.input, result, tt.expected)
			}
		})
	}
}

func TestSet(t *testing.T) {
	s := NewSet()
	s.Add(10)
	if !s.Contains(10) {
		t.Error("expected set to contain 10")
	}
	s.Add(100)
	if !s.Contains(100) {
		t.Error("expected set to contain 10")
	}
	if s.Contains(11) {
		t.Error("expected set to not contain 11")
	}
	s.Remove(10)
	if s.Contains(10) {
		t.Error("expected set to not contain 10")
	}
}
