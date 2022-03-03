package inkcallerlib

import "testing"

func TestParseTurnIndex(t *testing.T) {
	turnIndex, err := ParseTurnIndex("3")
	if err != nil {
		t.Fatal(err)
	}
	if *turnIndex != 3 {
		t.Fatal(*turnIndex)
	}
	s := turnIndex.String()
	if s != "3" {
		t.Fatal(s)
	}
}
