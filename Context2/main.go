package main

import (
	"context"
	"fmt"
	"time"
)

type Product struct {
	Id   int64
	Name string
}

var productChannel = make(chan Product)

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "correlationID", "abc123")
	F1(ctx)

	/*ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	go getProductDetailsFromExternalService(10)
	select {
	case productDetail := <-productChannel:
		fmt.Println("Product details : ", productDetail)

	case <-ctx.Done():
		fmt.Println("timeout..")
	}
	fmt.Println("End of the main")
	*/
}

func F1(ctx context.Context) {
	fmt.Println("f1", ctx.Value("correlationID"))
	F2(ctx)
}
func F2(ctx context.Context) {
	fmt.Println("f2", ctx.Value("correlationID"))
	F3(ctx)

}
func F3(ctx context.Context) {
	fmt.Println("f3", ctx.Value("correlationID"))
}

func getProductDetailsFromExternalService(productId int64) {
	time.Sleep(2 * time.Second)

	productChannel <- Product{
		Id:   productId,
		Name: "washing machine",
	}

}
