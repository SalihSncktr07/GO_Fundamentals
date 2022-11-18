package for_range

import "fmt"

func Demo2() {
	sayi := 1
	tek := 0

	for i := sayi; i < 10; i += 2 {		
			tek += i
		}
	fmt.Println(tek)
}