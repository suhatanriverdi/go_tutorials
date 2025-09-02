package main

import (
	"fmt"
	"os"
	"strconv"
)

func writeToFile(bytesChannel chan []byte, doneChannel chan bool, errChannel chan error) {
	// Sample destinations
	// file, err := os.Create("/home/user/Desktop/output.txt")  // Linux/Mac
	// OR
	// file, err := os.Create("C:\\Users\\user\\Desktop\\output.txt")  // Windows
	//
	// Create a file
	file, err := os.Create("output.txt")
	if err != nil {
		errChannel <- err
		return
	}

	// Close the file when done
	defer file.Close()

	// Write incoming data to the file
	for data := range bytesChannel {
		_, err := file.Write(data)
		if err != nil {
			errChannel <- err
			return
		}
	}

	doneChannel <- true
}

func main() {
	// Create Unbuffered Channels
	bytesChannel := make(chan []byte)
	doneChannel := make(chan bool)
	errChannel := make(chan error)

	// Goroutine başlat
	go writeToFile(bytesChannel, doneChannel, errChannel)

	// Sample data send
	// Equals to:
	// data := []byte{72, 101, 108, 108, 111}  // H=72, e=101, l=108, l=108, o=111
	// bytesChannel <- data
	bytesChannel <- []byte("Hello")
	bytesChannel <- []byte("World")

	// Channel'ı kapat (writeToFile fonksiyonunun döngüden çıkması için)
	close(bytesChannel)

	// Tamamlanmasını bekle veya hata kontrolü yap
	select {
	case <-doneChannel:
		fmt.Println("Dosya yazma işlemi tamamlandı")
	case err := <-errChannel:
		fmt.Printf("Hata: %v\n", err)
	}

	// Dosyayı oku ve ekrana yazdır
	content, err := os.ReadFile("output.txt")
	if err != nil {
		fmt.Printf("Dosya okuma hatası: %v\n", err)
		return
	}

	fmt.Printf("Dosya içeriği: %s\n", string(content))
}

// String <> Int Conversion Examples
func stringIntConversions() {
	// String to Int
	str := "123"
	num, err := strconv.Atoi(str)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	fmt.Printf("String '%s' -> Int %d\n", str, num)

	// Int to String
	number := 456
	text := strconv.Itoa(number)
	fmt.Printf("Int %d -> String '%s'\n", number, text)

	// Alternative methods
	text2 := fmt.Sprintf("%d", number)
	fmt.Printf("Int %d -> String '%s' (Sprintf ile)\n", number, text2)

	// ParseInt and FormatInt
	num2, err := strconv.ParseInt("789", 10, 64) // base 10, 64 bit
	if err == nil {
		fmt.Printf("ParseInt: %d\n", num2)
	}

	str2 := strconv.FormatInt(int64(number), 10) // Decimal bytes
	fmt.Printf("FormatInt: %s\n", str2)
}
