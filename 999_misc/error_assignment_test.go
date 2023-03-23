package main

import (
	"errors"
	"testing"
)

func dangerousThing(reallyDangerous bool) (int, error) {
	if reallyDangerous {
		return 1, errors.New("That didn't go well!")
	}

	return 42, nil
}

func Test_NewAssignment(t *testing.T) {
	score, err := dangerousThing(true)
	assertGotError(t, score, err)
}

func Test_PredeclaredAssignment_Error(t *testing.T) {
	var err error
	score, err := dangerousThing(true)
	assertGotError(t, score, err)
}

func Test_PredeclaredAssignment_NoError(t *testing.T) {
	var err error
	score, err := dangerousThing(false)
	assertNoError(t, score, err)
}

func assertGotError(t *testing.T, score int, err error) {
	if err == nil {
		t.Error("Expected error not to be nil, but it was nil")
	}
	if score != 1 {
		t.Error("Expected score not to be 1, but it was not")
	}
}

func assertNoError(t *testing.T, score int, err error) {
	if err != nil {
		t.Error("Expected error not be nil, but it was not nil")
	}
	if score != 42 {
		t.Error("Expected score not to be 42, but it was not")
	}
}
