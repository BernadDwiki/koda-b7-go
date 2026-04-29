package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	chn := make(chan string, 10)

	pengirimPesan("Andi", "Halo, apa kabar?", chn)
	pengirimPesan("Budi", "Selamat pagi!", chn)
	pengirimPesan("Citra", "Semoga harimu menyenangkan!", chn)

	wg.Add(1)
	go papanTulis(chn, &wg)
	close(chn)
	wg.Wait()
	fmt.Println("Semua pesan sudah ditulis di papan tulis!")
}

func pengirimPesan(pengirim string, message string, chn chan string) {
	pesan := fmt.Sprintf("Pesan dari %s : %s", pengirim, message)
	// wg.Add(1)
	chn <- pesan
}

func papanTulis(messageChn chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for message := range messageChn {
		fmt.Printf("Papan tulis menulis pesan: %s\n", message)
	}
}
