package linked_list

type SingleLinkedList[T any] struct{
    head *singleNode[T]
}

type singleNode[T any] struct{
    t T
    next *singleNode[T]
}

func (s *SingleLinkedList[T]) AddFirst(t T){
}

func (s *SingleLinkedList[T]) AddLast(t T){
}

func (s *SingleLinkedList[T]) RemoveFirst(t T) T{
    return nil
}

func (s *SingleLinkedList[T]) RemoveLast(t T) T{
    return nil
}

func (s *SingleLinkedList[T]) Add(i int, t T){
    if i < 0 {
        panic(IndexOutOfRange{})
        return
    } 
    node := singleNode[T]{t: t}
    if i == 0 {
        node.next = s.head
        s.head = &node
        return
    }
    j := 0
    runner := s.head
    for ; runner != nil && j < i - 1; runner = runner.next {
        j++
    }
    if runner == nil {
        panic(IndexOutOfRange{})
        return
    }
    node.next = runner.next
    runner.next = node
}

func (s *SingleLinkedList[T]) Remove(i int, t T){
}

func (s *SingleLinkedList[T]) Len() int{
    i := 0
    for runner := head; runner != nil; runner = runner.next {
        i++
    }
    return i 
}

func (s *SingleLinkedList[T]) Empty() bool{
    return s.Len() <= 0
}

func (s *SingleLinkedList[T]) String() string{
    var sb strings.Builder
    sb.WriteString("[")
    divider := ","
    for runner = s.head; runner != nil; runner = runner.next {
        t := reflect.ValueOf(runner.t)
        sb.WriteString(t.String())
        if runner.next != nil {
            sb.WriteString(divider)
        }
    }
    sb.WriteString("]")
    return sb.String()
}
