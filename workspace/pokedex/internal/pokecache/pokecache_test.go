package pokecache

import (
	"bytes"
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {
	nc := NewCache(5 * time.Second)

	nc.Add("haha", []byte("hihi"))

	bite, ok := nc.Get("haha")

	if !ok || !bytes.Equal(bite, []byte("hihi")) {
		t.Errorf("Add or Get fucked up")
	}
}

func TestReapLoop(t *testing.T) {
	nc := NewCache(5 * time.Millisecond)

	nc.Add("haha", []byte("hihi"))
	bite, ok := nc.Get("haha")

	if !ok || !bytes.Equal(bite, []byte("hihi")) {
		t.Errorf("Add or Get fucked up")
	}

	time.Sleep(1 * time.Second)

	bite, ok = nc.Get("haha")
	if ok {
		t.Errorf("ReapLoop fucked up")
	}
}
