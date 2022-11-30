package linked_list

import (
	"data_structures/errors"
	"fmt"
	"reflect"
	"strings"
)

type V interface {
	string | any
}

type SingleLinkedList[T any] struct {
	head *singleNode[T]
	size int
}

type singleNode[T any] struct {
	t    T
	next *singleNode[T]
}

func (s *SingleLinkedList[T]) AddFirst(t T) error {
	return s.Add(0, t)
}

func (s *SingleLinkedList[T]) AddLast(t T) error {
	return s.Add(s.size, t)
}

func (s *SingleLinkedList[T]) RemoveFirst() (*T, error) {
	return s.Remove(0)
}

func (s *SingleLinkedList[T]) RemoveLast() (*T, error) {
	return s.Remove(s.Len()-1)
}

func (s *SingleLinkedList[T]) Add(i int, t T) error {
	if i < 0 {
		return errors.IndexOutOfRange{}
	}
	node := singleNode[T]{t: t}
	if i == 0 {
		node.next = s.head
		s.head = &node
		s.size++
		return nil
	}
	j := 0
	runner := s.head
	for ; runner != nil && j < i-1; runner = runner.next {
		j++
	}
	if runner == nil {
		return errors.IndexOutOfRange{}
	}
	s.size++
	node.next = runner.next
	runner.next = &node
	return nil
}

func (s *SingleLinkedList[T]) Remove(i int) (*T, error) {
    if s.Empty() || i < 0 || i >= s.Len() {
        return nil, errors.IndexOutOfRange{}
    }
    if i == 0 {
        h := s.head
        s.size--
        s.head = s.head.next
        return &h.t, nil
    }
    runner := s.head
    for j := 0; j < i-1; j++ {
        runner = runner.next
    }
    n := runner.next
    runner.next = runner.next.next
    s.size--
	return &n.t, nil
}

func (s *SingleLinkedList[T]) Len() int {
	return s.size
}

func (s *SingleLinkedList[T]) Empty() bool {
	return s.Len() <= 0
}

func (s *SingleLinkedList[T]) String() string {
	var sb strings.Builder
    sb.WriteString("[")
    divider := ","
	for runner := s.head; runner != nil; runner = runner.next {
		t := reflect.ValueOf(runner.t)
		sb.WriteString(fmt.Sprint(t))
        if runner.next != nil {
            sb.WriteString(divider)
        }
	}
    sb.WriteString("]")
	return sb.String()
}

func (s singleNode[T]) String() string {
	return "singleNode"
}
