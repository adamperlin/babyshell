package history
import (
  "containers/lists"
  "parser"
  "strings"
)
type HistoryList struct {
  MainList *lists.List
}
func CreateNewHistoryInstance() *HistoryList {
  return &HistoryList{MainList: lists.New()}
}
func (this *HistoryList) AddEntry(lit string){
  str := strings.Join(cl.Args, " ")
  this.MainList.Append(str)
}
func (this *HistoryList) CurrentEntry(){
  return this.MainList.End()
}
