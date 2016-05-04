package shell
//what currently works: executing commands because of c routine.
//TODO: implement better error messages (take away various debug logs)
//TODO: implement piping and redirection, as well as reading a "&" on the end of a command, meaning to fork and not wait.
//TODO: implement some sort of history, and arrow key up and down to scroll through it.
//TODO: fix organization of project if it grows.
import (
  "fmt"
  "shell_util"
  "parser"
  "shell/colors"
//  nc "github.com/rthornton128/goncurses"
  "readline"

)
func Mainloop() {
  status := true
  var color string = colors.CLR_G
    for status {
      p := getNewLineParser(color + "> "+colors.CLR_N)
      if p == nil {continue}
      if err, cl := p.Parse(); err != nil {
        fmt.Println(err)
        color = colors.CLR_R
        p = getNewLineParser(color + "> " + colors.CLR_N)
      }else if cl == nil  {
        status = true
      }else {
        color = colors.CLR_G
        status = shell_util.Shell_exec(cl)
      }
    }
}
//TODO: not the best solution, but it does work. Look into a better way of doing this, instead of getting a new instance every time
func getNewLineParser(prompt string) *parser.Parser {

   reader := readline.ReadLine(prompt)
   if reader != nil {
   return parser.NewParser(reader)
 }else {
   return nil
 }
}

/*func runWindowLoop(){
  s, err := nc.Init()
  if err != nil {
    panic(err)
  }
  defer s.Delete()
  defer nc.End()
  nc.Raw(false)
  nc.Echo(true)
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
      case nc.KEY_ENTER:
        return
    }
  }
}*/
