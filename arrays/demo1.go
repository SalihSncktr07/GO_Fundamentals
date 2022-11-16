package arrays

import "fmt"

func Demo1() {
	var sehirler [5]string

	sehirler[0] = "Ankara"
	sehirler[1] = "İstanbul"
	sehirler[2] = "Antalya"
	sehirler[3] = "İzmir"
	sehirler[4] = "Adana"

	//fmt.Println(sehirler)
	for i := 0; i < 5; i++ {
		fmt.Println(sehirler[i])
	}
}
