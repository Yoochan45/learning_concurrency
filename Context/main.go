package main

import (
	"context" // Package untuk mengontrol eksekusi goroutine (cancel, timeout, pass data)
	"fmt"
	"time"
)

// Worker simulasi pekerjaan goroutine
// id : nomor worker
// ctx : context untuk nge-handle stop/cancel worker
func Worker(id int, ctx context.Context) {
	for {
		select {
		// <-- Cek channel dari context. Jika context di-cancel, channel akan close
		case <-ctx.Done():
			fmt.Println("Worker", id, "dihentikan") // worker berhenti
			return                                  // keluar dari loop & goroutine
		default:
			// Jika belum di-cancel, worker tetap jalan
			fmt.Println("Worker", id, "dijalankan")
			time.Sleep(500 * time.Millisecond) // simulasi kerja
		}
	}
}

func main() {
	// Buat context cancelable. Bisa dipakai untuk stop semua goroutine yang terhubung
	ctx, cancel := context.WithCancel(context.Background())

	// Jalankan 3 worker secara konkuren
	for i := 1; i <= 3; i++ {
		go Worker(i, ctx) // worker menerima context untuk bisa dihentikan
	}

	time.Sleep(2 * time.Second) // biarkan worker jalan dulu selama 2 detik
	fmt.Println("Stop semua proses")
	cancel() // cancel context → semua worker akan menerima sinyal stop

	time.Sleep(1 * time.Second) // beri waktu worker nge-print "dihentikan" sebelum program exit
}
