package main

import (
	"errors"
	"fmt"
)

/*type error interface {
	Error() string
}*/

/*func myFunction(arg int) (int, error) {
	if arg < 0 {
		return 0, errors.New("Negatif sayi kabul edilemez")
	}
	return arg * 2, nil
}*/

/*
	func checkValue(x int) error {
		if x == 42 {
			return errors.New("42 ile işlem yapılmaz")
		}
		return nil // hata yok
	}

	func checkWithContext(x int) error {
		if x == 42 {
			return fmt.Errorf("veri işlemedi: %d", x)
		}
		return nil
	}

var ErrOutOfTea = errors.New("out of tea")
var ErrPower = errors.New("power")

	func checkTeaStock(teaAmount int) error {
		if teaAmount == 0 {
			return ErrOutOfTea
		}
		return nil
	}
*/
func f(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("42 ile calisamam")
	}
	return arg + 3, nil
}

var ErrOutOfTea = fmt.Errorf("out of tea ")
var ErrPower = fmt.Errorf("suyu kaynatamiyoruz")

func makeTea(arg int) error {
	if arg == 2 {
		return ErrOutOfTea
	} else if arg == 4 {
		return fmt.Errorf("cay yapma hatasi: %w", ErrPower)
	}
	return nil
}
func main() {
	for _, i := range []int{7, 42} {
		if r, e := f(i); e != nil {
			fmt.Println("f basarisiz:", e)
		} else {
			fmt.Println("f basarili:", r)
		}
	}
	for i := 0; i < 5; i++ {
		if err := makeTea(i); err != nil {
			if errors.Is(err, ErrOutOfTea) {
				fmt.Println("yeni cay al")
			} else if errors.Is(err, ErrPower) {
				fmt.Println("elektrik kesildi!!")
			} else {
				fmt.Printf("bilinmeyen hata: %s \n", err)
			}
			continue
		}
		fmt.Println("cay hazir")
	}

	/*fmt.Println(checkValue(42))
	fmt.Println(checkWithContext(42))
	fmt.Println(checkTeaStock(0))

	err := checkTeaStock(0)
	if errors.Is(err, ErrOutOfTea) {
		fmt.Println("out of tea")
	}*/

}
