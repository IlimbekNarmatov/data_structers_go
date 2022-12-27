package linked_list

import (
	"strconv"
	"strings"
	"testing"
)

func TestSingleLinkedList_AddFirst(t *testing.T) {
	l := SingleLinkedList[int]{}
	l.AddFirst(0)
	got := l.String()
	want := "[0]"

	if got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}

	l.AddFirst(1)
	got = l.String()
	want = "[1,0]"

	if got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}

	l = SingleLinkedList[int]{}
	for i := 0; i < 10; i++ {
		l.AddFirst(i)
	}
	var wantSB strings.Builder
	wantSB.WriteString("[")
	for i := 9; i > -1; i-- {
		wantSB.WriteString(strconv.Itoa(i))
		if i-1 > -1 {
			wantSB.WriteString(",")
		}
	}
	wantSB.WriteString("]")
	got = l.String()
	want = wantSB.String()

	if got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}
}

func TestSingleLinkedList_AddLast(t *testing.T) {
	l := SingleLinkedList[int]{}
	got := l.String()
	want := "[]"

	if got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}

	l.AddLast(0)
	got = l.String()
	want = "[0]"

	if got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}

	l.AddLast(1)
	got = l.String()
	want = "[0,1]"

	if got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}

	l = SingleLinkedList[int]{}
	var wantSB strings.Builder
	wantSB.WriteString("[")
	for i := 0; i < 10; i++ {
		l.AddLast(i)
		wantSB.WriteString(strconv.Itoa(i))
		if i+1 < 10 {
			wantSB.WriteString(",")
		}
	}
	wantSB.WriteString("]")
	got = l.String()
	want = wantSB.String()

	if got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}
}

func TestSingleLinkedList_RemoveFirst(t *testing.T) {
	l := SingleLinkedList[int]{}
	l.AddFirst(0)
	l.RemoveFirst()
	got := l.String()
	want := "[]"

	if got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}

	l = SingleLinkedList[int]{}
	l.AddFirst(0)
	l.AddFirst(1)
	l.RemoveFirst()
	got = l.String()
	want = "[0]"

	if got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}

	l = SingleLinkedList[int]{}
	for i := 0; i < 10; i++ {
		l.AddLast(i)
	}
	l.RemoveFirst()
	got = l.String()
	want = "[1,2,3,4,5,6,7,8,9]"

	if got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}

	l.RemoveFirst()
	got = l.String()
	want = "[2,3,4,5,6,7,8,9]"

	if got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}

	l.RemoveFirst()
	got = l.String()
	want = "[3,4,5,6,7,8,9]"

	if got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}

	l.RemoveFirst()
	got = l.String()
	want = "[4,5,6,7,8,9]"

	if got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}

	l.RemoveFirst()
	got = l.String()
	want = "[5,6,7,8,9]"

	if got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}
}

func TestSingleLinkedList_RemoveLast(t *testing.T) {
	l := SingleLinkedList[int]{}
	l.AddLast(0)
	l.RemoveLast()
	got := l.String()
	want := "[]"

	if got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}

	l = SingleLinkedList[int]{}
	l.AddLast(0)
	l.AddLast(1)
	l.RemoveLast()
	got = l.String()
	want = "[0]"

	if got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}

	l = SingleLinkedList[int]{}
	for i := 0; i < 10; i++ {
		l.AddLast(i)
	}
	l.RemoveLast()
	got = l.String()
	want = "[0,1,2,3,4,5,6,7,8]"
	if got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}

	l.RemoveLast()
	got = l.String()
	want = "[0,1,2,3,4,5,6,7]"
	if got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}

	l.RemoveLast()
	got = l.String()
	want = "[0,1,2,3,4,5,6]"
	if got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}

	l.RemoveLast()
	got = l.String()
	want = "[0,1,2,3,4,5]"
	if got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}

	l.RemoveLast()
	got = l.String()
	want = "[0,1,2,3,4]"
	if got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}
}

func TestSingleLinkedList_Add(t *testing.T) {
	l := SingleLinkedList[int]{}
	l.Add(0, 1)
	got := l.String()
	want := "[1]"
	if got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}

	l = SingleLinkedList[int]{}
	l.Add(0, 1)
	l.Add(1, 2)
	got = l.String()
	want = "[1,2]"
	if got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}

	l = SingleLinkedList[int]{}
	l.Add(0, 1)
	l.Add(1, 2)
	l.Add(2, 3)
	got = l.String()
	want = "[1,2,3]"
	if got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}

	l.Add(1, 4)
	got = l.String()
	want = "[1,4,2,3]"
	if got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}
}

func TestSingleLinkedList_Remove(t *testing.T) {
	l := SingleLinkedList[int]{}
	l.AddLast(0)
	l.Remove(0)
	got := l.String()
	want := "[]"
	if got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}

	l = SingleLinkedList[int]{}
	l.AddLast(0)
	l.AddLast(1)
	l.Remove(1)
	got = l.String()
	want = "[0]"
	if got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}

	l = SingleLinkedList[int]{}
	l.AddLast(0)
	l.AddLast(1)
	l.AddLast(2)
	l.Remove(0)
	got = l.String()
	want = "[1,2]"
	if got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}

	l = SingleLinkedList[int]{}
	l.AddLast(0)
	l.AddLast(1)
	l.AddLast(2)
	l.Remove(1)
	got = l.String()
	want = "[0,2]"
	if got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}

	l = SingleLinkedList[int]{}
	l.AddLast(0)
	l.AddLast(1)
	l.AddLast(2)
	l.Remove(2)
	got = l.String()
	want = "[0,1]"
	if got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}
}

func TestSingleLinkedList_Len(t *testing.T) {
	l := SingleLinkedList[int]{}
	got := l.Len()
	want := 0
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}

	l = SingleLinkedList[int]{}
	l.AddLast(0)
	got = l.Len()
	want = 1
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}

	l = SingleLinkedList[int]{}
	l.AddLast(0)
	l.AddLast(1)
	got = l.Len()
	want = 2
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestSingleLinkedList_Empty(t *testing.T) {
	l := SingleLinkedList[int]{}
	got := l.Empty()
	want := true
	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}

	l = SingleLinkedList[int]{}
	l.AddLast(0)
	got = l.Empty()
	want = false
	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}

	l = SingleLinkedList[int]{}
	l.RemoveLast()
	got = l.Empty()
	want = true
	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}

	l = SingleLinkedList[int]{}
	l.AddLast(0)
	l.AddLast(1)
	l.RemoveLast()
	got = l.Empty()
	want = false
	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}
}
