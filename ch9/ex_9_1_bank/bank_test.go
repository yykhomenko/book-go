package bank

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBank(t *testing.T) {
	done := make(chan bool)
	go func() {
		Deposit(100)
		Withdraw(40)
		Deposit(80)
		done <- true
	}()

	go func() {
		Deposit(100)
		Withdraw(40)
		done <- true
	}()

	<-done
	<-done

	assert.Equal(t, 200, Balance())
}
