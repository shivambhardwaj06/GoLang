package main
import "fmt"

// enumerated types 
type OrderStatus int // custom type

const (
	Received OrderStatus = iota 
	Confirmed
	Prepared
	Delivered
	Canceled
)

type Orderstatus string

const (
	received Orderstatus = "Received"
	confirmed            = "Confirmed"
	prepared             = "Prepared"
	delivered            = "Delivered"
	canceled             = "Canceled"
)


func changeOrderStatus(status Orderstatus) {
	fmt.Println("Order status changed to:", status)	
}

func main() {
	// changeOrderStatus(Received)
	// changeOrderStatus(Confirmed)
	// changeOrderStatus(Prepared)
	// changeOrderStatus(Delivered)
	// changeOrderStatus(Canceled)
	// changeOrderStatus(5) // this will not give any error but it will not be a valid status
	// changeOrderStatus(Delivered) 
	changeOrderStatus(received)


} 

