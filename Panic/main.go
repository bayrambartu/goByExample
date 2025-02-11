package main

import (
	"fmt"
	"os"
)

/*
	func bolme(a, b int) int {
		if b == 0 {
			panic("sifira bolme hatasi")
		}
		return a / b
	}
*/
func main() {
	/*
		fmt.Println(bolme(10, 2))
		fmt.Println(bolme(10, 0))*/

	_, err := os.Create("/root/test.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("Dosya olusturuldu")
	/*

		panic("a problem")
		_, err := os.Create("/tmp/file")
		if err != nil {
			panic(err)
		}
	*/
	/*
		fmt.Println("Program başlıyor...")
		panic("Bir hata oluştu!")
		fmt.Println("Bu satır çalışmaz!")
	*/

	/*
		var x *int
		fmt.Println(*x) // deferece hatası
	*/

}
