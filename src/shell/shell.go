package shell
//what currently works: executing commands because of c routine.
//TODO: implement better error messages (take away various debug logs)
//TODO: implement piping and redirection, as well as reading a "&" on the end of a command, meaning to fork and not wait.
//TODO: implement some sort of history, and arrow key up and down to scroll through it.
//TODO: fix organization of project if it grows.
import (
  "fmt"
  "shell_util"
  "os"
  "parser"
  "shell/colors"
)

func Mainloop() {
  p := getNewStdinParser()
  status := true
  var color string = colors.CLR_G
    for status {
      fmt.Printf("%s>%s ",color,colors.CLR_N)
      if err, cl := p.Parse(); err != nil {
        fmt.Println(err)
        color = colors.CLR_R
        p = getNewStdinParser()
      }else if cl == nil  {
        status = true
      }else {
        color = colors.CLR_G
        status = shell_util.Shell_exec(cl)
      }
    }
}
//TODO: not the best solution, but it does work. Look into a better way of doing this, instead of getting a new instance every time
func getNewStdinParser() *parser.Parser{
   reader, err := os.Open("/dev/stdin")
   if err != nil {
     fmt.Println(err)
     os.Exit(-1)
   }
  return parser.NewParser(reader)
}
