package main

import (
	"study/application/service"
	"study/infrastructure/queue"
)

func main() {
	forever := make(chan bool)
	rabbit := queue.NewRabbit()
	sync := service.NewSync()
	go rabbit.Listen(sync.Execute)
	sync.StreamWorker.Wait()
	<-forever
}
