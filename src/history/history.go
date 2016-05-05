package history

import (
	"containers/lists"
)

type HistoryList struct {
	MainList *lists.List
}

func CreateNewHistoryInstance() *HistoryList {
	return &HistoryList{MainList: lists.New()}
}
func (this *HistoryList) AddEntry(str string) {
	this.MainList.Append(str)
}
func (this *HistoryList) CurrentEntry() *lists.Node {
	return this.MainList.End()
}
