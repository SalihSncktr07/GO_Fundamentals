package string_functions

import (
	"fmt"
	s "strings"
)

func Demo2() {
	sehir := "Antalya"
	fmt.Println(s.HasPrefix(sehir, "Ant")) //bununla başlıyor mu
	fmt.Println(s.HasSuffix(sehir, "ya"))  //bununla bitiyor mu
	fmt.Println(s.Index(sehir, "ta"))      //başladığı index değeri

	harfler := []string{"s", "a", "l", "i", "h"}
	fmt.Println(s.Join(harfler, "*")) //aralara "*" koyar

	fmt.Println(s.Replace(sehir, "a", "M", 2)) // 2 den itibaren gördüğü "a" harflerini "M" ile değiştirir

	fmt.Println(s.Split(sehir, "a")) // bulduğu "a" ları böler

	fmt.Println(s.Repeat(sehir, 4)) //4 defa tekrarlar
}
