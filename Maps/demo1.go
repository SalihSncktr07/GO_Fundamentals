package maps

import "fmt"

func Demo1() {
	sozluk := make(map[string]string)
	sozluk["table"] = "masa"
	sozluk["book"] = "kitap"
	sozluk["pencil"] = "kalem"

	fmt.Println(sozluk)
	fmt.Println(sozluk["book"])
	fmt.Println("Eleman sayısı:", len(sozluk))
	delete(sozluk, "book")
	fmt.Println(sozluk)
	fmt.Println("Eleman sayısı:", len(sozluk))

	deger, varMi := sozluk["pencil"]
	fmt.Println(deger)
	fmt.Println("Listede var mı:", varMi)

	sozluk2 := map[string]string{"glass": "cam", "door": "kapı"}
	fmt.Println(sozluk2)
}
