package golang_goroutine

import (
	"fmt"
	"sync"
	"testing"
)

// untuk memastikan sebuah function dieksekusi hanya sekali, dan jika ada banyak goroutine maka goroutine pertama yang akan dijalankan

var counter = 0

func OnlyOne() {
	counter++
}

func TestOnlyOne(t *testing.T) {
	var once sync.Once
	var group sync.WaitGroup

	for i := 0; i < 100; i++ {
		go func() {
			group.Add(1)
			// hanya function yang tidak memiliki parameter
			once.Do(OnlyOne)
			group.Done()
		}()
	}

	group.Wait()
	fmt.Println(counter) // isinya hanya akan ada 1 karena hanya 2 perulangan saja
	fmt.Println("proses menunggu telah selesai")
}
