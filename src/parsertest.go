package main
import (
  "parser"
  "os"
  "fmt"
)

func main(){
reader, err := os.Open("/dev/stdin")
if err != nil {
  fmt.Println(err)
}else {
  var p *parser.Parser = parser.NewParser(reader)
  for {
    err, cl := p.Parse()
    if err != nil {
      fmt.Println(err)
      break
    }
    if cl != nil {
      fmt.Printf("Stdout: %s\n", cl.Out)
      fmt.Printf("Stdin: %s\n", cl.In)
      fmt.Printf("Stderr: %s\n", cl.Err)
      for i := range cl.Commands {
        fmt.Printf("command %d: ", i)
        for j := range cl.Commands[i].Args {
          fmt.Printf("%s ",cl.Commands[i].Args[j])
        }
      }

    }
  }
}

}
