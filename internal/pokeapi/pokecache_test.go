package pokeapi

import (
	"testing"
	"time"
)

func TestNewPokeCache(t *testing.T) {
	// create a slice of test case structs
	cases := []struct {
		input    time.Duration
		expected *pokeCache
	}{ // create a slice of test case structs
		{
			input:    5 * time.Second,
			expected: &pokeCache{},
		},
	}

	for _, c := range cases {
		actual := newPokeCache(c.input)
		if actual == nil {
			t.Errorf("FAIL -- newPokeCache returned nil, expected a valid pointer")
			continue
		}
	}
}

func TestAdd(t *testing.T) {
	//create a new pokeCache
	pc := newPokeCache(5 * time.Second)
	pc.Add("test", []byte("test passed"))
	actual := pc.cacheMap["test"].val

	if string(actual) != "test passed" {
		t.Errorf("FAIL -- add failed to add new entry")
	}
}

func TestGet(t *testing.T) {
	//create a new pokeCache
	pc := newPokeCache(5 * time.Second)
	pc.Add("test", []byte("test passed"))
	actual, ok := pc.Get("test")

	if !ok {
		t.Errorf("FAIL -- get failed to retrieve entry")
	}

	if string(actual) != "test passed" {
		t.Errorf("FAIL -- get returned incorrect value")
	}
}
