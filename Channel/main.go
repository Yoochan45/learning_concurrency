package main

import (
	"fmt"
	"time"
)

// Worker function
// id      : nomor worker
// jobs    : channel read-only, menerima pekerjaan
// results : channel write-only, mengirim hasil pekerjaan
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs { // loop terus sampai channel jobs ditutup
		fmt.Println("Worker", id, "mengerjakan job", j)
		time.Sleep(time.Second) // simulasi pekerjaan 1 detik
		results <- j * 2        // kirim hasil ke channel results
	}
}

func main() {
	// Buat channel jobs dan results dengan kapasitas 5 (buffered)
	jobs := make(chan int, 5)
	results := make(chan int, 5)

	// Jalankan 3 worker secara konkuren
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// Kirim 5 pekerjaan ke channel jobs
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs) // penting: tutup channel jobs supaya worker berhenti loop

	// Ambil semua hasil dari channel results
	for a := 1; a <= 5; a++ {
		fmt.Println("Hasil:", <-results, "Dari Worker", a)
	}
}
