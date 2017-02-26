package main

import (
	"testing"
)

func testSum(t *testing.T) {
	if sum(1, 2) !=3 {
		t.Fatal("error")
	}
}