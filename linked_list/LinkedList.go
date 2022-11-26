package linked_list

type LinkedList[T any] interface {
    AddFirst(t T)
    AddLast(t T)
    RemoveFirst(t T) T
    RemoveLast(t T) T
    Add(i int, t T)
    Remove(i int) T
    Len() int
    Empty() bool
}
