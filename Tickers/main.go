package main

import (
	"fmt"
	"time"
)

func main() {
	// belli aralıkta düzenli olarak bir işlemi gerçekleştirmek için kullanırız. Timer'lardan farkı -> sürekli sinyal göndermesidir.
	/*ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()

	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	done <- true
	fmt.Println("Ticker stopped") */

	/*ticker := time.NewTicker(2 * time.Second)
	done := make(chan bool)
	go func() {
		for {
			select {
			case <-done:
				fmt.Println("Health check stopped.")
				return
			case d := <-ticker.C:
				fmt.Println("Performing health check at", d)
			}
		}
	}()
	time.Sleep(10 * time.Second)
	done <- true
	fmt.Println("Ticker stopped.")*/
	/*
		ticker := time.Tick(time.Second * 1)
		for i := 1; i <= 3; i++ {
			<-ticker // 1 saniye bekler
			fmt.Printf("Periyodik işlem %d gerçekleşti\n", i)
		}
		fmt.Println("tüm işlemler tamamlandı.")*/

	/*timeout := time.After(3 * time.Second)
	fmt.Println("beklekniyor.(3sn)")
	<-timeout
	fmt.Println("3 saniye bekledikten sonra islem gerceklesti")*/

	/*
		ticker := time.Tick(1 * time.Second)
		timeout := time.After(5 * time.Second)
		fmt.Println("basliyor")
		for {
			select {
			case <-ticker:
				fmt.Println("Periyodik islem gerceklesti")
			case <-timeout:
				fmt.Println("zaman asimi")
				return
			}
		}*/

	// ZAMANLAYICI VE KANALLARI BİRLEŞTİRME
	ch := make(chan string)
	go func() {
		time.Sleep(2 * time.Second)
		ch <- "kanal uzerinden veri gonderildi"
	}()
	timeout := time.After(3 * time.Second)
	select {
	case msg := <-ch:
		fmt.Println(msg)
	case <-timeout:
		fmt.Println("Zaman Asimi")
	}

}
