package commands
import (
  "os"
  "fmt"
  "parser"
  "errors"
  "history"
)
type Builtin struct {
  Name string
  Handler func([]string, interface{}) error
}

func Set(args []string, other interface{}) error {
  if args == nil {
    //output list of set shell variables
  }else if len(args) < 2 {
      //make new entry
  }else {
    //varname = args [0], data = args[1]
  }
  return nil
}

func Cd(args []string, other interface{})error {
  home := os.Getenv("HOME")
  var dir string
  if args == nil {
    dir = ""
  }else {
    dir = args[0]
  }
  if dir == "" || dir == "~" {
    os.Chdir(home)
    return nil
  }else {
    err := os.Chdir(dir)
    if err != nil{
      return err
    }
  return nil
  }
}

func Exit(args []string, other interface{}) error {
  os.Exit(0)
  return nil
}
func Help(args []string, other interface{}) error {
  fmt.Println("Welcome to Shell")
  fmt.Println("This is a simple Command line interpreter. Input commands in order to run them.")
  fmt.Println("Standard shell operations, such as pipes, io redirection, and backgrounding are supported, using '>', '<', '|', and '&'.")
  fmt.Println("Refer to man pages for more information about specific commands.")
  return nil
}

func History(args []string, p interface{}) error {
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
func BuiltinExec(cl *parser.BasicCommand, h *history.HistoryList) error {
  for _, b := range Builtins {
    var args []string
    if cl.Args[0] == b.Name {
      if len(cl.Args) > 1 {
        args = cl.Args[1:]
      }else {
        args = nil
      }
      switch b.Name {
        case "history":
          return b.Handler(args, h)
      }
      return b.Handler(args, nil)
    }
  }
  return errors.New("-sh: builtin not found")
}
