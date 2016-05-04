package readline
import (
  "unsafe"
 //"io"
 "bytes"
 )
/*
#include "readln.h"
#cgo LDFLAGS: -lreadline
*/
import "C"
func ReadLine(prompt string) *bytes.Buffer{
  cs := C.CString(prompt)
  defer C.free(unsafe.Pointer(cs))
  str := C.read_line(cs)
  defer C.free(unsafe.Pointer(str))
  ret := C.GoString(str)
  b := bytes.NewBufferString(ret)
  if b != nil {
  b.WriteRune('\n')
  }
  return b
}
