package main
import "fmt"
type payment struct {
	 gateway paymenter
}

type paymenter interface {
	pay(amount float64)
	refund(amount float64,account string)
}

func (p payment) makePayment(amount float64){
	// razorpayPaymentGw := razorpay{}
	// razorpayPaymentGw.pay(amount)
	// stripePaymentGw := stripe{}
	// stripePaymentGw.pay(amount)
	p.gateway.pay(amount)
}

type razorpay struct {}

func (r razorpay) pay(amount float64){
	// logic to make payment 
	fmt.Println("making payment using razorpay", amount)
}

type stripe struct {}

func(s stripe) pay(amount float64){
	// logic to make payment 
	fmt.Println("making payment using stripe", amount)
}

func(s stripe) refund(amount float64, account string){
	// logic to refund payment
	fmt.Println("refunding payment using stripe", amount, account)
}
type paypal struct {}

func(p paypal) pay(amount float64){	
	fmt.Println("making payment using paypal", amount)
}

func(p paypal) refund(amount float64, account string){
	// logic to refund payment
	fmt.Println("refunding payment using paypal", amount, account)
}

// dependecy inversion principle 
// also fix open closed principle
func main(){
	// stripePaymentGw := stripe{}
	// razorpayPaymentGw := razorpay{}
	paymentGw := paypal{}
	newPayment := payment{
		gateway: paymentGw,
	}
	newPayment.makePayment(1000)
}