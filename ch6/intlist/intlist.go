package main

import (
	"fmt"
)

type IntList struct {
	Value int
	Tail  *IntList
}

func (list *IntList) Map(f func(int) int) {
	if list == nil {
		return
	}
	list.Value = f(list.Value)
	list.Tail.Map(f)
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
	fmt.Printf("sum = %d\n", list.Sum())
	list.Map(func(i int) int {
		return i * 3
	})
	fmt.Printf("map i*2 = %d\n", list.Sum())
}
