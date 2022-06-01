package main

import "fmt"

type ArithmeticTool interface {
	add(base int) int
	multiply(base int) int
}

type Taschenrechner2000 struct {
	addBy      int
	multiplyBy int
}

func (a *Taschenrechner2000) add(base int) int {
	return a.addBy + base
}

func (a *Taschenrechner2000) multiply(base int) int {
	return a.multiplyBy * base
}

func (a Taschenrechner2000) giveMeACoolNumber() int {
	return a.multiplyBy * a.addBy
}

func doSomeMath() {
	var theTool ArithmeticTool = &Taschenrechner2000{1, 42}

	fmt.Println(theTool.add(42))
	fmt.Println(theTool.multiply(42))
}
