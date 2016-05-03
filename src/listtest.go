 package main
 import (
   "containers/lists"
   "fmt"
 )

 func main(){
   list := lists.New()
/*  for i := 0; i < 100; i++ {
    list.Append(i)
  }
   for i := list.Begin(); i != nil; i = i.Next() {
     fmt.Println(i.Data)
   }
   fmt.Println("**********************************************")
   for i := list.End(); i != nil; i = i.Prev() {
     fmt.Println(i.Data)
   }*/
   list.Append("hello")
   list.Append("bizarre")
   list.Append("world")
   list.RemoveAt(0)
   list.Push_front("Whats good")
  // list.Append("!")
   for i := list.Begin(); i != nil; i = i.Next() {
     fmt.Println(i.Data)
   }
   list.InsertAt(0, "whatup")
   for i := list.Begin(); i != nil; i = i.Next() {
     fmt.Println(i.Data)
   }
 }
