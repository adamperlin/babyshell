package main
import (
  "fmt"
  "scanner"
  "os"
)

func main() {
  reader, err := os.Open("/dev/stdin")
  if err != nil {
    fmt.Println(err)
  }
  sc  := scanner.NewScanner(reader)
  for {
      tok, ident := sc.Scan()
      switch tok {
      case scanner.AMPERSAND:
        fmt.Println("Scanned &\n")
      case scanner.GREAT:
        fmt.Println("Scanned >\n")
      case scanner.PIPE:
          fmt.Println("Scanned |")
      case scanner.GREATAMPERSAND:
          fmt.Println("Scanned >&\n")
      case scanner.IDENT:
          fmt.Printf("Scanned word: %s\n", ident)
      case scanner.AMPERSANDGREAT:
          fmt.Println("Scanned &>\n")
      case scanner.GREATGREAT:
          fmt.Println("Scanned >>\n")
      case scanner.LESS:
          fmt.Println("Scanned <\n")
      case scanner.RIGHTARROW:
          fmt.Println("Scanned RIGHTARROW")
          break
      case scanner.LEFTARROW:
        fmt.Println("Scanner LEFTARROW")
        break
      case scanner.UPARROW:
        fmt.Println("Scanned UPARROW")
        break
      case scanner.DOWNARROW:
        fmt.Println("Scanned DOWNARROW")
        break
      }
  }
}
