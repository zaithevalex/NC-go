package lib

import (
	"math/rand"
	"net/http"
)

func GenerateClients(amount int) []*Client {
	clients := make([]*Client, amount)
	for i := 0; i < amount; i++ {
		clients[i] = &Client{i, &http.Client{}}
	}

	return clients
}

func GenerateReportType() ReportType {
	return ReportType(rand.Intn(amountReportTypes))
}
