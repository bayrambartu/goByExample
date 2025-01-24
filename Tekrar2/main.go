package main

import "fmt"

/*
var mt sync.Mutex

	func paraCek(bakiye *float64, cekilecekMiktar float64, wg *sync.WaitGroup) {
		mt.Lock()

		*bakiye -= 15
		fmt.Printf("Yeni Bakiye: %.2f\n", *bakiye)
		mt.Unlock()
		fmt.Println("çekme işlemi tamamlandı")
		wg.Done()
	}

	func paraYatir(bakiye *float64, yatirılacakMiktar float64, wg *sync.WaitGroup) {
		mt.Lock()
		*bakiye += 65
		fmt.Printf("Yeni Bakiye: %.2f\n", *bakiye)
		mt.Unlock()
		fmt.Println("Yatırma işlemi tamamlandı.")
		wg.Done()
	}
*/

type Request struct {
	Action  string
	Key     string
	Value   int
	ReplyCh chan interface{}
}

/*
type WriteRequest struct {
	Key   string
	Value int
}*/

func main() {
	/*
		data := make(map[string]int)
		writeRequest := make(chan WriteRequest)
		done := make(chan bool)

		// Map yönetimi
		go func() {
			for req := range writeRequest {
				data[req.Key] = req.Value
				fmt.Printf("Set %s -> %d\n", req.Key, req.Value)
			}
			done <- true
		}()

		// Yazma işlemleri
		writeRequest <- WriteRequest{Key: "a", Value: 10}
		writeRequest <- WriteRequest{Key: "b", Value: 20}

		close(writeRequest)
		<-done
		fmt.Println("Final Map State", data)

	*/
	/*
		var wg sync.WaitGroup
		wg.Add(2)
		var bakiye float64 = 100
		fmt.Printf("İlk Bakiye: %.2f\n", bakiye)

		go paraCek(&bakiye, 25, &wg)
		go paraYatir(&bakiye, 65, &wg)
		wg.Wait()*/

	/*
		var mu sync.Mutex
		data := make(map[string]int)

		mu.Lock()
		data["key1"] = 100
		data["key2"] = 200
		mu.Unlock()

		mu.Lock()
		fmt.Println("Map durumu:", data)
		mu.Unlock()

		mu.Lock()
		delete(data, "key1")
		fmt.Println("Veri silindi:key1")
		mu.Unlock()

		mu.Lock()
		fmt.Println("son Map durumu:", data)
		mu.Unlock()
	*/
	/*
		var mu sync.Mutex
		data := make(map[string]int)
		var wg sync.WaitGroup

		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			data["key1"] = 100
			data["key2"] = 200
			fmt.Println("Veriler eklendi")
			mu.Unlock()
		}()

		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			fmt.Println("map durumu:", data)
			mu.Unlock()
		}()

		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			delete(data, "key1")
			fmt.Println("Delete eklendi : key1")
			mu.Unlock()
		}()
		wg.Wait()
		fmt.Println("tüm işlemler tamamlandı")
		// her goroutine sync.Mutex kullanarak map'e güvenli bir şekilde erişiyor.
		//sync.WaitGroup tüm işlemleri senkronize ediyor ve ana goroutine diğer işlemleri bekliyor.

	*/

	// kanal tabanlı Map yönetimi
	data := make(map[string]int)
	requests := make(chan Request)

	go func() {
		for req := range requests {
			switch req.Action {
			case "set":
				data[req.Key] = req.Value
				fmt.Printf("Set: %s -> %d\n", req.Key, req.Value)
				req.ReplyCh <- "success"
			case "get":
				value, exists := data[req.Key]
				if exists {
					req.ReplyCh <- value
				} else {
					req.ReplyCh <- "not found"
				}
			case "delete":
				delete(data, req.Key)
				fmt.Printf("Deleted: %s\n", req.Key)
				req.ReplyCh <- "deleted"
			}
		}

	}()

	replyCh := make(chan interface{})
	requests <- Request{Action: "set", Key: "a", Value: 42, ReplyCh: replyCh}
	fmt.Println(<-replyCh)

	requests <- Request{Action: "get", Key: "a", ReplyCh: replyCh}
	fmt.Println("Value of a:", <-replyCh)

	requests <- Request{Action: "delete", Key: "a", ReplyCh: replyCh}
	fmt.Println(<-replyCh)

	requests <- Request{Action: "get", Key: "a", ReplyCh: replyCh}
	fmt.Println("Value of a:", <-replyCh)

	close(requests)

	/*
		data := make(map[string]int)
		data["a"] = 10
		data["b"] = 20

		// veri okuma
		fmt.Println("Value of a:", data["a"])

		// veri silme
		delete(data, "a")
		fmt.Println("Value of a after deletion:", data["a"])
	*/

	/*
		data := make(map[string]int)
		done := make(chan bool)

		go func() {
			data["a"] = 10
			data["b"] = 20
			fmt.Println("Veriler eklendi:", data)

			// map'ten veri silme
			delete(data, "a")
			fmt.Println("Value of a after deletion:", data["a"])

			done <- true

		}()

		<-done // işlemler tamamlanana kadar bekle
		fmt.Println("tüm işlemler tamamlandı") */

}
