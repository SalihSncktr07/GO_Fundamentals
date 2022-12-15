package defer_statment

import "fmt"

func A() { // A
	fmt.Println("a fonk çalıştı")
}

func B() { // B C A
	defer A()
	defer C()
	fmt.Println("b fonk çalıştı")
}

func C() { // C
	fmt.Println("c fonk çalıştı")
}
