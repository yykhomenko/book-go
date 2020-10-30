package main

import (
	"fmt"
)

type IntList struct {
	Value int
	Tail  *IntList
}

func (list *IntList) Sum() int {
	if list == nil {
		return 0
	}
	return list.Value + list.Tail.Sum()
}

func (list *IntList) String() string {
	if list == nil {
		return ""
	}
	return fmt.Sprintf("%d %s", list.Value, list.Tail.String())
}

func main() {
	list := IntList{1, &IntList{2, &IntList{3, nil}}}
	fmt.Println(list.String())
}
