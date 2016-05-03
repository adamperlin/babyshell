package lists
//a simple linked list implementation
import "errors"
import "fmt"

type node struct {
  data interface{}
  next *node
  prev *node
}
type List struct {
  head *node
  current *node
  previous *node
}

func New() *List {
  return &List{}
//  return new(List)
}

func(l *List) Append(val interface{}) {
  if l.head == nil {
    l.head = &node{next: nil, prev: nil, data: val}
  //  fmt.Println(*l.head)
    l.current = l.head
  }else {
    l.MoveToLast()
    l.current.next = new(node)
    l.current.prev = l.current
    l.current = l.current.next
    l.current.data = val
  }
}

func (l *List) MoveToLast(){
  temp := l.head
  for temp.next.next != nil {
    temp = temp.next
  }
  l.previous = temp
  l.current = temp.next
}
func (l *List) Insert(index int, val interface{}){
  temp := l.head
  for i := 0; i <= index; i++{
    temp = temp.next
    if temp == nil {
      errors.New("Error: Overflow, attempt to insert value out of bounds of list. To add a new value and extend the list, use List.append")
      return
    }
  }

  //finish
  l.current = temp
  l.previous = temp.prev
  l.current = l.previous
  l.previous = l.previous.prev
  l.current.next = new(node)
}
