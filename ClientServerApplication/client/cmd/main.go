package main

import (
	"ClientServerApplication/lib"
	"time"
)

const (
	minClients = 2
	maxClients = 100

	minSendTime = 1
	maxSendTime = 30
)

func main() {
	//clients := lib.GenerateClients(minClients + rand.Intn(maxClients-minClients+1))
	clients := lib.GenerateClients(100)
	for _, client := range clients {
		_, _ = client.Send("http://localhost:8080/", &lib.Report{client.Id, lib.GenerateReportType()})
		time.Sleep(time.Second)
	}
}
