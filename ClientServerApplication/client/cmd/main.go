package main

import (
	"ClientServerApplication/lib/client"
	"ClientServerApplication/lib/report"
	"log"
	"math/rand"
	"net/http"
	"time"
)

// consts
const (
	minClients  = 2
	maxClients  = 10
	minSendTime = 1
	maxSendTime = 30
)

func main() {
	// generate clients
	amountOfClients := minClients + rand.Intn(maxClients-minClients+1)
	clients, done := make([]*client.Client, amountOfClients), make(chan bool)
	for i := 0; i < amountOfClients; i++ {
		clients[i] = &client.Client{Id: i, Client: &http.Client{}}
	}

	log.Println("Amount of clients:", len(clients))

	for id, c := range clients {
		go send(id, c)
	}

	<-done
}

// send is designed to send requests from each goroutine.
func send(id int, client *client.Client) {
	for {
		r := &report.Report{Id: client.Id, Type: report.ReportType(rand.Intn(report.AmountReportTypes))}
		delay := time.Duration(minSendTime+rand.Intn(maxSendTime-minSendTime+1)) * time.Second

		_, _ = client.Send("http://localhost:8080/", r)
		log.Printf("sent report `{%v, %v}` with processing delay %v\n", r.Id, r.Type, delay)
		time.Sleep(delay)
	}
}
