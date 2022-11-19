package main

import "fmt"

func main() {
	tekCn := make(chan int)
	ciftCn := make(chan int)
	go channels.TekSayilar(tekCn)
	go channels.CiftSayilar(ciftCn)

	ciftSayiToplam, tekSayiToplam := <-ciftCn, <-tekCn

	carpim := ciftSayiToplam * tekSayiToplam
	fmt.Println("Çarpım:", carpim)
}
