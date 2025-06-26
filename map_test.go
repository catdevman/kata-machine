package katamachine

import (
	"testing"
)

func TestMap(t *testing.T) {
	m := Map[string, int]{}

	m.set("foo", 55)

	if m.size() != 1 {
		t.Fatal("Map size() should return 1")
	}

	m.set("fool", 75)

	if m.size() != 2 {
		t.Fatal("Map size() should return 2")
	}

	m.set("foolish", 105)

	if m.size() != 3 {
		t.Fatal("Map size() should return 3")
	}

	m.set("bar", 69)

	if m.size() != 4 {
		t.Fatal("Map size() should return 4")
	}

	if v, _ := m.get("bar"); v != 69 {
		t.Fatal("bar should be 69")
	}

	if _, ok := m.get("blaz"); ok {
		t.Fatal("blaz ok value should be false")
	}

	m.delete("barblar")
	if m.size() != 4 {
		t.Fatal("Map size() should return 4")
	}

	m.delete("bar")
	if m.size() != 3 {
		t.Fatal("Map size() should return 3")
	}

	if _, ok := m.get("bar"); ok {
		t.Fatal("bar ok value should be false")
	}

}
