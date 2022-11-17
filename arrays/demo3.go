package arrays

import "fmt"

func Demo3() {
	var sayilar [2][3]int

	sayilar[0][0] = 0
	sayilar[0][1] = 1
	sayilar[0][2] = 2
	sayilar[1][0] = 3
	sayilar[1][1] = 4
	sayilar[1][2] = 5

	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			fmt.Print(sayilar[i][j])
			fmt.Print(" ")
		}
		fmt.Println()
	}
}
