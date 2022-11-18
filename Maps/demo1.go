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

	sozlukIS := map[int]string{1: "bir", 2: "iki", 3: "üç"}
	fmt.Println(sozlukIS)
	fmt.Println(sozlukIS[2])

	sozlukSI := map[string]int{"bir": 1, "iki": 2, "üç": 3}
	fmt.Println(sozlukSI)
	fmt.Println(sozlukSI["üç"])

	sozlukII :=map[int]int{1: 11, 2: 22, 3: 33}
	fmt.Println(sozlukII)
	fmt.Println(sozlukII[1])
}
