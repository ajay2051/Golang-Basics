package main

import "fmt"

type OrderStatus int

const (
	Received OrderStatus = iota
	Confirmed
	Prepared
	Delivered
)

// for string
// type Orderstatus string
// const (
// Received OrderStatus = "received"
// Confirmed = "confirmed"
// )

func changeOrderstatus(status OrderStatus) {
	fmt.Println("Changed Order Status", status)
}

func main() {
	changeOrderstatus(Received)
}
