package commands
import (
  "os"
  "fmt"
  "parser"
  "strings"
  "errors"
  "history"
)
type Builtin struct {
  Name string
  Handler func(interface{}) error
}
func Cd(dir interface{})error {
  home := os.Getenv("HOME")
  strdir, ok := dir.(string)
  if  !ok {
    return errors.New("-sh: error: cd must be given a string as an argument")
  }
  if strdir == "" || strdir == "~" {
    os.Chdir(home)
    return nil
  }else {
    err := os.Chdir(strdir)
    if err != nil{
      return err
    }
  return nil
  }
}

func Exit(code interface{}) error {
  os.Exit(0)
  return nil
}
func Help(c interface{}) error {
  fmt.Println("Welcome to Shell")
  fmt.Println("This is a simple Command line interpreter. Input commands in order to run them.")
  fmt.Println("Refer to man pages for more information")
  return nil
}

func History(p interface{}) error {
  history, ok := p.(*history.HistoryList)
  if !ok {
    return errors.New("Wrong value given to History")
  }else {
    inc := 0
    for i := history.MainList.Begin(); i != history.MainList.End(); i = i.Next(){
      fmt.Printf("%d %s \n", inc,i.Data)
      inc++
    }
  }
  return nil
}

 var Builtins []Builtin = []Builtin {
  Builtin{Name: "cd", Handler: Cd },
  Builtin{Name: "exit", Handler: Exit},
  Builtin{Name:"help", Handler: Help},
  Builtin{Name:"history", Handler: History},
}
func IsBuiltin(bin string) bool {
  for _, b := range Builtins {
    if b.Name == bin {
      return true
    }
  }
  return false
}

func BuiltinExec(cl *parser.BasicCommand, h *history.HistoryList) error{
  for _, b := range Builtins {
    if cl.Args[0] == b.Name {
      switch b.Name {
        case "history":
          return b.Handler(h)
      }
      var args string
      if len(cl.Args) < 2 {
        args = ""
      }else {
        args = strings.Join(cl.Args[1:], " ")
      }
      return b.Handler(args)
    }
  }
return errors.New("-sh: builtin not found")
}
