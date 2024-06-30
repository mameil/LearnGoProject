package main

import (
	"fmt"
	"time"
)

type Courier struct {
	name string
}

type Product struct {
	name  string
	price int
	ID    int
}

type Parcel struct {
	pdt           *Product
	ShippedTime   time.Time
	DeliveredTime time.Time
}

func (c *Courier) SendProduct(product *Product) Parcel {
	return Parcel{pdt: product, ShippedTime: time.Now()}
}

func (p *Parcel) Delivered() *Product {
	p.DeliveredTime = time.Now()
	return p.pdt
}

func main() {
	courier := Courier{name: "shim"}
	product := Product{name: "laptop", price: 1000000, ID: 1}

	parcel := courier.SendProduct(&product)
	fmt.Println(parcel)
}
