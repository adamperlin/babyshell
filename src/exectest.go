package main
import (
  "fmt"
  "os/exec"
  "io/ioutil"
)

func main(){
  cmd := exec.Command("sudo", "pacman", "-Syu")
  pipe, err := cmd.StdoutPipe()
  if err != nil {
    panic(err)
  }
  if err := cmd.Start(); err != nil {
    fmt.Println(err)
  }
  if err := cmd.Wait(); err != nil {
    fmt.Println(err)
  }
  if dat,err := ioutil.ReadAll(pipe); err == nil {
    print(string(dat))
  }
}
