package commands
import (
  "os"
  "fmt"
  "parser"
  "strings"
  "errors"
)
type Builtin struct {
  Name string
  Handler func(string)error
}
func Cd(dir string)error {
home := os.Getenv("HOME")
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

func Exit(code string) error {
  fmt.Printf("Exiting with code %s ", code)
  os.Exit(0)
  return nil
}
func Help(c string) error {
  fmt.Println("Welcome to Shell")
  fmt.Println("This is a simple Command line interpreter. Input commands in order to run them.")
  fmt.Println("Refer to man pages for more information")
  return nil
}

 var Builtins []Builtin = []Builtin {
  Builtin{Name: "cd", Handler: Cd },
  Builtin{Name: "exit", Handler: Exit},
  Builtin{Name:"help", Handler: Help},
}
func IsBuiltin(bin string) bool {
  for _, b := range Builtins {
    if b.Name == bin {
      return true
    }
  }
  return false
}

func BuiltinExec(cl *parser.BasicCommand) error{
  for _, b := range Builtins {
    if cl.Args[0] == b.Name {
      var args string
      if len(cl.Args) < 2{
        args = ""
      }else {
        args = strings.Join(cl.Args[1:], " ")
      }
      return b.Handler(args)
    }
  }
return errors.New("-sh: builtin not found")
}
