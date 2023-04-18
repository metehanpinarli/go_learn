package main

import (
	"fmt"
	"os"
	"os/signal"
)

func main() {
	//Sinyallerin gönderileceği kanalı oluşturalım.
	channelOne := make(chan os.Signal, 1)

	//Gelen sinyalleri kanal1'e yönlendirelim
	signal.Notify(channelOne)

	// kanal1'e sinyal gelene kadar programı bekletelim.
	signalType := <-channelOne
	fmt.Println("Sinyal Türü:", signalType)
}
