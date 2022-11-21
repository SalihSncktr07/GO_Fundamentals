package defer_statment

import "fmt"

func A() {
	fmt.Println("a fonk çalıştı")
}

func B() {
	defer A()
	defer C()
	fmt.Println("b fonk çalıştı")
}

func C() {
	fmt.Println("c fonk çalıştı")
}
