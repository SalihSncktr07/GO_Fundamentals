package for_range

import "fmt"

func Demo1() {
	sehirler := []string{"Istanbul", "Ankara", "Izmir"}
	sehir_sayisi := 0

	for i := 0; i < len(sehirler); i++ {
		fmt.Println(sehirler[i])
	}

	for i, sehir := range sehirler {
		fmt.Println(sehir)
		sehir_sayisi = i
	}
	fmt.Println(sehir_sayisi + 1)
}
