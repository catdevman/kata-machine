package katamachine

import (
	"testing"
)

func TestDoublyLinkedList(t *testing.T) {
	list := DoublyLinkedList[int]{}
	list.append(5)
	list.append(7)
	list.append(9)

	if list.get(2) != 9 {
		t.Fatal("second element not equal to 9")
	}
	if list.removeAt(1) != 7 {
		t.Fatal("removeAt(1) does not return 7")
	}
	if list.length != 2 {
		t.Fatal("length not equal to 2")
	}

	list.append(11)

	if list.removeAt(1) != 9 {
		t.Fatal("removing index 1 did not return 9")
	}
	if list.remove(9) != 0 {

		t.Fatal("removing value 9 did not return index 0")
	}

	if val := list.removeAt(0); val != 5 {
		t.Fatal("removing index 0 did not return value of 5", val)
	}

	if val := list.removeAt(0); val != 11 {
		t.Fatal("removing index 0 did not return value of 11", val)
	}

	if list.length != 0 {
		t.Fatal("list length did not equal 0")
	}

	list.prepend(5)
	list.prepend(7)
	list.prepend(9)

	if list.get(2) != 5 {
		t.Fatal("2 index was not equal to 5")
	}

	if list.get(0) != 9 {
		t.Fatal("0 index was not equal to 9")
	}

	if val := list.remove(9); val != 9 {
		t.Fatal("did not remove value 9 from list instead got", val)
	}

	if list.length != 2 {
		t.Fatal("list length did not equal 2")
	}

	if list.get(0) != 7 {
		t.Fatal("0 index was not equal to 7")
	}
}
