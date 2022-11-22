package error_handling

import (
	"fmt"
	"os"
)

func Demo1() {
	f.err := os.Open("demo1.txt") //dosyayı bulursa f (dosyayı açar), dosya yoksa err çalışır

	if err != nil{
		if pErr, ok := err.(*os.PathError); ok{ //Dosya yolunda hata varsa çalışır
			fmt.Println("Dosya bulunamadı", pErr.Path)
			return
		} else {
			fmt.Println("Dosya bulunamadı")
			return
		}
	} else {
		fmt.Println("Dosya mevcut", f.Name())
	}
}
