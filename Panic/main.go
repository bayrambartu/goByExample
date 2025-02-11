package main

import "fmt"

func bolme(a, b int) int {
	if b == 0 {
		panic("sifira bolme hatasi")
	}
	return a / b
}
func main() {
	fmt.Println(bolme(10, 0))
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
