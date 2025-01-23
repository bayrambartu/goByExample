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

func main() {
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
}
