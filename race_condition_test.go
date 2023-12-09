package golang_goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// race condition adalah kondisi dimana 2 goroutine atau lebih mengakses data yang sama secara bersamaan
// dan minimal 1 goroutine yang melakukan write ke data tersebut
// berbahaya karena bisa mengakibatkan data rusak

func TestRaceCondition(t *testing.T) {
	x := 0
	for i := 0; i < 1000; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				x = x + 1
			}
		}()
	}
	time.Sleep(5 * time.Second)
	fmt.Printf("Counter = %d", x)
}

// cara handle dengan menggunakan mutex
// mutex adalah sebuah lock yang digunakan untuk menjamin bahwa hanya 1 goroutine yang bisa mengakses data pada satu waktu
// cara kerja mutex adalah ketika 1 goroutine mengakses data, maka goroutine lain harus menunggu hingga mutex tersebut di unlock
// cara menggunakan mutex adalah dengan menggunakan function Lock dan Unlock

func TestMutex(t *testing.T) {
	x := 0
	var mutex sync.Mutex

	for i := 0; i < 1000; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				mutex.Lock()
				x = x + 1
				mutex.Unlock()
			}
		}()
	}
	time.Sleep(5 * time.Second)
	fmt.Printf("Counter = %d", x)
}
