package main

import (
	"math/rand"
	"reflect"
	"testing"
)

func TestGetBucket(t *testing.T) {
	tests := []struct {
		name     string
		input    int
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
				if got := s.Contains(int(value)); got != want {
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

func TestIntersection(t *testing.T) {
	type test struct {
		name     string
		set1     []int
		set2     []int
		expected []int
	}
	tests := []test{
		{
			name:     "Basic intersection",
			set1:     []int{1, 2, 100},
			set2:     []int{1, 100},
			expected: []int{1, 100},
		},
		{
			name:     "No intersection",
			set1:     []int{1, 2, 3},
			set2:     []int{4, 5, 6},
			expected: []int{},
		},
		{
			name:     "Identical sets",
			set1:     []int{10, 20, 30},
			set2:     []int{10, 20, 30},
			expected: []int{10, 20, 30},
		},
		{
			name:     "One empty set",
			set1:     []int{1, 2, 3},
			set2:     []int{},
			expected: []int{},
		},
		{
			name:     "Both empty sets",
			set1:     []int{},
			set2:     []int{},
			expected: []int{},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			s1 := NewSet()
			s2 := NewSet()
			for _, val1 := range tc.set1 {
				s1.Add(val1)
			}

			for _, val2 := range tc.set2 {
				s2.Add(val2)
			}
			result := s1.Intersection(s2)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("expected: %v, got %v", tc.expected, result)
			}
		})
	}
}

func BenchmarkSet_Add(b *testing.B) {
	s := NewSet()
	for i := 0; i < b.N; i++ {
		s.Add(rand.Intn(1000000))
	}
}

func BenchmarkMap_Add(b *testing.B) {
	m := make(map[int]struct{})
	for i := 0; i < b.N; i++ {
		m[rand.Intn(1000000)] = struct{}{}
	}
}

func BenchmarkSet_Contains(b *testing.B) {
	s := NewSet()
	for i := 0; i < 1000000; i++ {
		s.Add(int(i))
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = s.Contains(rand.Intn(1000000))
	}
}

func BenchmarkMap_Contains(b *testing.B) {
	m := make(map[int]struct{})
	for i := 0; i < 1000000; i++ {
		m[int(i)] = struct{}{}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = m[rand.Intn(1000000)]
	}
}

func BenchmarkSet_Remove(b *testing.B) {
	s := NewSet()
	for i := 0; i < 1000000; i++ {
		s.Add(int(i))
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Remove(rand.Intn(1000000))
	}
}

func BenchmarkMap_Remove(b *testing.B) {
	m := make(map[int]struct{})
	for i := 0; i < 1000000; i++ {
		m[int(i)] = struct{}{}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		delete(m, rand.Intn(1000000))
	}
}

func generateRandomInts(size, maxValue int) []int {
	nums := make([]int, size)
	for i := range nums {
		nums[i] = rand.Intn(maxValue)
	}
	return nums
}

func populateSet(values []int) *Set {
	set := NewSet()
	for _, v := range values {
		set.Add(v)
	}
	return set
}

func populateMap(values []int) map[int]struct{} {
	m := make(map[int]struct{})
	for _, v := range values {
		m[v] = struct{}{}
	}
	return m
}

func mapIntersection(a, b map[int]struct{}) []int {
	intersection := []int{}
	for k := range a {
		if _, exists := b[k]; exists {
			intersection = append(intersection, k)
		}
	}
	return intersection
}

func BenchmarkSetIntersection(b *testing.B) {
	setSize := 10000000
	otherSetSize := 10000000
	maxValue := 100000000

	// Generate test data
	setData := generateRandomInts(setSize, maxValue)
	otherSetData := generateRandomInts(otherSetSize, maxValue)

	// Populate sets
	setA := populateSet(setData)
	setB := populateSet(otherSetData)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = setA.Intersection(setB)
	}
}

func BenchmarkMapIntersection(b *testing.B) {
	setSize := 10000000
	otherSetSize := 10000000
	maxValue := 100000000

	// Generate test data
	setData := generateRandomInts(setSize, maxValue)
	otherSetData := generateRandomInts(otherSetSize, maxValue)

	// Populate maps
	mapA := populateMap(setData)
	mapB := populateMap(otherSetData)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = mapIntersection(mapA, mapB)
	}
}
