package main

import (
	"fmt"
	"reflect"
	"sort"
)

type person struct {
	name string
	age  int
}

type persons []person

func (p persons) Len() int {
	return len(p)
}

func (p persons) Less(i, j int) bool {
	return p[i].age < p[j].age
}
func (p persons) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
func main() {
	p1 := persons{
		person{"saksham", 38},
		person{"ali", 20},
	}
	fmt.Println(reflect.ValueOf(p1))
	sort.Sort(p1)
	fmt.Println(p1)

}
