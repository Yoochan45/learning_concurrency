package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// kode warna ANSI untuk bikin output terminal jadi warna-warni
const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
)

// struct sederhana buat ngebungkus nama task
type Task struct {
	Name string
}

// fungsi driver = worker yang ngerjain task
func driver(id int, tasks <-chan Task, wg *sync.WaitGroup) {
	defer wg.Done() // kalau fungsi selesai, tandai 1 goroutine selesai ke WaitGroup

	// loop ambil task dari channel sampai channel ditutup
	for task := range tasks {
		var driverColor string
		if id == 1 {
			driverColor = Green // driver 1 = hijau
		} else {
			driverColor = Blue // driver lain = biru
		}

		// warna default task
		taskColor := Reset
		if task.Name == "Emergency Medical Transport" {
			taskColor = Red // kalau nama task tertentu, kasih warna merah
		}

		// print siapa driver-nya dan task apa yang dikerjain
		fmt.Printf("%sDriver %d:%s Processing Task: %s%s%s\n",
			driverColor, id, Reset, taskColor, task.Name, Reset)

		// simulasi kerja butuh waktu random 200–700 ms
		time.Sleep(time.Duration(rand.Intn(500)+200) * time.Millisecond)

	}
}

// fungsi simulasi sekali run
func runSimulation(runID int, drivers int, jobList []Task) {
	fmt.Printf("Run %d:\n\n", runID)

	// bikin channel buat distribusi task
	taskChan := make(chan Task)
	var wg sync.WaitGroup

	// nyalain beberapa driver (sesuai jumlah drivers)
	for i := 1; i <= drivers; i++ {
		wg.Add(1)                   // nambah 1 counter ke WaitGroup
		go driver(i, taskChan, &wg) // jalanin driver sebagai goroutine
	}

	// kirim semua task ke channel biar driver bisa ambil
	for _, t := range jobList {
		taskChan <- t
	}

	// tutup channel → sinyal kalau ga ada task baru
	close(taskChan)

	// tunggu semua goroutine driver selesai ngerjain task
	wg.Wait()

	// semua task udah selesai
	fmt.Printf("%sAll tasks have completed%s\n\n", Yellow, Reset)
}

func main() {
	// seed random biar hasil Sleep beda tiap run
	rand.Seed(time.Now().UnixNano())

	// daftar pekerjaan
	jobs := []Task{
		{"Emergency Medical Transport"},
		{"Delivery at Zone A"},
		{"Pickup at Zone B"},
	}

	// jalanin simulasi 3 kali, tiap run ada 2 driver
	for i := 1; i <= 3; i++ {
		runSimulation(i, 2, jobs)
	}
}
