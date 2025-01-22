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
	/*ch := make(chan string)
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
	}*/

	ch1 := make(chan string)
	ch2 := make(chan string)
	go func() {
		time.Sleep(time.Second * 2)
		ch1 <- "kanal 1 gorev tamamlandi!"
	}()
	go func() {
		time.Sleep(time.Second * 6)
		ch2 <- "kanal 2 gorev tamamlandi!"
	}()

	// kanal basina bir zaman asimi belirliyoruz.!!
	timeout1 := time.After(3 * time.Second)
	timeout2 := time.After(5 * time.Second)

	for {
		select {
		case msg := <-ch1:
			fmt.Println(msg)
			ch1 = nil // kanal islenince artik dinlenmeyecek
			timeout1 = nil

		case <-timeout1:
			fmt.Println("kanal 1 zaman asimina ugradi")
			ch1 = nil
			timeout1 = nil

		case msg := <-ch2:
			fmt.Println(msg)
			ch2 = nil
			timeout2 = nil

		case <-timeout2:
			fmt.Println("Kanal 2 zaman asimina ugradi")
			ch2 = nil
			timeout2 = nil

		}
		if ch1 == nil && ch2 == nil {
			break
		}
	}
	fmt.Println("Tum isler tamamlandi")
}
