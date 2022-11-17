package slices

import "fmt"

func Demo2() {
	sehirler := []string{"Antalya", "Istanbul", "Ankara"}

	fmt.Println(sehirler)

	kopya := make([]string, len(sehirler))
	fmt.Println(kopya)

	copy(kopya, sehirler)
	fmt.Println(kopya)

	sehirler = append(sehirler, "Trabzon")
	fmt.Println(sehirler)
	fmt.Println(kopya)
	copy(kopya, sehirler)
	fmt.Println(kopya) //kopyanın alanı sehirlerin ilk halinin uzunluğu kadar boyut verildiği için daha fazlasını yazamaz
}
