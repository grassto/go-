package main

import (
	"fmt"
	_ "os"
)

type Phone interface {
	call()
	call1()
}

type NokiaPhone struct {
}

func (nokiaPhone NokiaPhone) call() {
	fmt.Println("I am Nokia, I can call you!")
}
func (nokiaPhone NokiaPhone) call1() {
	fmt.Println("I am Nokia, I can call you!")
}

//
type IPhone struct {
	call2 func()
}

func (iPhone IPhone) call() {
	fmt.Println("I am iPhone, I can call you!")
}
func (iPhone IPhone) call1() {
	fmt.Println("I am iPhone, I can call1 you!")
}

func main() {
	var phone Phone

	phone = new(NokiaPhone)
	phone.call()

	phone = new(IPhone)
	phone.call()

	phone.call1()
	var p1 IPhone
	p1.call2 = func() {

	}
	p1.call2()

	if 9 > 2 {
		return
	} else if 3 > 1 {
	} else {

	}

	fmt.Println("ss")
}
