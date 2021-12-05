package main

import "testing"

func TestMakeRange(t *testing.T) {
	expectedIncremental := []int{1, 2, 3, 4, 5}
	expectedDecremental := []int{5, 4, 3, 2, 1}

	outputIncremental := makeRange(1, 5)
	outputDecremental := makeRange(5, 1)

	if len(outputIncremental) != len(expectedIncremental) {
		t.Fail()
		t.Logf("expected: %d, got: %d\n", len(expectedIncremental), len(outputIncremental))
	}

	if len(outputDecremental) != len(expectedDecremental) {
		t.Fail()
		t.Logf("expected: %d, got: %d\n", len(expectedIncremental), len(outputIncremental))
	}

	for i := range outputIncremental {
		if outputIncremental[i] != expectedIncremental[i] {
			t.Fail()
			t.Logf("expected: %d, got: %d\n", expectedIncremental[i], outputIncremental[i])
		}
	}

	for i := range outputDecremental {
		if outputDecremental[i] != expectedDecremental[i] {
			t.Fail()
			t.Logf("expected: %d, got: %d\n", expectedDecremental[i], outputDecremental[i])
		}
	}
}
