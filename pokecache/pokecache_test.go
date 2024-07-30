package pokecache

import (
	"testing"
	"time"
)

func TestCreatingCache(t *testing.T) {
	cache := NewCache(time.Millisecond)
	if cache.cache == nil {
		t.Error("cache is nil")
	}
}

func TestAddingCache(t *testing.T) {

	cache := NewCache(time.Millisecond)

	cases := []struct {
		input string
		val   []byte
	}{
		{
			input: "Entry",
			val:   []byte("Hello world"),
		},
		{
			input: "Entry2",
			val:   []byte("Hello world2"),
		},
	}

	for _, cas := range cases {
		cache.Add(cas.input, cas.val)
		actual, ok := cache.Get(cas.input)
		if !ok {
			t.Error("cache couldnt add")
		}
		if string(actual) != string(cas.val) {
			t.Errorf("cache failing in retriving: %s != %s", string(actual), string(cas.val))
		}
	}
}
