package golang_goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// dilakukan untuk menunggu sebuah proses selesai dilakukan
// ketika ada beberapa go routine dijalankan dengan menggunakan wg anda dapat menunggu semua goroutine selesai dijalankan baru aplikasi akan terhenti

func RunAsynchronoous(wg *sync.WaitGroup) {
	// mengurangi
	defer wg.Done()

	// menaikkan jumlah tugas yang harus ditunggu
	wg.Add(1)

	fmt.Println("Hello")
	time.Sleep(1 * time.Second)
}

func TestWaitGroup(t *testing.T) {

	//initialized wg
	group := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go RunAsynchronoous(group)
	}

	// proses untuk menunggu
	group.Wait()
	fmt.Println("Proses menunggu telah selesai")
}
