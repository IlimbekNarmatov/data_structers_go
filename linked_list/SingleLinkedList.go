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

func (s *SingleLinkedList[T]) RemoveFirst(t T) (*T, error) {
	return nil, nil
}

func (s *SingleLinkedList[T]) RemoveLast(t T) (*T, error) {
	return nil, nil
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
	return nil, nil
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
    divider := ", "
	for runner := s.head; runner != nil; runner = runner.next {
		t := reflect.ValueOf(runner.t)
		sb.WriteString(fmt.Sprint(t))
        if(runner.next != nil){
            sb.WriteString(divider)
        }
	}
    sb.WriteString("]")
	return sb.String()
}
