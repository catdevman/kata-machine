package katamachine

import "testing"

func TestArrayList(t *testing.T) {
	list := ArrayList[int]{}
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
		t.Fail()
	}
	if list.remove(9) != 0 {
		t.Fail()
	}

	if val := list.removeAt(0); val != 5 {
		t.Fail()
	}

	if val := list.removeAt(0); val != 11 {
		t.Fail()
	}

	if list.length != 0 {
		t.Fail()
	}

	list.prepend(5)
	list.prepend(7)
	list.prepend(9)

	if list.get(2) != 5 {
		t.Fail()
	}

	if list.get(0) != 9 {
		t.Fail()
	}

	if val := list.remove(9); val != 9 {
		t.Fail()
	}

	if list.length != 2 {
		t.Fail()
	}

	if list.get(0) != 7 {
		t.Fail()
	}

}
