package pkg

import (
	"reflect"
	"testing"
)

func TestZip(t *testing.T) {
	tests := []struct {
		name     string
		input    [][]any
		expected [][]any
	}{
		{
			name:     "two slices of equal length",
			input:    [][]any{{1, 2, 3}, {4, 5, 6}},
			expected: [][]any{{1, 4}, {2, 5}, {3, 6}},
		},
		{
			name:     "three slices of equal length",
			input:    [][]any{{1, 2}, {"a", "b"}, {true, false}},
			expected: [][]any{{1, "a", true}, {2, "b", false}},
		},
		{
			name:     "single slice",
			input:    [][]any{{1, 2, 3}},
			expected: [][]any{{1}, {2}, {3}},
		},
		{
			name:     "empty slices",
			input:    [][]any{{}, {}},
			expected: [][]any{},
		},
		{
			name:     "slices of different lengths - shorter first",
			input:    [][]any{{1, 2}, {4, 5, 6}},
			expected: [][]any{{1, 4}, {2, 5}},
		},
		{
			name:     "slices of different lengths - longer first",
			input:    [][]any{{1, 2, 3}, {4, 5}},
			expected: [][]any{{1, 4}, {2, 5}},
		},
		{
			name:     "one empty slice with non-empty slice",
			input:    [][]any{{}, {1, 2, 3}},
			expected: [][]any{},
		},
		{
			name:     "multiple mixed types",
			input:    [][]any{{1, 2.5, "hello"}, {"world", 42, true}},
			expected: [][]any{{1, "world"}, {2.5, 42}, {"hello", true}},
		},
		// ADDED: Previously separate test cases
		{
			name:     "string slices",
			input:    [][]any{{"a", "b", "c"}, {"x", "y", "z"}},
			expected: [][]any{{"a", "x"}, {"b", "y"}, {"c", "z"}},
		},
		{
			name:     "integer slices",
			input:    [][]any{{1, 2, 3}, {10, 20, 30}, {100, 200, 300}},
			expected: [][]any{{1, 10, 100}, {2, 20, 200}, {3, 30, 300}},
		},
		{
			name:     "with nil slice",
			input:    [][]any{{1, 2}, nil, {3, 4}},
			expected: [][]any{},
		},
		{
			name:     "nil argument",
			input:    nil,
			expected: [][]any{},
		},
		{
			name:     "as nil slice",
			input:    [][]any{},
			expected: [][]any{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := zip(tt.input...)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("zip(%v) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestZipNoArgs(t *testing.T) {
	result := zip()
	expected := [][]any{}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("zip() = %v, want %v", result, expected)
	}
}

func TestZipLargeSlices(t *testing.T) {
	// Test with larger slices
	size := 1000
	input1 := make([]any, size)
	input2 := make([]any, size)
	expected := make([][]any, size)

	for i := 0; i < size; i++ {
		input1[i] = i
		input2[i] = i * 2
		expected[i] = []any{i, i * 2}
	}

	result := zip(input1, input2)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("zip large slices failed")
	}
}

// Benchmark tests
func BenchmarkZipSmall(b *testing.B) {
	input1 := []any{1, 2, 3}
	input2 := []any{4, 5, 6}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		zip(input1, input2)
	}
}

func BenchmarkZipMedium(b *testing.B) {
	size := 100
	input1 := make([]any, size)
	input2 := make([]any, size)

	for i := 0; i < size; i++ {
		input1[i] = i
		input2[i] = i * 2
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		zip(input1, input2)
	}
}

func BenchmarkZipLarge(b *testing.B) {
	size := 1000
	input1 := make([]any, size)
	input2 := make([]any, size)

	for i := 0; i < size; i++ {
		input1[i] = i
		input2[i] = i * 2
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		zip(input1, input2)
	}
}
