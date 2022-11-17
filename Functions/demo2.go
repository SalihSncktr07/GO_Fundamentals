package functions

func DortIslem(sayi1 int, sayi2 int) (int, int, int, float32) {
	sum := sayi1 + sayi2
	diff := sayi1 - sayi2
	imp := sayi1 * sayi2
	div := float32(sayi1 / sayi2)

	return sum, diff, imp, div
}

	/*main
		sonuc1, _ , sonuc3, _ := functions.DortIslem(10, 4)
		fmt.Println("Toplam", sonuc1)
		//fmt.Println("Çıkarma", sonuc2)
		fmt.Println("Çarpım", sonuc3)
		//fmt.Println("Bölüm", sonuc4)
	*/