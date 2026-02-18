package main

import (
	"testing"
)

func TestSliceContains(t *testing.T) {
	slice := []string{"a", "b", "c"}

	if !sliceContains(slice, "a") {
		t.Errorf("Expected slice to contain 'a', but it didn't")
	}

	if sliceContains(slice, "d") {
		t.Errorf("Expected slice to not contain 'd', but it did")
	}
}

func TestJoinSlices(t *testing.T) {
	newSlice := []string{"d", "e"}
	existingSlice := []string{"a", "b", "c"}

	result := joinSlices(newSlice, existingSlice)

	if len(result) != 5 {
		t.Errorf("Expected result length to be 5, got %d", len(result))
	}

	expected := []string{"a", "b", "c", "d", "e"}
	for i, v := range expected {
		if result[i] != v {
			t.Errorf("Expected result[%d] to be %s, got %s", i, v, result[i])
		}
	}

	// Test with duplicates
	duplicateSlice := []string{"a", "f"}
	result = joinSlices(duplicateSlice, result)

	if len(result) != 6 {
		t.Errorf("Expected result length to be 6, got %d", len(result))
	}

	if result[5] != "f" {
		t.Errorf("Expected last element to be 'f', got %s", result[5])
	}
}
