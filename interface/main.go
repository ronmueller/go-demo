package main

import "fmt"

type myinterface interface {
	my()
}

type otherinterface interface {
	other()
}

type foo struct{}

func (f foo) my() {
	fmt.Println("foo.my")
}

func (f foo) other() {
	fmt.Println("my.other")
}

type bar struct{}

func (b bar) my() {
	fmt.Println("bar.my")
}

func main() {

	f := foo{}
	b := bar{}

	f.my()
	f.other()

	b.my()

	var myi []myinterface

	if myi == nil {
		fmt.Println("myi is nil")
	}

	myi = append(myi, f)
	myi = append(myi, b)

	_ = myi

	for _, x := range myi {
		if y, ok := x.(otherinterface); ok {
			y.other()
		}

		switch y := x.(type) {
		case otherinterface:
			y.other()
		case myinterface:
			fmt.Println("i am not an otherinterface but a myinterface")
		default:
			fmt.Println("what the heck!")
		}
	}

}
