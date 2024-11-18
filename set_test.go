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
		{"Bucket for 64", 64, 1},
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
	tests := []struct {
		name     string
		actions  func(*Set)
		expected map[int]bool // Key is the value to check, value is whether it should be contained in the set
	}{
		{
			name: "Add 0 and 10, check contains",
			actions: func(s *Set) {
				s.Add(0)
				s.Add(10)
			},
			expected: map[int]bool{10: true, 0: true, 11: false},
		},
		{
			name: "Add 100, check contains",
			actions: func(s *Set) {
				s.Add(100)
			},
			expected: map[int]bool{100: true, 11: false},
		},
		{
			name: "Add 100, check contains",
			actions: func(s *Set) {
				s.Add(1000000)
			},
			expected: map[int]bool{1000000: true, 11: false},
		},
		{
			name: "Add 100, check contains",
			actions: func(s *Set) {
				s.Add(1)
				s.Remove(1)
			},
			expected: map[int]bool{1: false},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewSet()
			tt.actions(s)

			for value, want := range tt.expected {
				if got := s.Contains(int64(value)); got != want {
					if want {
						t.Errorf("expected set to contain %d, but it did not", value)
					} else {
						t.Errorf("expected set to not contain %d, but it did", value)
					}
				}
			}
		})
	}
}
