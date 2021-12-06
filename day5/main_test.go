package main

import (
	"reflect"
	"testing"
)

func TestMakeRange(t *testing.T) {
	expectedIncremental := []int{1, 2, 3, 4, 5}
	expectedDecremental := []int{5, 4, 3, 2, 1}

	outputIncremental := makeRange(1, 5)
	outputDecremental := makeRange(5, 1)

	if !reflect.DeepEqual(outputIncremental, expectedIncremental) {
		t.Fatalf("expected: %d, got: %d\n", expectedIncremental, outputIncremental)
	}

	if !reflect.DeepEqual(outputDecremental, expectedDecremental) {
		t.Fatalf("expected: %d, got: %d\n", expectedDecremental, outputDecremental)
	}
}
