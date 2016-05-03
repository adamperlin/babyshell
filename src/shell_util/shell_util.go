package shell_util
import (
  "unsafe"
  "parser"
  "fmt"
  "os"
  "commands"
)
/*
#include "execute.h"
int open2(char* pathname, int flags, int mode){
  if (mode != 0){
    return open(pathname, flags, mode);
  }else {
  return open(pathname, flags);
  }
}
*/
import "C"
func Shell_exec(cl *parser.CommandList) bool {
    // get stdin and stdout file descriptors
      tmpin := C.dup(0)
      tmpout := C.dup(1)
      var in_f C.int
      var out_f C.int
      var pid C.pid_t
      if cl.In != "" {
        temp := C.CString(cl.In)
        defer C.free(unsafe.Pointer(temp))
        in_f = C.open2(temp, C.O_RDONLY,0)
        if in_f == C.int(-1){
          fmt.Fprintf(os.Stderr, "Error: file %s not found \n", temp)
        }
      }else {
        in_f = C.dup(tmpin)
      }
      for i := range cl.Commands {
        C.dup2(in_f, 0)
        C.close(in_f)
        if i == len(cl.Commands) - 1 { //last command
          if cl.Out != "" {
            temp := C.CString(cl.Out)
            defer C.free(unsafe.Pointer(temp))
            out_f = C.open2(temp, C.O_WRONLY|C.O_CREAT, 0700)
            if (out_f == -1){
              fmt.Fprintf(os.Stderr, "-sh: Error: failed to redirect to file")
            }
          }else {
            out_f = C.dup(tmpout)
          }
        }else {
          var fd [2]C.int
          C.pipe(&fd[ 0 ])
          out_f = fd[ 1 ]
          in_f = fd[ 0 ]
        }
        C.dup2(out_f, 1)
        C.close(out_f)
        carr, cleanup := toCStringArray(cl.Commands[ i ].Args)
        defer cleanup(carr)
        if commands.IsBuiltin(cl.Commands[ i ].Args[0]) {
          err := commands.BuiltinExec(cl.Commands[ i ])
          if err != nil {
            fmt.Println(err)
          }
        }else {
        pid = C.process_launch(&carr[ 0 ])
        }
      }
  C.dup2(tmpin, 0)
  C.dup2(tmpout, 1)
  C.close(tmpin)
  C.close(tmpout)
  if !cl.Background {
    err := C.waitfor(pid)
    if err == 1 {
      return true
    } else {
      return false
    }
  }
  return true
}
func toCStringArray(arr []string)([]*C.char, func([]*C.char)){
  carr := make([]*C.char,len(arr))
  for i := range arr {
    carr[i] = C.CString(arr[i])
  }
  carr = append(carr, nil)
  cleanup := func(carr []*C.char){
    for i := range carr { //this memory isn't garbage collected, must manually free
      C.free(unsafe.Pointer(carr[i]))
    }
  }
  return carr, cleanup
}
