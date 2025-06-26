package katamachine

import "testing"

func TestLRU(t *testing.T) {
	lru := LRU[string, int]{}

	if _, ok := lru.get("foo"); ok {
		t.Fatal("unset foo should have a false on ok")
	}

	lru.update("foo", 69)

	if v, ok := lru.get("foo"); !ok || v != 69 {
		t.Fatal("foo should be set to 69 and the ok value should be true ")
	}

	lru.update("bar", 420)

	if v, ok := lru.get("bar"); !ok || v != 420 {
		t.Fatal("bar should be set to 420 and the ok value should be true ")
	}

	lru.update("baz", 1337)

	if v, ok := lru.get("baz"); !ok || v != 1337 {
		t.Fatal("baz should be set to 1337 and the ok value should be true ")
	}

	lru.update("ball", 69420)

	if v, ok := lru.get("ball"); !ok || v != 69420 {
		t.Fatal("ball should be set to 69420 and the ok value should be true ")
	}

	if _, ok := lru.get("foo"); ok {
		t.Fatal("unset foo should have a false on ok")
	}

	if v, ok := lru.get("bar"); !ok || v != 420 {
		t.Fatal("bar should be ok and should have a value of 420")
	}

	lru.update("foo", 69)

	if v, ok := lru.get("bar"); !ok || v != 420 {
		t.Fatal("bar should be ok and should have a value of 420")
	}

	if v, ok := lru.get("foo"); !ok || v != 69 {
		t.Fatal("foo should be ok and should have a value of 69")
	}

	// shouldn't of been deleted, but since bar was get'd, bar was added to the
	// front of the list, so baz became the end
	if _, ok := lru.get("baz"); ok {
		t.Fatal("baz should have ok as false")
	}
}
