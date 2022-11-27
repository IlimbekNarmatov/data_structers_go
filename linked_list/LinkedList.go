package linked_list

type LinkedList[T any] interface {
    AddFirst(t T) error
    AddLast(t T) error
    RemoveFirst(t T) (*T, error)
    RemoveLast(t T) (*T, error)
    Add(i int, t T) error
    Remove(i int) (*T, error)
    Len() int
    Empty() bool
}
