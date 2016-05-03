package lists
//a simple linked list implementation
import "errors"


type Node struct {
  Data interface{}
  next *Node
  prev *Node
}
type List struct {
  head *Node
  current *Node
  previous *Node
}

func New() *List {
  return new(List)
}
func(l *List) Append(val interface{}) {
  if l.head == nil {
    l.head = &Node{next: nil, prev: nil, Data: val}
    l.current = l.head
  }else {
    l.MoveToLast()
    l.previous = l.current
    l.current.next  = &Node{next: nil, prev: l.previous, Data: val}
    l.current = l.current.Next()

  }
}

func (l *List) MoveToLast(){
  temp := l.head
  for temp.next != nil {
    temp = temp.next
  }
  l.current = temp
  l.previous = temp.prev
}

func (l *List) MoveToFirst(){
  l.current = l.head
  l.previous = nil
}
func (l *List) InsertAt(index int, val interface{}){
  temp := l.head
  for i := 0; i < index; i++{
    temp = temp.next
    if temp == nil {
      errors.New("Error: Overflow, attempt to insert value out of bounds of list. To add a new value and extend the list, use List.append")
      return
    }
  }
  //finish
if temp == l.head {
  old := l.head
  l.head = &Node{prev: nil, next: old.next, Data: val}
  return
}
l.current = temp
l.previous = temp.prev
n := &Node{next: nil, prev: nil, Data: val}
l.previous.next = n
l.current.next.prev = n
n.next = l.current.next
}
func (l *List) Begin() *Node {
  return l.head
}

func (this *Node) Next() *Node {
  if elem := this.next; elem == nil {
    return nil
  }
  return this.next
}

func (this *List) End() *Node {
  this.MoveToLast()
  return this.current
}

func (this *Node) Prev() *Node {
  if elem := this.prev; elem == nil {
    return nil
  }
  return this.prev
}

func (l *List) RemoveAt(index int){
  temp := l.head
  for i := 0; i != index; i++{
    temp = temp.next
    if temp == nil {
      errors.New("Error: Overflow, attempt to remove value outside of the list")
      return
    }
  }
    if temp == l.head {
      l.head = l.head.next
      return
    }
    l.current = temp
    l.previous = temp.prev
    //n := &Node{next: nil, prev: nil, Data: val}
    l.previous.next = l.current.next
    l.current.next.prev = l.current.prev
    l.current = l.current.next
}

func (l *List) Push_front(val interface{}){
  p := &Node{prev: nil, next: l.head, Data: val}
  l.current = p
  l.previous = nil
  l.head =  p
}
