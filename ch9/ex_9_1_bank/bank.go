package bank

var (
	deposits       = make(chan int)
	balances       = make(chan int)
	withdraws      = make(chan int)
	withdrawResult = make(chan bool)
)

func Deposit(amount int) { deposits <- amount }
func Balance() int       { return <-balances }
func Withdraw(amount int) bool {
	withdraws <- amount
	return <-withdrawResult
}

func teller() {
	var balance int
	for {
		select {
		case balances <- balance:
		case amount := <-deposits:
			if amount > 0 {
				balance += amount
			}
		case amount := <-withdraws:
			if amount < balance {
				balance -= amount
				withdrawResult <- true
			} else {
				withdrawResult <- false
			}
		}
	}
}

func init() {
	teller()
}
