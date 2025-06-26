package katamachine

import "testing"

func TestMinHeap(t *testing.T) {
	heap := new(MinHeap)
	if heap.length != 0 {
		t.Fatal("heap length was not initlized correctly")
	}

	heap.insert(5)
	heap.insert(3)
	heap.insert(69)
	heap.insert(420)
	heap.insert(4)
	heap.insert(1)
	heap.insert(8)
	heap.insert(7)

	if heap.length != 8 {
		t.Fatal("heap length should be 8")
	}

	if v := heap.delete(); v != 1 {
		t.Fatal("delete should have returned 1")
	}

	if v := heap.delete(); v != 3 {
		t.Fatal("delete should have returned 3")
	}

	if v := heap.delete(); v != 4 {
		t.Fatal("delete should have returned 4")
	}

	if v := heap.delete(); v != 5 {
		t.Fatal("delete should have returned 5")
	}

	if heap.length != 4 {
		t.Fatal("heap length should be 4")
	}

	if v := heap.delete(); v != 7 {
		t.Fatal("delete should have returned 7")
	}

	if v := heap.delete(); v != 8 {
		t.Fatal("delete should have returned 8")
	}

	if v := heap.delete(); v != 69 {
		t.Fatal("delete should have returned 69")
	}

	if v := heap.delete(); v != 420 {
		t.Fatal("delete should have returned 420")
	}

	if heap.length != 0 {
		t.Fatal("heap length should be 0")
	}
}
