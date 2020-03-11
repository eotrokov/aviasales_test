package unit_test

import (
	"aviasales/services"
	"aviasales/store"
	"testing"
)

func TestGetAnnogramEmptyStore(t *testing.T) {
	word := "foobar"
	result := services.GetAnnogram(word)
	if result != nil {
		t.Errorf("GetAnnogram(%s) == %s, want %v", word, result, nil)
	}
}

func TestGetAnnogramExsist(t *testing.T) {
	store.AddToStore([]string{"foobar", "barfoo", "boofar"})
	word := "foobar"
	result := services.GetAnnogram(word)
	if result == nil {
		t.Errorf("GetAnnogram(%s) == %s, want %v", word, result, nil)
	}
}
func TestGetAnnogramNonExsist(t *testing.T) {
	store.AddToStore([]string{"foobar", "barfoo", "boofar"})
	word := "abba"
	result := services.GetAnnogram(word)
	if result != nil {
		t.Errorf("GetAnnogram(%s) == %s, want %v", word, result, nil)
	}
}
