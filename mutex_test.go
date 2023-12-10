package golang_goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type UserBalance struct {
	sync.Mutex
	Name    string
	Balance int
}

func (user *UserBalance) Lock() {
	user.Mutex.Lock()
}

func (user *UserBalance) Unlock() {
	user.Mutex.Unlock()
}

func (user *UserBalance) change(amount int) {
	user.Balance += amount
}

func Transfer(user1 *UserBalance, user2 *UserBalance, amount int) {
	user1.Lock()
	fmt.Println("Lock user 1", user1.Name)
	user1.change(-amount)

	time.Sleep(1 * time.Second)

	user2.Lock()
	fmt.Println("Lock user 2", user2.Name)
	user2.change(amount)

	time.Sleep(1 * time.Second)

	user1.Unlock()
	user2.Unlock()
}

func TestDeadLock(t *testing.T) {
	user1 := UserBalance{
		Name:    "Arief",
		Balance: 1000000,
	}

	user2 := UserBalance{
		Name:    "Juan",
		Balance: 5000000,
	}

	go Transfer(&user1, &user2, 100000)
	go Transfer(&user2, &user1, 200000)

	fmt.Println("User 1 balance :", user1.Balance)
	fmt.Println("User 2 balance :", user2.Balance)

	time.Sleep(5 * time.Second)
}
