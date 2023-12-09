package golang_goroutine

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Arief Mahendra"
		// block hingga ada data yang dikirimkan ke channel dan diambil oleh channel lain
		fmt.Println("Selesai mengirim data ke channel")
	}()

	data := <-channel
	fmt.Println(data)
	time.Sleep(5 * time.Second)
}

// Channel sebagai parameter
// untuk channel sudah otomatis pass by reference
func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Arief Mahendra"
}

func TestChanAsParameter(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go GiveMeResponse(channel)

	data := <-channel
	fmt.Println(data)
	time.Sleep(5 * time.Second)
}

// Channel in hanya untuk mengirim data ke channel
func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "Arief Mahendra"
}

// Channel out hanya untuk mengambil data dari channel
func OnlyOut(channel <-chan string) {
	data := <-channel
	fmt.Println(data)
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(2 * time.Second)
}

// buffered channel adalah channel yang bisa menampung data sebanyak buffer yang kita tentukan (kapasitas)
// buffered channel akan memblok goroutine yang mengirim data ke buffered channel jika buffer penuh
// secara default channel adalah unbuffered channel (buffer = 0)
// cocok digunakan untuk mengirim data dari goroutine yang berbeda ke satu channel tanpa harus menunggu proses pengambilan data oleh channel lain selesai

func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 3)
	defer close(channel)

	channel <- "Arief"

	fmt.Println(<-channel)

	channel <- "Mahendra"

	// visualnya seperti ini : [Arief Mahendra] (buffer = 3)

	// menghitung panjang data yang ada di channel
	fmt.Println(len(channel))

	// menghitung kapasitas channel
	fmt.Println(cap(channel))
}

// range channel
// digunakan untuk melakukan iterasi terhadap data yang ada di channel
// range channel akan terus berjalan hingga channelnya ditutup

func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- "Perulangan ke " + strconv.Itoa(i)
		}

		// menutup channel agar range channel bisa berhenti atau deadlock
		close(channel)
	}()

	for data := range channel {
		fmt.Println("Menerima data", data)
	}
}

// select channel adalah proses seleksi data yang ada di channel
// select channel digunakan untuk mengecek data yang ada di channel secara bergantian
// select channel digunakan untuk menangani kasus dimana kita tidak tahu kapan data akan dikirimkan ke channel
// select channel akan mengecek semua channel yang ada di case
