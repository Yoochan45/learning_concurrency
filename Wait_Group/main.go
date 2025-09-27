package main

import (
	"fmt"
	"sync"
	"time"
)

// Worker itu fungsi untuk mensimulasikan pekerjaan tiap worker
func Worker(id int, wg *sync.WaitGroup) {
	fmt.Println("Worker :", id, "dimulai") // Cetak worker mulai
	time.Sleep(10 * time.Millisecond)      // Simulasi kerja sebentar
	fmt.Println("Worker :", id, "selesai") // Cetak worker selesai

	wg.Done() // Memberitahu WaitGroup kalau worker ini selesai
}

func main() {
	var wg sync.WaitGroup // Buat WaitGroup baru

	for i := 0; i < 3; i++ { // Loop 3 kali buat bikin 3 worker
		wg.Add(i)         // Tambah counter WaitGroup
		go Worker(i, &wg) // Jalankan worker di goroutine
	}

	wg.Wait() // Tunggu semua worker selesai
	fmt.Println("Semua worker selesai")
}
