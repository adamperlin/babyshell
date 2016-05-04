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
  "readline"
  "history"
)
var historylist *history.HistoryList
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
        status = shell_util.Shell_exec(cl, historylist)
      }
    }
}
//TODO: not the best solution, but it does work. Look into a better way of doing this, instead of getting a new instance every time
func getNewLineParser(prompt string) *parser.Parser {
  if (historylist == nil){
    historylist = history.CreateNewHistoryInstance()
  }
   reader := readline.ReadLine(prompt)
   if reader != nil {
     historylist.MainList.Append(reader.String())
   return parser.NewParser(reader)
 }else {
   historylist.MainList.Append("")
   return nil
 }
}
