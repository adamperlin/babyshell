package main
import (
  nc "github.com/rthornton128/goncurses"
  "os"
  "fmt"
)

func main(){
  s, err := nc.Init()
  defer s.Delete()
  if err != nil {
    panic(err)
  }
  defer nc.End()
  nc.Raw(true)
  nc.Echo(false)
//  s.Clear()
  s.Keypad(true)


  for {
    ch := s.GetChar()
    switch ch {
    case 'q': os.Exit(0)
  case nc.KEY_LEFT:
      y, x := s.CursorYX()
      s.Move(y, x-1)
      s.Refresh()
    case nc.KEY_RIGHT:
      y, x := s.CursorYX()
      s.Move(y, x+1)
      s.Refresh()
    case nc.KEY_BACKSPACE:
      y, x := s.CursorYX()
      s.Move (y, x - 1)
        err := s.DelChar()
        if err != nil {
          fmt.Println(err)
        }
      s.Refresh()
    }
  }
}
