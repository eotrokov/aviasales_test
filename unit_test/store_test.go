package unit_test

import (
	"aviasales/store"
	"testing"
)

func TestStore(t *testing.T) {
	words := []string{"foobar", "barfoo", "boofar"}
	store.AddToStore(words)
	result := store.GetStore()
	if len(result) != len(words) {
		t.Errorf("GetStore() == %v, want %v", result, words)
	}
}
