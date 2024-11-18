package main

import (
	"math/rand"
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

func BenchmarkSet_Add(b *testing.B) {
	s := NewSet()
	for i := 0; i < b.N; i++ {
		s.Add(rand.Int63n(1000000))
	}
}

func BenchmarkMap_Add(b *testing.B) {
	m := make(map[int64]struct{})
	for i := 0; i < b.N; i++ {
		m[rand.Int63n(1000000)] = struct{}{}
	}
}

func BenchmarkSet_Contains(b *testing.B) {
	s := NewSet()
	for i := 0; i < 1000000; i++ {
		s.Add(int64(i))
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = s.Contains(rand.Int63n(1000000))
	}
}

func BenchmarkMap_Contains(b *testing.B) {
	m := make(map[int64]struct{})
	for i := 0; i < 1000000; i++ {
		m[int64(i)] = struct{}{}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = m[rand.Int63n(1000000)]
	}
}

func BenchmarkSet_Remove(b *testing.B) {
	s := NewSet()
	for i := 0; i < 1000000; i++ {
		s.Add(int64(i))
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Remove(rand.Int63n(1000000))
	}
}

func BenchmarkMap_Remove(b *testing.B) {
	m := make(map[int64]struct{})
	for i := 0; i < 1000000; i++ {
		m[int64(i)] = struct{}{}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		delete(m, rand.Int63n(1000000))
	}
}
