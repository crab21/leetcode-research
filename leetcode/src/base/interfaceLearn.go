package base

import (
	"linkedlist"
)

type Duck interface {
	Quack()
}
type Cat struct {
	name string
}

func (car *Cat) Quack() {
	print(car.name)
}

func T1() {
	var ca Duck = &Cat{name: "wang"}
	ca.Quack()
	ca.(*Cat).Quack()
	linkedlist.TTTMall()
}
