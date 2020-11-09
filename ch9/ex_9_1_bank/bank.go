package bank

var (
	deposits  = make(chan int)
	balances  = make(chan int)
	withdraws = make(chan withdrawal)
)

type withdrawal struct {
	amount  int
	success chan bool
}

func Deposit(amount int) { deposits <- amount }
func Balance() int       { return <-balances }
func Withdraw(amount int) bool {
	w := withdrawal{amount, make(chan bool)}
	withdraws <- w
	return <-w.success
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
		case w := <-withdraws:
			if w.amount < balance {
				balance -= w.amount
				w.success <- true
			} else {
				w.success <- false
			}
		}
	}
}

func init() {
	go teller()
}
