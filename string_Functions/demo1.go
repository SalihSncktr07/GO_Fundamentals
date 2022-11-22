package string_functions

import (
	"fmt"
	s "strings" // string yerinde "s" kullanabilmek için
)

func Demo1() {
	isim := "Trabzon"
	fmt.Println(s.Count(isim, "a"))    // kaç defa olduğunu bulur
	fmt.Println(s.Contains(isim, "a")) // varsa T, yoksa F verir
	fmt.Println(s.Index(isim, "a"))    //kaçıncı indexde olduğunu verir

	sonuc := s.Index(isim, "a")

	if sonuc != -1 {
		fmt.Println("Içerisinde aranan harf Var")
	} else {
		fmt.Println("Içerisinde aranan harf Yok")
	}

	fmt.Println(s.ToLower(isim))
	fmt.Println(s.ToUpper(isim))
	fmt.Println(isim)
}
