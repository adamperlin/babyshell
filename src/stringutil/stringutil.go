package stringutil
import "strings"

func Contains(arr []string, delim string)bool{
  for i := range arr {
    if arr[i] == delim || strings.Contains(arr[i], delim){
      return true
    }
  }
  return false
}
