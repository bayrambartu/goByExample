package main

import (
	"fmt"
	"sync"
	"time"
)

/*
	func islem(i int) {
		fmt.Printf("Islem %d baslasi \n", i)
		time.Sleep(time.Second * 2)
		fmt.Printf("Islem %d tamamlandi \n", i)
	}
*/

/*func islem(i int) {
	fmt.Printf("Islemlem %d başladi\n", i)
	time.Sleep(time.Duration(i) * time.Second) // Islem suresi degisken
	fmt.Printf("Islem %d tamamlandi\n", i)
}*/

/*
	func islem(i int, wg *sync.WaitGroup) {
		defer wg.Done()
		fmt.Printf("islem %d basladi \n", i)
		time.Sleep(time.Duration(i) * time.Second)
		fmt.Printf("islem %d finished \n", i)

}
*/
func islem(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	start := time.Now()
	fmt.Printf("islem %d basladi. Baslama zamani: %s \n", i, start.Format("15:04:05"))
	time.Sleep(time.Duration(i) * time.Second)
	end := time.Now()
	fmt.Printf("islem %d bitti. Bitis zamani: %s \n", i, end.Format("15:04:05"))

}

func main() {
	fmt.Println("Ana fonksiyon basliyor.")

	var wg sync.WaitGroup
	wg.Add(3)

	go islem(2, &wg)
	go islem(1, &wg)
	go islem(3, &wg)

	wg.Wait()
	fmt.Println("Ana fonksiyon tamamlandi")

	/*
		fmt.Println("Ana fonksiyon basladi")

		var wg sync.WaitGroup
		wg.Add(3)

		go islem(2, &wg) // 2 sn surecek
		go islem(1, &wg) // 1 sn surecek
		go islem(3, &wg) // 3 sn surecek

		wg.Wait()
		fmt.Println("Ana fonksyion tamamlandi")

	*/
	/*
		fmt.Println("Ana fonksiyon basladi")

		go islem(2) // 2 sn bekle
		go islem(1) // 1 sn
		go islem(3) // 3 sn

		time.Sleep(6 * time.Second)

		fmt.Println("Ana fonksiyon tamamlandi")
	*/

	/*
		fmt.Println("Ana fonksiyon basladi")

		go islem(1)
		time.Sleep(time.Second * 3)
		fmt.Println("Ana fonksiyon basladi")
	*/
}
