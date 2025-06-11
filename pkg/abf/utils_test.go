package abf

import (
	"reflect"
	"testing"
)

type zipper [][]any

func TestZip(t *testing.T) {
	tests := []struct {
		name     string
		input    zipper
		expected zipper
	}{
		{
			name:     "two slices of equal length",
			input:    zipper{{1, 2, 3}, {4, 5, 6}},
			expected: zipper{{1, 4}, {2, 5}, {3, 6}},
		},
		{
			name:     "three slices of equal length",
			input:    zipper{{1, 2}, {"a", "b"}, {true, false}},
			expected: zipper{{1, "a", true}, {2, "b", false}},
		},
		{
			name:     "single slice",
			input:    zipper{{1, 2, 3}},
			expected: zipper{{1}, {2}, {3}},
		},
		{
			name:     "empty slices",
			input:    zipper{{}, {}},
			expected: zipper{},
		},
		{
			name:     "slices of different lengths - shorter first",
			input:    zipper{{1, 2}, {4, 5, 6}},
			expected: zipper{{1, 4}, {2, 5}},
		},
		{
			name:     "slices of different lengths - longer first",
			input:    zipper{{1, 2, 3}, {4, 5}},
			expected: zipper{{1, 4}, {2, 5}},
		},
		{
			name:     "one empty slice with non-empty slice",
			input:    zipper{{}, {1, 2, 3}},
			expected: zipper{},
		},
		{
			name:     "multiple mixed types",
			input:    zipper{{1, 2.5, "hello"}, {"world", 42, true}},
			expected: zipper{{1, "world"}, {2.5, 42}, {"hello", true}},
		},
		// ADDED: Previously separate test cases
		{
			name:     "string slices",
			input:    zipper{{"a", "b", "c"}, {"x", "y", "z"}},
			expected: zipper{{"a", "x"}, {"b", "y"}, {"c", "z"}},
		},
		{
			name:     "integer slices",
			input:    zipper{{1, 2, 3}, {10, 20, 30}, {100, 200, 300}},
			expected: zipper{{1, 10, 100}, {2, 20, 200}, {3, 30, 300}},
		},
		{
			name:     "with nil slice",
			input:    zipper{{1, 2}, nil, {3, 4}},
			expected: zipper{},
		},
		{
			name:     "nil argument",
			input:    nil,
			expected: zipper{},
		},
		{
			name:     "as nil slice",
			input:    zipper{},
			expected: zipper{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Zip(tt.input...)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Zip(%v) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestZipNoArgs(t *testing.T) {
	result := Zip[any, []any]()
	expected := zipper{}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Zip() = %v, want %v", result, expected)
	}
}

func TestZipLargeSlices(t *testing.T) {
	// Test with larger slices
	size := 1000
	input1 := make([]any, size)
	input2 := make([]any, size)
	expected := make(zipper, size)

	for i := range size {
		input1[i] = i
		input2[i] = i * 2
		expected[i] = []any{i, i * 2}
	}

	result := Zip(input1, input2)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Zip large slices failed")
	}
}

// Benchmark tests
func BenchmarkZipSmall(b *testing.B) {
	input1 := []any{1, 2, 3}
	input2 := []any{4, 5, 6}

	b.ResetTimer()
	for b.Loop() {
		Zip(input1, input2)
	}
}

func BenchmarkZipMedium(b *testing.B) {
	size := 100
	input1 := make([]any, size)
	input2 := make([]any, size)

	for i := range size {
		input1[i] = i
		input2[i] = i * 2
	}

	b.ResetTimer()
	for b.Loop() {
		Zip(input1, input2)
	}
}

func BenchmarkZipLarge(b *testing.B) {
	size := 1000
	input1 := make([]any, size)
	input2 := make([]any, size)

	for i := range size {
		input1[i] = i
		input2[i] = i * 2
	}

	b.ResetTimer()
	for b.Loop() {
		Zip(input1, input2)
	}
}
