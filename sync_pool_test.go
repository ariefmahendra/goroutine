package golang_goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// implementasi object pool patern
// sederhananya, pool digunakan untuk menyimpan data, selanjutnya jika membutuhkan data tersebut maka kita dapat mengambil dari pool
// dan jika sudah selesai maka kita dapat mengembalikan data tersebut ke pool
// implementasi pool di golang ini sudah aman dari problem race condition

func TestPool(t *testing.T) {

	pool := sync.Pool{}

	// proses memasukkan data kedalam pool
	pool.Put("Arief")
	pool.Put("Mahendra")

	for i := 0; i < 10; i++ {
		go func() {
			// mengambil data pool dipindahkan ke variabel data
			data := pool.Get()

			fmt.Println(data)

			// mengembalikan data yang telah diambil ke pool
			pool.Put(data)
		}()
	}

	time.Sleep(3 * time.Second)
	fmt.Println("Selesai")
}
