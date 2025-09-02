package main

import (
	"fmt"
	"time"
)

// DEADLOCK örneği - Unbuffered channel, goroutine yok
func deadlockExample() {
	// ch := make(chan int) // unbuffered

	// ❌ Bu deadlock yapar!
	// ch <- 42  // Gönderen bekler ama alan yok!
	// fmt.Println("Bu satır çalışmaz")
}

// Doğru kullanım - Goroutine ile
func correctUsage() {
	ch := make(chan int) // unbuffered

	// Goroutine başlat (alıcı)
	go func() {
		value := <-ch // Burda bekliyor
		fmt.Printf("Alındı: %d\n", value)
	}()

	// Ana thread'de gönder
	ch <- 42 // ✅ Goroutine alacağı için blocking olmaz

	time.Sleep(100 * time.Millisecond) // Goroutine'in bitmesini bekle
}

// Bizim örneğimiz nasıl çalışıyor
func ourExample() {
	ch := make(chan []byte) // unbuffered
	done := make(chan bool)

	// Goroutine başlat (alıcı)
	go func() {
		fmt.Println("Goroutine başladı, veri bekliyor...")

		// Channel'dan veri oku
		for data := range ch {
			fmt.Printf("Alındı: %s\n", string(data))
		}

		done <- true
	}()

	// Ana thread'de gönder
	fmt.Println("Veri gönderiliyor...")
	ch <- []byte("Hello") // ✅ Goroutine bekliyor, çalışır
	ch <- []byte("World") // ✅ Goroutine bekliyor, çalışır

	close(ch) // Channel'ı kapat
	<-done    // Goroutine'in bitmesini bekle
}

// Buffered channel örneği
func bufferedExample() {
	ch := make(chan int, 2) // 2 elemanlık buffer

	// Buffer dolana kadar blocking olmaz
	ch <- 1 // ✅ Buffer'a gider
	ch <- 2 // ✅ Buffer'a gider
	// ch <- 3  // ❌ Bu blocking olur, buffer dolu!

	fmt.Println("Buffer dolu, şimdi okuyalım:")
	fmt.Println(<-ch) // 1
	fmt.Println(<-ch) // 2
}

func test() {
	fmt.Println("=== Doğru Kullanım (Goroutine ile) ===")
	correctUsage()

	fmt.Println("\n=== Bizim Örneğimiz ===")
	ourExample()

	fmt.Println("\n=== Buffered Channel ===")
	bufferedExample()
}
